package services

import (
	"context"
	"crypto/sha256"
	"github.com/Nerzal/gocloak/v11"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"uacs/sso-server/internal/config"
	"uacs/sso-server/internal/models"
)

type Services struct {
	Log            *zap.SugaredLogger
	Cfg            *config.Config
	KeyloackClient gocloak.GoCloak
}

func (s *Services) Registration(newUser models.User) error {
	token, err := s.KeyloackClient.LoginAdmin(context.Background(), s.Cfg.KeycloakAdminUsername, s.Cfg.KeycloakAdminPassword, "master")
	if err != nil {
		return err
	}

	user := gocloak.User{
		FirstName: gocloak.StringP(newUser.FirstName),
		LastName:  gocloak.StringP(newUser.LastName),
		Email:     gocloak.StringP(newUser.Email),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP(newUser.Username),
	}

	usrId, err := s.KeyloackClient.CreateUser(context.Background(), token.AccessToken, s.Cfg.KeycloakRealmName, user)
	if err != nil {
		return err
	}

	bytePwd := sha256.Sum256([]byte(newUser.Password))

	err = s.KeyloackClient.SetPassword(context.Background(), token.AccessToken, usrId, s.Cfg.KeycloakRealmName, string(bytePwd[:]), false)
	if err != nil {
		return err
	}

	return nil
}

func (s *Services) Login(loginReq models.LoginRequest) (models.LoginResponse, error) {
	bytePwd := sha256.Sum256([]byte(loginReq.Password))

	jwt, err := s.KeyloackClient.Login(context.Background(),
		s.Cfg.KeycloakClientId,
		s.Cfg.KeycloakClientSecret,
		s.Cfg.KeycloakRealmName,
		loginReq.Username,
		string(bytePwd[:]),
	)
	if err != nil {
		return models.LoginResponse{}, err
	}

	resp := models.LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	return resp, nil
}

func (s *Services) ValidateAccessToken(token string) (bool, error) {
	rptResult, err := s.KeyloackClient.RetrospectToken(context.Background(), token, s.Cfg.KeycloakClientId, s.Cfg.KeycloakClientSecret, s.Cfg.KeycloakRealmName)
	if err != nil {
		return false, err
	}

	if !*rptResult.Active {
		return false, errors.New("token is not active")
	}

	return true, nil
}
