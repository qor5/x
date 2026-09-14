# Dynamic Configuration Design

## Background

Based on the discussion:
- Users need to dynamically modify configurations and have them take effect
- Configurations are stored in DB, with API for modifications
- Some configurations need encryption (e.g., OAuth2 client secret)
- Use `kx` library for encryption
- Cache with TTL as fallback, no need for go-bus broadcast mechanism

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           Application Layer                                  │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │                     DynamicConfig[T]                                 │    │
│  │  - Get(ctx) (T, error)      // Get config (auto cache handling)     │    │
│  │  - Reload(ctx) (T, error)   // Force reload from DB                 │    │
│  │  - Invalidate()             // Invalidate cache                     │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                           Storage Layer                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │                     ConfigStore[T]                                   │    │
│  │  - Load(ctx) (T, error)     // Load config from DB                  │    │
│  │  - Save(ctx, T) error       // Save config to DB                    │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                           Encryption Layer (kx)                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │                     kx.Manager                                       │    │
│  │  - EncryptStruct(ctx, obj, encCtx) (encObj, ciphertext, error)      │    │
│  │  - DecryptStruct(ctx, encObj, ciphertext, encCtx) (obj, error)      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              Database                                        │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  configs table                                                       │    │
│  │  - key: string (PK)         // Config identifier                    │    │
│  │  - data: jsonb              // Encrypted config data                │    │
│  │  - ciphertext: text         // Encrypted sensitive fields           │    │
│  │  - updated_at: timestamp                                            │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────────────────┘
```

## Core Components

### 1. DynamicConfig[T] - Cache Layer

Manages configuration caching with TTL.

```go
type DynamicConfig[T any] struct {
    store     ConfigStore[T]
    cache     T
    cacheSet  bool
    ttl       time.Duration
    expiresAt time.Time
    mu        sync.RWMutex
    validator Validator
}

type DynamicConfigOption[T any] func(*DynamicConfig[T])

// WithTTL sets cache TTL. If 0, no caching (always query DB).
func WithTTL[T any](ttl time.Duration) DynamicConfigOption[T]

// WithValidator sets config validator.
func WithValidator[T any](v Validator) DynamicConfigOption[T]
```

**Methods:**

| Method | Description |
|--------|-------------|
| `Get(ctx) (T, error)` | Get config. Returns cached value if valid, otherwise loads from DB. |
| `Reload(ctx) (T, error)` | Force reload from DB, update cache. |
| `Invalidate()` | Mark cache as invalid. Next `Get` will reload. |

**Behavior:**

```
Get(ctx):
  1. RLock, check if cache valid (cacheSet && time.Now() < expiresAt)
  2. If valid, return cached value
  3. If invalid, Lock, double-check, then call store.Load(ctx)
  4. Validate loaded config
  5. Update cache and expiresAt
  6. Return config
```

### 2. ConfigStore[T] - Storage Interface

Abstract interface for config persistence.

```go
type ConfigStore[T any] interface {
    // Load loads config from storage.
    Load(ctx context.Context) (T, error)
    
    // Save saves config to storage.
    Save(ctx context.Context, config T) error
}
```

### 3. EncryptedConfigStore[T] - Encrypted Storage Implementation

Implements `ConfigStore[T]` with kx encryption support.

```go
type EncryptedConfigStore[T any] struct {
    db                *gorm.DB
    kxManager         *kx.Manager
    configKey         string
    encryptionContext map[string]string
}

// ConfigRecord is the DB model
type ConfigRecord struct {
    Key        string    `gorm:"primaryKey"`
    Data       []byte    `gorm:"type:jsonb"`       // Config with sensitive fields removed
    Ciphertext string    `gorm:"type:text"`        // Encrypted sensitive fields
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
```

**Encryption Flow:**

```
Save(ctx, config):
  1. Register struct with kx.Registry (define which fields to encrypt)
  2. Call kxManager.EncryptStruct(ctx, &config, encryptionContext)
     - Returns: encryptedObj (sensitive fields cleared), ciphertext
  3. Marshal encryptedObj to JSON
  4. Save to DB: { key, data: json, ciphertext }

Load(ctx):
  1. Load record from DB
  2. Unmarshal data to config struct
  3. Call kxManager.DecryptStruct(ctx, &config, ciphertext, encryptionContext)
  4. Return decrypted config
```

### 4. API Layer - Config Management

HTTP handlers for config CRUD operations.

```go
type ConfigAPI[T any] struct {
    dynamicConfig *DynamicConfig[T]
    store         ConfigStore[T]
}

// Endpoints:
// GET  /api/config         - Get current config (from cache)
// PUT  /api/config         - Update config (save to DB, invalidate cache)
// POST /api/config/reload  - Force reload config
```

**Update Flow:**

```
PUT /api/config:
  1. Parse request body to config struct
  2. Validate config
  3. Call store.Save(ctx, config)
  4. Call dynamicConfig.Invalidate()
  5. Return success
```

## Usage Example

### 1. Define Config Struct

```go
type AppConfig struct {
    ServerPort    int    `confx:"serverPort" json:"serverPort"`
    DatabaseURL   string `confx:"databaseURL" json:"databaseUrl"`

    // Sensitive fields - will be encrypted
    OAuth2ClientID     string `confx:"oauth2ClientID" json:"oauth2ClientId"`
    OAuth2ClientSecret string `confx:"oauth2ClientSecret" json:"oauth2ClientSecret"`
}
```

### 2. Setup kx Registry

```go
registry, _ := kx.NewRegistry()
registry.MustRegisterStruct(&AppConfig{},
    kx.WithRegularField("OAuth2ClientSecret", false), // Encrypt, no hash
)
```

### 3. Create Store and DynamicConfig

```go
// Create kx manager
kxManager, _ := kx.NewManagerByConfig(&kx.Config{
    KMSKeyID: "arn:aws:kms:...",
    HashKey:  "base64-encoded-key",
}, registry)

// Create encrypted store
store := confx.NewEncryptedConfigStore[*AppConfig](
    db,
    kxManager,
    "app_config",                           // config key
    map[string]string{"app": "myapp"},      // encryption context
)

// Create dynamic config with 5 min TTL
dynamicConfig := confx.NewDynamicConfig[*AppConfig](
    store,
    confx.WithTTL[*AppConfig](5 * time.Minute),
    confx.WithValidator[*AppConfig](validator.New()),
)

// Use in application
cfg, err := dynamicConfig.Get(ctx)
```

### 4. Setup API

```go
configAPI := confx.NewConfigAPI(dynamicConfig, store)

r := chi.NewRouter()
r.Get("/api/config", configAPI.GetConfig)
r.Put("/api/config", configAPI.UpdateConfig)
r.Post("/api/config/reload", configAPI.ReloadConfig)
```

## DB Schema

```sql
CREATE TABLE configs (
    key         VARCHAR(255) PRIMARY KEY,
    data        JSONB NOT NULL,
    ciphertext  TEXT,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_configs_updated_at ON configs(updated_at);
```

## Cache Strategy

| Scenario | TTL | Behavior |
|----------|-----|----------|
| Strong consistency required | 0 | Always query DB |
| Normal use case | 1-5 min | Cache with TTL, auto refresh on expiry |
| Low frequency changes | 10-30 min | Longer cache, rely on `Invalidate()` for immediate updates |

**TTL as Fallback:**
- TTL ensures eventual consistency even if `Invalidate()` is not called
- Prevents stale data from persisting indefinitely
- Reduces DB load for frequently accessed configs

## Security Considerations

1. **Encryption Context**: Use meaningful context (e.g., `{"app": "myapp", "env": "prod"}`) to bind ciphertext to specific use case
2. **Key Rotation**: kx/awskms supports key rotation via AWS KMS
3. **Audit Logging**: Log config changes with user info and timestamp
4. **Access Control**: Implement proper authorization for config API

## Error Handling

| Error | Handling |
|-------|----------|
| DB connection failure | Return error, keep serving stale cache if available |
| Decryption failure | Return error, log for investigation |
| Validation failure | Return error, do not update cache |
| KMS unavailable | Return error, consider fallback strategy |

## Additional Considerations

### 1. Concurrency Safety

- Multiple goroutines may call `Get()` simultaneously when cache expires
- Use `sync.RWMutex` for read-heavy workloads
- Consider `singleflight` to prevent thundering herd on cache miss

```go
import "golang.org/x/sync/singleflight"

type DynamicConfig[T any] struct {
    // ...
    sf singleflight.Group
}

func (c *DynamicConfig[T]) load(ctx context.Context) (T, error) {
    result, err, _ := c.sf.Do("load", func() (interface{}, error) {
        return c.store.Load(ctx)
    })
    return result.(T), err
}
```

### 2. Config Not Found (First Load)

- What if config doesn't exist in DB on first `Get()`?
- Options:
  - Return error
  - Return zero value
  - Use default value provided at creation

```go
func WithDefault[T any](defaultValue T) DynamicConfigOption[T]
```

### 4. Config Migration / Schema Evolution

- Config struct may change over time (add/remove fields)
- JSON unmarshaling handles missing fields gracefully
- Consider versioning for breaking changes

```go
type ConfigRecord struct {
    Key        string
    Version    int       // Schema version
    Data       []byte
    Ciphertext string
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
```

### 4. Multiple Config Keys

- Application may have multiple independent configs
- Each `DynamicConfig[T]` instance handles one config key
- Consider a registry pattern for managing multiple configs

### 5. Initialization Order

- Config may be needed before DB connection is ready
- Consider lazy initialization or explicit `Init()` method

```go
func (c *DynamicConfig[T]) Init(ctx context.Context) error {
    _, err := c.Reload(ctx)
    return err
}
```

### 6. Partial Update

- Current design replaces entire config
- For large configs, may want partial update API
- Trade-off: complexity vs. flexibility

### 7. Testing Support

- Provide mock store for unit tests
- Allow injecting fixed config for testing

```go
// For testing
func NewFixedConfig[T any](value T) *DynamicConfig[T]
```

## File Structure

```
confx/
├── confx.go                # Existing static config initialization
├── ...
└── dynamic/                # New package for dynamic config
    ├── dynconfx.go          # DynamicConfig[T] implementation
    ├── dynamic_test.go
    ├── store.go            # ConfigStore interface
    ├── store_encrypted.go  # EncryptedConfigStore implementation
    ├── store_encrypted_test.go
    ├── api.go              # HTTP handlers (optional)
    └── api_test.go
```

## Implementation Priority

1. **Phase 1**: `DynamicConfig[T]` with simple loader function (no encryption)
2. **Phase 2**: `ConfigStore` interface and `EncryptedConfigStore` with kx integration
3. **Phase 3**: HTTP API layer
4. **Phase 4**: Documentation and examples
