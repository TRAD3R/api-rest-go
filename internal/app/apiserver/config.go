package apiserver

import "github.com/trad3r/rest-api-go/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` // для конфигурации используется toml файл
	LogLevel string `toml:"log_level"` // указание уровня логирования для логгера
	Store    *store.Config
}

// NewConfig - установка конфигурации сервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
