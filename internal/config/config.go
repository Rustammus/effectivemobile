package config

import (
	"EffectiveMobile/pkg/logging"
	"encoding/json"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	Server struct {
		Port         string `yaml:"port" env:"APP_PORT"`
		ExternalURL  string `yaml:"external_url" env:"APP_EXTERNAL_URL"`
		MigrationDir string `yaml:"migration_dir" env:"APP_MIGRATION_DIR"`
		MaxRowLimit  int    `yaml:"max_row_limit" env:"APP_MAX_ROW_LIMIT"`
		LogLevel     string `yaml:"log_level" env:"APP_LOG_LEVEL"`
	} `yaml:"server"`
	Storage `yaml:"storage"`
}

type Storage struct {
	Username string `yaml:"username" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     string `yaml:"port" env:"POSTGRES_PORT"`
	Database string `yaml:"database" env:"POSTGRES_DATABASE"`
}

var once sync.Once
var instance *Config

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		confPath := ".env"
		if value, ok := os.LookupEnv("CONFIG_FILE"); ok {
			confPath = value
		} else {
			logger.Info("env: CONFIG_FILE not set. Default: \".env\"")
		}

		logger.Infof("Try to read config file %s", confPath)
		instance = &Config{}
		err := cleanenv.ReadConfig(confPath, instance)
		if err != nil {
			logger.Fatal("Failed to read config file. Abort start app.\n\t", err.Error())
		}
		configJSON, _ := json.Marshal(&instance)
		logger.Info("Config file read. Config: ", string(configJSON))
	})
	return instance
}
