package config

import (
	"fmt"
	"log"
	"os"

	"github.com/pruser/allegro-feed-generator/utils"

	"github.com/pruser/allegro-feed-generator/model"
)

const (
	ENV_ALLEGRO_WEB_API_KEY = "ALLEGRO_WEB_API_KEY"
	ENV_ALLEGRO_URL_BASE    = "ALLEGRO_URL_BASE"
)

type Config struct {
	WebAPIKey model.WebAPIKey
	UrlBase   string
}

func ReadConfig() (Config, error) {
	webApiKey := model.WebAPIKey(os.Getenv(ENV_ALLEGRO_WEB_API_KEY))
	if webApiKey == "" {
		return Config{}, fmt.Errorf("missing environment variable '%s'", ENV_ALLEGRO_WEB_API_KEY)
	}
	urlBase := os.Getenv(ENV_ALLEGRO_URL_BASE)
	if urlBase == "" {
		return Config{}, fmt.Errorf("missing environment variable '%s'", ENV_ALLEGRO_URL_BASE)
	}

	return Config{webApiKey, urlBase}, nil
}

func (c Config) Print() {
	log.Printf("WebAPIKey: %s", utils.Cover(string(c.WebAPIKey), 6))
	log.Printf("URLBase: %s", c.UrlBase)
}
