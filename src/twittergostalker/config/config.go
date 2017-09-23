package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	UserName        string
	PollingInterval time.Duration

	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string

	TelegramBotToken string
}

func Parse(pathToCfg string) *Config {
	data, err := ioutil.ReadFile(pathToCfg)
	if err != nil {
		log.Fatalf("config: could not read %s: %s\n", pathToCfg, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalln("config: could not unmarshal config:", err)
	}

	return &cfg
}
