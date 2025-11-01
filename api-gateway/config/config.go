package config

import "os"

type Config struct {
	WalletURL       string
	DiscountURL     string
	Port            string
	ProxyTimeoutSec int
}

func LoadConfig() *Config {
	return &Config{
		WalletURL:       getEnv("WALLET_SERVICE_URL", "http://wallet-service:8081"),
		DiscountURL:     getEnv("DISCOUNT_SERVICE_URL", "http://discount-service:8080"),
		Port:            getEnv("PORT", "8082"),
		ProxyTimeoutSec: 10,
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
