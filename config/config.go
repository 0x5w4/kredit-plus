package config

import (
	"os"

	"github.com/0x5w4/kredit-plus/pkg/logger"
)

type Config struct {
	ServiceName string
	Version     string
	Environment string
	RpcServer   RpcServer
	HttpServer  HttpServer
	Logger      logger.Config
}

type Logger struct {
	Level   string
	Encoder string
}

type RpcServer struct {
	Network string
	Port    string
	Tls     bool
}

type HttpServer struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		ServiceName: getEnv("SERVICE_NAME", "kredit-plus"),
		Version:     getEnv("VERSION", "v1.0.0"),
		Environment: getEnv("ENVIRONMENT", "development"),
		RpcServer: RpcServer{
			Network: getEnv("RPC_SERVER_NETWORK", "tcp"),
			Port:    getEnv("RPC_SERVER_PORT", "8001"),
		},
		HttpServer: HttpServer{
			Port: getEnv("HTTP_SERVER_PORT", "5001"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
