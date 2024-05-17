package logger

// LoggerConfig holds the configuration for creating a new logger.
type Config struct {
	Encoding   string `mapstructure:"encoding"`
	Level      string `mapstructure:"level"`
	OutputPath string `mapstructure:"ouputPath"`
	ErrorPath  string `mapstructure:"errorPath"`
}
