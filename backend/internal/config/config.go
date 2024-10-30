package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBConfig struct {
	DBHost     string `yaml:"db_host" env-required:"true"`
	DBPort     string `yaml:"db_port" env-required:"true"`
	DBUser     string `yaml:"db_user" env-required:"true"`
	DBPassword string `yaml:"db_pass" env-required:"true"`
	DBName     string `yaml:"db_name" env-required:"true"`
}

type Config struct {
	Env      string `yaml:"env" env-default:"production"`
	Port     string `yaml:"port" env-required:"true"`
	DBConfig `yaml:"db_config"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists on the given path: %s", configPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}
	return &cfg
}
