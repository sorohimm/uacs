package controllers

import (
	"github.com/Nerzal/gocloak/v11"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"uacs/sso-server/internal/config"
)

type Controllers struct {
	Log            *zap.SugaredLogger
	KeyloackClient gocloak.GoCloak
	Cfg            *config.Config
}

func (c *Controllers) Registration(ctx *gin.Context) {
	token, err := c.KeyloackClient.LoginAdmin(ctx, c.Cfg.KeycloakAdminUsername, c.Cfg.KeycloakAdminPassword, "realmName")
	if err != nil {
		panic("Something wrong with the credentials or url")
	}

	user := gocloak.User{
		FirstName: gocloak.StringP("Bob"),
		LastName:  gocloak.StringP("Uncle"),
		Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("CoolGuy"),
	}

	_, err = c.KeyloackClient.CreateUser(ctx, token.AccessToken, "realm", user)
	if err != nil {
		panic("Oh no!, failed to create user :(")
	}
}
