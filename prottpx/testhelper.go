package prottpx

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Client is a helper for testing client-facing HTTP endpoints
type Client struct {
	*http.Client
	baseURL            string
	authorizationToken string
}

// WrapClient wraps an existing *http.Client with prottpx functionality
func WrapClient(client *http.Client, baseURL string) *Client {
	return &Client{
		Client:  client,
		baseURL: baseURL,
	}
}

// WithAuthorization sets the authorization token for subsequent requests
func (c *Client) WithAuthorization(token string) *Client {
	newClient := *c
	newClient.authorizationToken = token
	return &newClient
}

// Call makes an HTTP POST request to the prottpx endpoint
func (c *Client) Call(ctx context.Context, method string, req proto.Message, reply proto.Message) error {
	reqBody, err := protojson.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	url := c.baseURL + method
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "failed to create HTTP request")
	}

	httpReq.Header.Set(HeaderContentType, JSONContentType)
	httpReq.Header.Set(HeaderAccept, JSONContentType)
	if c.authorizationToken != "" {
		httpReq.Header.Set("Authorization", c.authorizationToken)
	}

	httpResp, err := c.Client.Do(httpReq)
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	if err != nil {
		return errors.Wrap(err, "failed to send HTTP request")
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	if httpResp.StatusCode != http.StatusOK {
		return errors.Errorf("HTTP error: status=%d, body=%s", httpResp.StatusCode, string(respBody))
	}

	if err := protojson.Unmarshal(respBody, reply); err != nil {
		return errors.Wrap(err, "failed to unmarshal response")
	}

	return nil
}
