package config

import (
	"fmt"
	"os"
)

type Config struct {
	SsoServerDevPort      string
	SsoServerProdPort     string
	KeycloakURL           string
	KeycloakAdminUsername string
	KeycloakAdminPassword string
	KeycloakClientSecret  string
	KeycloakRealmName     string
	KeycloakClientId      string
}

func New() (*Config, error) {
	return &Config{
		SsoServerDevPort:      os.Getenv("SSO_SERVER_DEV_PORT"),
		SsoServerProdPort:     os.Getenv("SSO_SERVER_PROD_PORT"),
		KeycloakURL:           fmt.Sprintf("http://localhost:%s", os.Getenv("KEYCLOAK_PORT")),
		KeycloakAdminUsername: os.Getenv("KEYCLOAK_ADMIN_USERNAME"),
		KeycloakAdminPassword: os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
		KeycloakClientSecret:  os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		KeycloakRealmName:     os.Getenv("KEYCLOAK_REALM_NAME"),
		KeycloakClientId:      os.Getenv("KEYCLOAK_CLIENT_ID"),
	}, nil
}
