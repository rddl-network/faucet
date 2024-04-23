package config

import "sync"

const DefaultConfigTemplate = `
address="{{ .Address }}"
service-bind="{{ .ServiceBind }}"
service-port={{ .ServicePort }}
`

type Config struct {
	Address     string `mapstructure:"address"`
	ServiceBind string `mapstructure:"service-bind"`
	ServicePort int    `mapstructure:"service-port"`
}

// global singleton
var (
	config     *Config
	initConfig sync.Once
)

// DefaultConfig returns the default config.
func DefaultConfig() *Config {
	return &Config{
		Address:     "plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty",
		ServiceBind: "localhost",
		ServicePort: 8080,
	}
}

// GetConfig returns the current config.
func GetConfig() *Config {
	initConfig.Do(func() {
		config = DefaultConfig()
	})
	return config
}
