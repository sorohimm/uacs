package config

import "os"

type Config struct {
	DevPort     string
	ProdPort    string
	DBAuthData  DBAuthenticationData
	Collections Collections
	SsoCfg      SsoCfg
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
	Competitions   string
	Participants   string
	Judges         string
	Qualifications string
}

type SsoCfg struct {
	TokenValidateEndpoint string
}

func New() (*Config, error) {
	return &Config{
		DevPort:  os.Getenv("MAIN_SERVER_DEV_PORT"),
		ProdPort: os.Getenv("MAIN_SERVER_PROD_PORT"),
		DBAuthData: DBAuthenticationData{
			Username: os.Getenv("DB_DEV_USERNAME"),
			Password: os.Getenv("DB_DEV_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			URL:      os.Getenv("DB_URL"),
		},
		Collections: Collections{
			Competitions:   os.Getenv("COMPETITIONS_COLLECTION"),
			Participants:   os.Getenv("PARTICIPANTS_COLLECTION"),
			Judges:         os.Getenv("JUDGES_COLLECTION"),
			Qualifications: os.Getenv("QUALIFICATIONS_COLLECTION"),
		},
		SsoCfg: SsoCfg{
			TokenValidateEndpoint: os.Getenv("KEYCLOAK_TOKEN_VALIDATE_ENDPOINT"),
		},
	}, nil
}
