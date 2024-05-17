package config

import (
	"flag"
	"fmt"
	"os"

	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	mongoClient "github.com/0x5w4/kredit-plus/pkg/mongo"
	postgresClient "github.com/0x5w4/kredit-plus/pkg/postgres"
	redisClient "github.com/0x5w4/kredit-plus/pkg/redis"
	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway microservice config path")
}

type Config struct {
	ServiceName      string                `mapstructure:"serviceName"`
	Logger           loggerClient.Config   `mapstructure:"logger"`
	KafkaTopics      KafkaTopics           `mapstructure:"kafkaTopics"`
	Grpc             Grpc                  `mapstructure:"grpc"`
	Kafka            kafkaClient.Config    `mapstructure:"kafka"`
	Postgresql       postgresClient.Config `mapstructure:"postgres"`
	Mongo            mongoClient.Config    `mapstructure:"mongo"`
	Redis            redisClient.Config    `mapstructure:"redis"`
	MongoCollections MongoCollections      `mapstructure:"mongoCollections"`
	ServiceSettings  ServiceSettings       `mapstructure:"serviceSettings"`
}

type Grpc struct {
	Port string `mapstructure:"port"`
}

type MongoCollections struct {
	Kredit string `mapstructure:"kredit"`
}
type ServiceSettings struct {
	RedisProductPrefixKey string `mapstructure:"redisKreditPrefixKey"`
}

type KafkaTopics struct {
	ProductCreated kafkaClient.TopicConfig `mapstructure:"productCreated"`
	ProductUpdated kafkaClient.TopicConfig `mapstructure:"productUpdated"`
	ProductDeleted kafkaClient.TopicConfig `mapstructure:"productDeleted"`
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
			configPath = fmt.Sprintf("%s/reader-service/config/config.yaml", getwd)
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

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort != "" {
		cfg.Grpc.Port = grpcPort
	}
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI != "" {
		//cfg.Mongo.URI = "mongodb://host.docker.internal:27017"
		cfg.Mongo.URI = mongoURI
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr != "" {
		cfg.Redis.Addr = redisAddr
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
