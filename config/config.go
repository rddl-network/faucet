package config

import "sync"

const DefaultConfigTemplate = `
address="{{ .Address }}"
amount={{ .Amount }}
denom="{{ .Denom }}"
service-bind="{{ .ServiceBind }}"
service-port={{ .ServicePort }}
`

type Config struct {
	Address     string `mapstructure:"address"`
	Amount      int    `mapstructure:"amount"`
	Denom       string `mapstructure:"denom"`
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
		Amount:      100,
		Denom:       "plmnt",
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
