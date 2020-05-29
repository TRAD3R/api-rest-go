package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"` // для конфигурации используется toml файл
	LogLevel    string `toml:"log_level"` // указание уровня логирования для логгера
	DatabaseUrl string `toml:"database_url"`
}

// NewConfig - установка конфигурации сервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
