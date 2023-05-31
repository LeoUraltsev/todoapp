package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool `yaml:"is_debug" env:"ISDEBUG" env-default:"true"`
	Listen  struct {
		Port string `yaml:"port" env:"PORT" env-description:"server port" env-default:"10000"`
		Host string `yaml:"host" env:"HOST" env-description:"server host" env-default:"0.0.0.0"`
	} `yaml:"listen"`
	StorageConfig `yaml:"database"`
}

type StorageConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

var instanse *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instanse = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instanse); err != nil {
			help, _ := cleanenv.GetDescription(instanse, nil)
			fmt.Printf("INFO: %s\n", help)
			log.Fatalf("FATAL: %v", err)
		}
	})
	return instanse
}
