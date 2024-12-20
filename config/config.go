package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	TelegramApiKey  string   `json:"telegram_api_key"`
	TelegramChannel string   `json:"telegram_channel"`
	ConsumerKey     string   `json:"consumer_key"`
	ConsumerSecret  string   `json:"consumer_secret"`
	Host            string   `json:"host"`
	PkgName         string   `json:"pkgname"`
	Filters         []Filter `json:"filters"`
}

type Filter struct {
	Include        string  `json:"include"`
	Exclude        string  `json:"exclude"`
	IncludeWebsite string  `json:"include_website"`
	ExcludeWebsite string  `json:"exclude_website"`
	LowestPrice    float64 `json:"lowest_price"`
	MaximumPrice   float64 `json:"maximum_price"`
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
