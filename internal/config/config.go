package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"127.0.0.1:3000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {

	// Load .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading environment variable: ", err)
	}

	// Get config path from .env
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		log.Fatal("Config path is not set")
	}

	// Load config
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		log.Fatal("Config file does not exist:", config_path)
	}

	var cfg Config

	// Init config
	if err := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), &cfg); err != nil {
		log.Fatal("Error loading config: ", err)
	}

	return &cfg

}
