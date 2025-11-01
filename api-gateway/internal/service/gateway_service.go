package service

import (
	"api-gateway/config"
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type GatewayService struct {
	cfg        *config.Config
	httpClient *http.Client
}

func NewGatewayService(cfg *config.Config) *GatewayService {
	return &GatewayService{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.ProxyTimeoutSec) * time.Second,
		},
	}
}

func (s *GatewayService) DoForward(ctx context.Context, method, targetBase, path string, headers http.Header, body []byte) (int, http.Header, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, targetBase+path, bytes.NewReader(body))
	if err != nil {
		return 0, nil, nil, err
	}
	// copy headers
	for k, vv := range headers {
		for _, v := range vv {
			req.Header.Add(k, v)
		}
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, resp.Header, b, nil
}
