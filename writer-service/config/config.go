package config

import (
	"flag"
	"fmt"
	"os"

	grpcServer "github.com/0x5w4/kredit-plus/pkg/grpc-server"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	postgresClient "github.com/0x5w4/kredit-plus/pkg/postgres"
	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway microservice config path")
}

type Config struct {
	ServiceName string                `mapstructure:"serviceName"`
	Logger      loggerClient.Config   `mapstructure:"logger"`
	KafkaTopics KafkaTopics           `mapstructure:"kafkaTopics"`
	GrpcServer  grpcServer.Config     `mapstructure:"grpcServer"`
	Kafka       kafkaClient.Config    `mapstructure:"kafka"`
	Postgresql  postgresClient.Config `mapstructure:"postgres"`
}

type KafkaTopics struct {
	KonsumenCreate   kafkaClient.TopicConfig `mapstructure:"konsumenCreate"`
	KonsumenCreated  kafkaClient.TopicConfig `mapstructure:"konsumenCreated"`
	LimitCreate      kafkaClient.TopicConfig `mapstructure:"limitCreate"`
	LimitCreated     kafkaClient.TopicConfig `mapstructure:"limitCreated"`
	TransaksiCreate  kafkaClient.TopicConfig `mapstructure:"transaksiCreate"`
	TransaksiCreated kafkaClient.TopicConfig `mapstructure:"transaksiCreated"`
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
			configPath = fmt.Sprintf("%s/writer-service/config/config.yaml", getwd)
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

	grpcPort := os.Getenv("GRPC_SERVER_PORT")
	if grpcPort != "" {
		cfg.GrpcServer.Port = grpcPort
	}
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}
	//kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	//if kafkaBrokers != "" {
	//	cfg.Kafka.Brokers = []string{"host.docker.internal:9092"}
	//}
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
