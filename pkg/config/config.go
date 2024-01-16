package config

type Config struct {
	PostgresURL         string
	ComissionOfTransfer float64
}

func NewConfig() *Config {
	return &Config{
		PostgresURL:         "postgres://postgres:admin@localhost:5432/bank_cli",
		ComissionOfTransfer: 5.0,
	}
}
