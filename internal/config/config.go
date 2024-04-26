package config

type Config struct {
	API_PORT               int
	AI_CLIENT_API_ENDPOINT string
}

var APP_CONFIG = Config{
	API_PORT:               8080,
	AI_CLIENT_API_ENDPOINT: "https://www.freeaitherapist.com/api/message",
}
