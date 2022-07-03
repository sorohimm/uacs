package interfaces

import "uacs/sso-server/internal/models"

type IService interface {
	Registration(user models.User) error
	Login(response models.LoginRequest) (models.LoginResponse, error)
	ValidateAccessToken(token string) (bool, error)
}
