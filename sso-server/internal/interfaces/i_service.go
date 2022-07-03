package interfaces

import "uacs/sso-server/internal/models"

type IService interface {
	Registration(user models.User) error
	Login(response models.LoginRequest) (models.Session, error)
	ValidateAccessToken(session models.Session) (bool, error)
	GetUserId(session models.Session) (string, error)
}
