package model

type Config struct {
	Port     string         `yaml:"port"`
	DBConfig DatabaseConfig `yaml:"db"`
}

const (
	minUsernameLen = 3
	maxUsernameLen = 20
	minPasswordLen = 6
	maxPasswordLen = 100
)
