package config

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "postgres://postgres:3302@localhost:5432/postgres",
	}
}
