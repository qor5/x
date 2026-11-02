package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/qor5/x/v3/dynconfx"
)

// AppConfig represents application configuration
type AppConfig struct {
	ServerPort  int    `json:"serverPort"`
	DatabaseURL string `json:"databaseUrl"`
	RedisAddr   string `json:"redisAddr"`

	// Sensitive fields - will be encrypted
	OAuth2ClientID     string `json:"oauth2ClientId"`
	OAuth2ClientSecret string `json:"oauth2ClientSecret"`
}

func main() {
	ctx := context.Background()

	// Note: This is a simplified example using a mock store.
	// In a real application with encryption, you would:
	//
	// 1. Create kx Registry and register config struct:
	//    import "github.com/qor5/kx"
	//
	//    registry, _ := kx.NewRegistry()
	//    registry.MustRegisterStruct(&AppConfig{},
	//        kx.WithRegularField("OAuth2ClientSecret", false),
	//    )
	//
	// 2. Create kx Manager:
	//    kxManager, _ := kx.NewManagerByConfig(&kx.Config{
	//        KMSKeyID: "arn:aws:kms:...",
	//        HashKey:  "base64-encoded-key",
	//    }, registry)
	//
	// 3. Create EncryptedConfigStore:
	//    store := dynconfx.NewEncryptedConfigStore[*AppConfig](
	//        db,
	//        kxManager,
	//        "app_config",
	//        map[string]string{"app": "myapp"},
	//    )

	// Mock store for demonstration
	store := &mockStore{}

	// Create DynamicConfig with 5 minute TTL
	dynamicConfig := dynconfx.NewConfigProvider[*AppConfig](
		store,
		dynconfx.WithTTL[*AppConfig](5*time.Minute),
	)

	// Initialize config by loading it
	config, err := dynamicConfig.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Loaded config: %+v", config)

	// Create HTTP API
	configAPI := dynconfx.NewConfigAPI[*AppConfig](dynamicConfig, store)

	// Set up HTTP routes (stdlib mux, method-based patterns)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/config", configAPI.GetConfig)
	mux.HandleFunc("PUT /api/config", configAPI.UpdateConfig)
	mux.HandleFunc("POST /api/config/reload", configAPI.ReloadConfig)

	// Start HTTP server
	log.Println("Starting HTTP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// mockStore is a mock implementation of ConfigStore for demonstration
type mockStore struct{}

func (m *mockStore) Load(ctx context.Context) (*AppConfig, error) {
	return &AppConfig{
		ServerPort:         8080,
		DatabaseURL:        "postgres://localhost/mydb",
		RedisAddr:          "localhost:6379",
		OAuth2ClientID:     "my-client-id",
		OAuth2ClientSecret: "my-secret",
	}, nil
}

func (m *mockStore) Save(ctx context.Context, config *AppConfig) error {
	log.Printf("Saving config: %+v", config)
	return nil
}

func (m *mockStore) Update(ctx context.Context, updateFunc func(*AppConfig) *AppConfig) error {
	config, err := m.Load(ctx)
	if err != nil {
		return err
	}
	updated := updateFunc(config)
	return m.Save(ctx, updated)
}
