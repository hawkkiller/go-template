package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"template/app/pkg/logging"
)

type Config struct {
	Listen struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	}
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Database    string `yaml:"database"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxAttempts int8   `yaml:"maxAttempts"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
