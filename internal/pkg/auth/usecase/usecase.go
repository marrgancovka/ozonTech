package usecase

import (
	"errors"
	"fmt"
	"os"
	"ozonTech/internal/pkg/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (u *AuthUsecase) generateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": expirationTime.Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (u *AuthUsecase) Login(name, password string) (string, error) {
	// Проверка пользователя в репозитории
	id, err := u.authRepo.CheckUser(name, password)
	if err != nil {
		return "", err
	}

	token, err := u.generateJWT(id)
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

	token, err := u.generateJWT(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyJWT проверяет и извлекает информацию из JWT токена
func (u *AuthUsecase) VerifyJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
