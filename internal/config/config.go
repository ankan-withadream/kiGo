package config

import "time"

type Config struct {
	API_PORT               int
	AI_CLIENT_API_ENDPOINT string
	ReadTimeout            time.Duration
	WriteTimeout           time.Duration
}

var APP_CONFIG = Config{
	API_PORT:               8080,
	AI_CLIENT_API_ENDPOINT: "https://www.freeaitherapist.com/api/message",
	ReadTimeout:            10 * time.Second,
	WriteTimeout:           10 * time.Second,
}
