package postgres

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
}
