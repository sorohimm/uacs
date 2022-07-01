package config

import "os"

type Config struct {
	Port        string
	DBAuthData  DBAuthenticationData
	Collections Collections
}

type DBAuthenticationData struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
	URL      string
}

type Collections struct {
	Competitions string
}

func New() (*Config, error) {
	return &Config{
		Port: os.Getenv("PORT"),
		DBAuthData: DBAuthenticationData{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			URL:      os.Getenv("DB_URL"),
		},
		Collections: Collections{
			Competitions: os.Getenv("COMPETITIONS_COLLECTION"),
		},
	}, nil
}
