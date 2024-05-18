package config

import (
	"flag"
	"fmt"
	"os"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway microservice config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	Logger      loggerClient.Config `mapstructure:"logger"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	Http        Http                `mapstructure:"http"`
	GrpcClient  grpcServer.Config   `mapstructure:"grpcClient"`
	Kafka       kafkaClient.Config  `mapstructure:"kafka"`
}

type Http struct {
	Port                string   `mapstructure:"port"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath"`
	KreditPath          string   `mapstructure:"kreditPath"`
	DebugHeaders        bool     `mapstructure:"debugHeaders"`
	HttpClientDebug     bool     `mapstructure:"httpClientDebug"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
}

type KafkaTopics struct {
	KonsumenCreate  kafkaClient.TopicConfig `mapstructure:"konsumenCreate"`
	LimitCreate     kafkaClient.TopicConfig `mapstructure:"limitCreate"`
	TransaksiCreate kafkaClient.TopicConfig `mapstructure:"transaksiCreate"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/api-gateway-service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort != "" {
		cfg.Http.Port = httpPort
	}

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	readerServicePort := os.Getenv("READER_SERVICE")
	if readerServicePort != "" {
		cfg.GrpcClient.Port = readerServicePort
	}

	return cfg, nil
}
