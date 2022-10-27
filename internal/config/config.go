package config

import (
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/hashicorp/go-multierror"

	"github.com/spf13/viper"
)

const LocalEnv = "local"

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

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = LocalEnv
	}

	if err := parseConfig(pathToConfigs, appEnv); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshall(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfig(path string, env string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == LocalEnv {
		return replaceEnvConfig()
	}

	viper.SetConfigName(env)

	if err := viper.MergeInConfig(); err != nil {
		return err
	}

	return replaceEnvConfig()
}

func unmarshall(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	return nil
}

func replaceEnvConfig() error {
	var result error

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			// обрезаем подстроку ${}
			envValut, err := getValueEnv(value[len("${") : len(value)-len("}")])
			result = multierror.Append(result, err)

			value = envValut
		}

		viper.Set(k, value)
	}

	return nil
}

func getValueEnv(key string) (string, error) {
	key, value, hasDef := parsEnv(key)

	envValue, ok := os.LookupEnv(key)
	if ok {
		return envValue, nil
	}

	if hasDef {
		return value, nil
	}

	return "", errors.Errorf("no %s env", key)
}

func parsEnv(env string) (key, value string, hasDef bool) {
	s := strings.SplitN(env, ":", 2)
	key = s[0]

	if len(s) == 2 {
		value = s[1]
		hasDef = true
	}

	return
}
