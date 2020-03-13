package config

type Database struct {
	Driver       string `yaml:"driver"`
	Address      string `yaml:"address"`
	MaxConns     int    `yaml:"max_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}
