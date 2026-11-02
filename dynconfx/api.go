package dynconfx

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// ConfigAPI provides HTTP handlers for config management operations.
type ConfigAPI[T any] struct {
	dynamicConfig *ConfigProvider[T]
	store         ConfigStore[T]
	validator     Validator
}

// NewConfigAPI creates a new ConfigAPI instance.
func NewConfigAPI[T any](provider *ConfigProvider[T], store ConfigStore[T], opts ...ConfigAPIOption[T]) *ConfigAPI[T] {
	api := &ConfigAPI[T]{
		dynamicConfig: provider,
		store:         store,
	}
	for _, opt := range opts {
		opt(api)
	}
	return api
}

// ConfigAPIOption is a function that configures a ConfigAPI.
type ConfigAPIOption[T any] func(*ConfigAPI[T])

// WithAPIValidator sets a validator for the API (separate from ConfigProvider validator).
func WithAPIValidator[T any](v Validator) ConfigAPIOption[T] {
	return func(api *ConfigAPI[T]) {
		api.validator = v
	}
}

// GetConfig handles GET requests to retrieve the current config.
// It returns the cached config if available.
//
// Example:
//
//	r.Get("/api/config", configAPI.GetConfig)
func (api *ConfigAPI[T]) GetConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	config, err := api.dynamicConfig.Get(ctx)
	if err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to get config"))
		return
	}

	respondJSON(w, http.StatusOK, config)
}

// UpdateConfig handles PUT requests to update the config.
// It validates the config, saves it to storage, and invalidates the cache.
//
// Example:
//
//	r.Put("/api/config", configAPI.UpdateConfig)
func (api *ConfigAPI[T]) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Decode request body
	var config T
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		respondError(w, http.StatusBadRequest, errors.Wrap(err, "failed to decode request body"))
		return
	}

	// Validate config if validator is set
	if api.validator != nil {
		if err := api.validator.StructCtx(ctx, config); err != nil {
			respondError(w, http.StatusBadRequest, errors.Wrap(err, "validation failed"))
			return
		}
	}

	// Save config to storage
	if err := api.store.Save(ctx, config); err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to save config"))
		return
	}

	// Invalidate cache so next Get loads the new config
	api.dynamicConfig.Invalidate()

	// Return success response
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Config updated successfully",
	})
}

// ReloadConfig handles POST requests to force reload the config.
// It forces a reload from storage and updates the cache.
//
// Example:
//
//	r.Post("/api/config/reload", configAPI.ReloadConfig)
func (api *ConfigAPI[T]) ReloadConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	config, err := api.dynamicConfig.Reload(ctx)
	if err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to reload config"))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Config reloaded successfully",
		"config":  config,
	})
}

// PartialUpdateConfig handles PATCH requests to update specific fields.
// It loads the current config, applies partial updates, and saves it.
//
// Request body should be a JSON object with fields to update.
// Example: {"server_port": 9090}
//
// Example:
//
//	r.Patch("/api/config", configAPI.PartialUpdateConfig)
func (api *ConfigAPI[T]) PartialUpdateConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Decode request body to map[string]any (partial update)
	var updates map[string]any
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		respondError(w, http.StatusBadRequest, errors.Wrap(err, "failed to decode request body"))
		return
	}

	// Load current config
	config, err := api.store.Load(ctx)
	if err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to load config"))
		return
	}

	// Convert config to map and apply updates
	configJSON, err := json.Marshal(config)
	if err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to marshal config"))
		return
	}

	var configMap map[string]any
	if err := json.Unmarshal(configJSON, &configMap); err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to unmarshal config"))
		return
	}

	// Apply updates to config map
	for key, value := range updates {
		configMap[key] = value
	}

	// Convert back to config struct
	updatedJSON, err := json.Marshal(configMap)
	if err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to marshal updated config"))
		return
	}

	var updatedConfig T
	if err := json.Unmarshal(updatedJSON, &updatedConfig); err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to unmarshal updated config"))
		return
	}

	// Validate config if validator is set
	if api.validator != nil {
		if err := api.validator.StructCtx(ctx, updatedConfig); err != nil {
			respondError(w, http.StatusBadRequest, errors.Wrap(err, "validation failed"))
			return
		}
	}

	// Save updated config
	if err := api.store.Save(ctx, updatedConfig); err != nil {
		respondError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to save config"))
		return
	}

	// Invalidate cache so next Get loads the new config
	api.dynamicConfig.Invalidate()

	// Return success response
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Config updated successfully",
	})
}

// respondJSON writes a JSON response with the given status code.
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// respondError writes an error response with the given status code.
func respondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
