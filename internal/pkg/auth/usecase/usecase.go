package usecase

import (
	"crypto/sha256"
	"encoding/hex"
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
	id, err := u.authRepo.CheckUser(name, hashString(password))
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
	id, err := u.authRepo.CreateUser(name, hashString(password))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func hashString(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))

	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}
