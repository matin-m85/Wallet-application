package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type WalletClient struct {
	baseURL string
	client  *http.Client
}

func NewWalletClient(baseURL string, timeoutSec int) *WalletClient {
	return &WalletClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: time.Duration(timeoutSec) * time.Second,
		},
	}
}

// Forward an HTTP request to wallet service and return response
func (wc *WalletClient) DoForward(ctx context.Context, method, path string, headers http.Header, body []byte) (int, http.Header, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, wc.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return 0, nil, nil, err
	}
	req.Header = headers.Clone()

	resp, err := wc.client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, resp.Header, respBody, nil
}
