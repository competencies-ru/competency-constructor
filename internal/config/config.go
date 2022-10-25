package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP HTTP
	}

	HTTP struct {
		Port            string
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		ShutdownTimeout time.Duration
	}
)

const (
	defaultHTTTPort             = "8080"
	defaultHTTPReadWriteTimeout = 10 * time.Second
	defaultHTTPShutdown         = 10 * time.Second
)

func setDefaults() {
	viper.SetDefault("http.port", defaultHTTTPort)
	viper.SetDefault("http.readTimeout", defaultHTTPReadWriteTimeout)
	viper.SetDefault("http.writeTimeout", defaultHTTPReadWriteTimeout)
	viper.SetDefault("http.shutdownTimeout", defaultHTTPShutdown)
}

func Parse(pathToConfigs string) (*Config, error) {
	setDefaults()

	var cfg Config

	if err := unmarshall(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshall(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	return nil
}
