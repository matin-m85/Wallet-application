package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type DiscountClient struct {
	baseURL string
	client  *http.Client
}

func NewDiscountClient(baseURL string, timeoutSec int) *DiscountClient {
	return &DiscountClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: time.Duration(timeoutSec) * time.Second,
		},
	}
}

func (dc *DiscountClient) DoForward(ctx context.Context, method, path string, headers http.Header, body []byte) (int, http.Header, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, dc.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return 0, nil, nil, err
	}
	req.Header = headers.Clone()

	resp, err := dc.client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, resp.Header, respBody, nil
}
