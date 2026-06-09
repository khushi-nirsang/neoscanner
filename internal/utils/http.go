package utils

import (
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

// HTTPClient is a reusable HTTP client
type HTTPClient struct {
	Client *http.Client
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(timeout int) *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// Response wraps http.Response with body content
type Response struct {
	*http.Response
	BodyContent string
}

// Get makes a GET request and returns enhanced response
func (h *HTTPClient) Get(url string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "NeoScanner/1.0[](https://github.com/khushi-nirsang/neoscanner)")

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read body
	bodyBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close() // Close original body

	bodyStr := string(bodyBytes)

	color.Green("[+] Connected to %s → Status: %d", url, resp.StatusCode)

	return &Response{
		Response:    resp,
		BodyContent: bodyStr,
	}, nil
}
