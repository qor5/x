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

func TestConfigAPI_GetConfig(t *testing.T) {
	type Config struct {
		Port    int
		Host    string
		Enabled bool
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{
				Port:    8080,
				Host:    "localhost",
				Enabled: true,
			}, nil
		},
	}

	cp := NewConfigProvider[Config](store)
	api := NewConfigAPI[Config](cp, store)

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/api/config", nil)
	w := httptest.NewRecorder()

	// Call handler
	api.GetConfig(w, req)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)

	var config Config
	err := json.Unmarshal(w.Body.Bytes(), &config)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "localhost", config.Host)
	assert.True(t, config.Enabled)
}

func TestConfigAPI_UpdateConfig(t *testing.T) {
	type Config struct {
		Port    int    `json:"port"`
		Host    string `json:"host"`
		Enabled bool   `json:"enabled"`
	}

	store := &mockStore[Config]{
		saveFunc: func(ctx context.Context, config Config) error {
			return nil
		},
	}

	cp := NewConfigProvider[Config](store)
	api := NewConfigAPI[Config](cp, store)

	// Test with empty body - should fail
	req := httptest.NewRequest(http.MethodPut, "/api/config", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	api.UpdateConfig(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test with valid body
	bodyStr := `{"port":9090,"host":"example.com","enabled":false}`
	req = httptest.NewRequest(http.MethodPut, "/api/config", strings.NewReader(bodyStr))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	api.UpdateConfig(w, req)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "Config updated successfully", response["message"])

	// Verify cache was invalidated
	assert.Equal(t, 1, store.saveCount)
}

func TestConfigAPI_UpdateConfig_ValidationError(t *testing.T) {
	type Config struct {
		Port int    `json:"port" validate:"required,min=1,max=65535"`
		Host string `json:"host" validate:"required"`
	}

	store := &mockStore[Config]{}

	validator := &mockValidator{
		validateFunc: func(ctx context.Context, v any) error {
			return errors.New("port must be between 1 and 65535")
		},
	}

	cp := NewConfigProvider[Config](store)
	api := NewConfigAPI[Config](cp, store, WithAPIValidator[Config](validator))

	// Create request with invalid config
	body := `{"port":0,"host":""}`
	req := httptest.NewRequest(http.MethodPut, "/api/config", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	api.UpdateConfig(w, req)

	// Check response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Contains(t, response["error"], "validation failed")
}

func TestConfigAPI_ReloadConfig(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{Port: 8080}, nil
		},
	}

	cp := NewConfigProvider[Config](store)
	api := NewConfigAPI[Config](cp, store)

	// Create test request
	req := httptest.NewRequest(http.MethodPost, "/api/config/reload", nil)
	w := httptest.NewRecorder()

	// Call handler
	api.ReloadConfig(w, req)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "Config reloaded successfully", response["message"])

	// Check config in response
	config := response["config"].(map[string]interface{})
	assert.Equal(t, float64(8080), config["Port"])

	// Verify store was called
	assert.Equal(t, 1, store.loadCount)
}

func TestConfigAPI_GetConfig_Error(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			var zero Config
			return zero, errors.New("database connection failed")
		},
	}

	dc := NewConfigProvider[Config](store)
	api := NewConfigAPI[Config](dc, store)

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/api/config", nil)
	w := httptest.NewRecorder()

	// Call handler
	api.GetConfig(w, req)

	// Check response
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Contains(t, response["error"].(string), "failed to get config")
}
