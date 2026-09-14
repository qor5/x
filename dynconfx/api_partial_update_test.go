package dynconfx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigAPI_PartialUpdateConfig(t *testing.T) {
	type TestConfig struct {
		ServerPort  int    `json:"serverPort"`
		DatabaseURL string `json:"databaseUrl"`
		RedisAddr   string `json:"redisAddr"`
	}

	t.Run("partial update single field", func(t *testing.T) {
		var currentConfig = TestConfig{
			ServerPort:  8080,
			DatabaseURL: "postgres://localhost/db1",
			RedisAddr:   "localhost:6379",
		}

		store := &mockStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				return currentConfig, nil
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				currentConfig = config
				return nil
			},
		}

		dc := NewConfigProvider[TestConfig](store)
		api := NewConfigAPI[TestConfig](dc, store)

		// Partial update - only change server port
		body := `{"serverPort": 9090}`
		req := httptest.NewRequest(http.MethodPatch, "/api/config", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		api.PartialUpdateConfig(w, req)

		// Check response
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.True(t, response["success"].(bool))

		// Verify only server port changed
		assert.Equal(t, 9090, currentConfig.ServerPort)
		assert.Equal(t, "postgres://localhost/db1", currentConfig.DatabaseURL)
		assert.Equal(t, "localhost:6379", currentConfig.RedisAddr)
	})

	t.Run("partial update multiple fields", func(t *testing.T) {
		var currentConfig TestConfig = TestConfig{
			ServerPort:  8080,
			DatabaseURL: "postgres://localhost/db1",
			RedisAddr:   "localhost:6379",
		}

		store := &mockStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				return currentConfig, nil
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				currentConfig = config
				return nil
			},
		}

		dc := NewConfigProvider[TestConfig](store)
		api := NewConfigAPI[TestConfig](dc, store)

		// Partial update - change multiple fields
		body := `{"serverPort": 9090, "databaseUrl": "postgres://localhost/db2"}`
		req := httptest.NewRequest(http.MethodPatch, "/api/config", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		api.PartialUpdateConfig(w, req)

		// Check response
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.True(t, response["success"].(bool))

		// Verify updated fields
		assert.Equal(t, 9090, currentConfig.ServerPort)
		assert.Equal(t, "postgres://localhost/db2", currentConfig.DatabaseURL)
		assert.Equal(t, "localhost:6379", currentConfig.RedisAddr)
	})

	t.Run("partial update with validation error", func(t *testing.T) {
		store := &mockStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				return TestConfig{
					ServerPort:  8080,
					DatabaseURL: "postgres://localhost/db1",
					RedisAddr:   "localhost:6379",
				}, nil
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				return nil
			},
		}

		validator := &mockValidator{
			validateFunc: func(ctx context.Context, v any) error {
				cfg := v.(TestConfig)
				if cfg.ServerPort < 1024 {
					return errors.New("server port must be >= 1024")
				}
				return nil
			},
		}

		dc := NewConfigProvider[TestConfig](store)
		api := NewConfigAPI[TestConfig](dc, store, WithAPIValidator[TestConfig](validator))

		// Partial update with invalid port
		body := `{"serverPort": 1023}`
		req := httptest.NewRequest(http.MethodPatch, "/api/config", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		api.PartialUpdateConfig(w, req)

		// Check response
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"].(string), "validation failed")
	})

	t.Run("partial update with invalid JSON", func(t *testing.T) {
		store := &mockStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				return TestConfig{
					ServerPort:  8080,
					DatabaseURL: "postgres://localhost/db1",
					RedisAddr:   "localhost:6379",
				}, nil
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				return nil
			},
		}

		dc := NewConfigProvider[TestConfig](store)
		api := NewConfigAPI[TestConfig](dc, store)

		// Invalid JSON
		body := `{"serverPort": invalid}`
		req := httptest.NewRequest(http.MethodPatch, "/api/config", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		api.PartialUpdateConfig(w, req)

		// Check response
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"].(string), "failed to decode request body")
	})
}
