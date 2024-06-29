package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"ozonTech/internal/pkg/auth"
	"ozonTech/internal/utils"
)

type AuthUsecase struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(authRepo auth.AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}

// Claims - структура для хранения данных в JWT
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func (u *AuthUsecase) Login(name, password string) (string, error) {
	// Проверка пользователя в репозитории
	id, err := u.authRepo.CheckUser(name, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *AuthUsecase) SignUp(name, password string) (string, error) {
	id, err := u.authRepo.CreateUser(name, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(id)
	if err != nil {
		return "", err
	}
	return token, nil
}
