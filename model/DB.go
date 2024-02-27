package model

type DatabaseConfig struct {
	Host     string `yaml:"db.host"`
	Port     string `yaml:"db.port"`
	Username string `yaml:"db.username"`
	Password string `yaml:"db.password"`
	DBName   string `yaml:"db.dbname"`
	SSLMode  string `yaml:"db.sslmode"`
}
