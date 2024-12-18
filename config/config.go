package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	TelegramApiKey  string `json:"telegram_api_key"`
	TelegramChannel string `json:"telegram_channel"`
	ConsumerKey     string `json:"consumer_key"`
	ConsumerSecret  string `json:"consumer_secret"`
	Host            string `json:"host"`
	PkgName         string `json:"pkgname"`
}

func ReadConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Fatal error opening config file: %s", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Fatal error reading config file: %s", err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatalf("Fatal error parsing config file: %s", err)
	}

	if config.TelegramChannel == "" || config.TelegramApiKey == "" ||
		config.ConsumerKey == "" || config.ConsumerSecret == "" ||
		config.Host == "" || config.PkgName == "" {
		log.Fatal("Empty values in config file")
	}

	return config
}
