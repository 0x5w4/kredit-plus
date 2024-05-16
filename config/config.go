package config

import (
	"os"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	"github.com/0x5w4/kredit-plus/pkg/logger"
)

type Config struct {
	ServiceName string
	Version     string
	Environment string
	GrpcServer  grpcServer.Config
	HttpServer  HttpServer
	Logger      logger.Config
}

type HttpServer struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		ServiceName: getEnv("SERVICE_NAME", "kredit-plus"),
		Version:     getEnv("VERSION", "v1.0.0"),
		Environment: getEnv("ENVIRONMENT", "development"),
		GrpcServer: grpcServer.Config{
			Network: getEnv("RPC_SERVER_NETWORK", "tcp"),
			Port:    getEnv("RPC_SERVER_PORT", "8001"),
			Tls:     false,
		},
		HttpServer: HttpServer{
			Port: getEnv("HTTP_SERVER_PORT", "5001"),
		},
		Logger: logger.Config{
			Encoding:   getEnv("LOGGER_ENCODING", "json"),
			Level:      getEnv("LOGGER_LEVEL", "json"),
			OutputPath: getEnv("LOGGER_OUPUT_PATH", "./logs"),
			ErrorPath:  getEnv("LOGGER_ERROR_PATH", "./logs"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
