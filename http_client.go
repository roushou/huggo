package huggo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const DefaultAPIBaseURL = "https://huggingface.co/api"

type options struct {
	apiKey  string
	baseURL string
}

// Option defines a function that can customize the Client.
type Option func(options *options) error

// WithAPIKey returns an Option that sets the api key for API requests.
func WithAPIKey(apiKey string) Option {
	return func(options *options) error {
		if apiKey == "" {
			return errors.New("api key should not be empty")
		}
		options.apiKey = apiKey
		return nil
	}
}

// WithBaseURL returns an Option that sets the base URL for API requests.
func WithBaseURL(baseURL string) Option {
	return func(options *options) error {
		if baseURL == "" {
			return errors.New("base URL should not be empty")
		}
		options.baseURL = baseURL
		return nil
	}
}

// HttpClient represents a client for interacting with the HuggingFace API.
type HttpClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewHttpClient creates a new HTTP client with default settings and optional configurations.
func NewHttpClient(apiKey string, opts ...Option) (*HttpClient, error) {
	options := &options{
		apiKey:  apiKey,
		baseURL: DefaultAPIBaseURL,
	}
	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}

	return &HttpClient{
		apiKey:     options.apiKey,
		baseURL:    options.baseURL,
		httpClient: http.DefaultClient,
	}, nil
}

// newRequest constructs a new HTTP request.
func (c *HttpClient) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	url := c.baseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	return req, nil
}

// doRequest sends an HTTP request and decodes the response into the provided interface.
func (c *HttpClient) doRequest(req *http.Request, out interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed (status: %s, body: %s)", resp.Status, string(body))

	}
	return json.NewDecoder(resp.Body).Decode(out)

}

// Get sends a GET request.
func (c *HttpClient) Get(path string, out interface{}) error {
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("GET request failed: %v", err)
	}
	return c.doRequest(req, out)
}

// Post sends a POST request.
func (c *HttpClient) Post(path string, payload interface{}, out interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to serialize body: %v", err)
	}
	req, err := c.newRequest(http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	return c.doRequest(req, out)
}

// Put sends a PUT request.
func (c *HttpClient) Put(path string, payload interface{}, out interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to serialize body: %v", err)
	}
	req, err := c.newRequest(http.MethodPut, path, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	return c.doRequest(req, out)
}

// Delete sends a DELETE request.
func (c *HttpClient) Delete(path string, payload interface{}, out interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to serialize body: %v", err)
	}
	req, err := c.newRequest(http.MethodDelete, path, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	return c.doRequest(req, out)
}
