package usecase

import (
	"errors"
	"holamundo/users/domain/repository"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type LoginUserUseCase struct {
	UserRepo repository.IUserRepository
}

func NewLoginUseCase(repo repository.IUserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{UserRepo: repo}
}

func (uc *LoginUserUseCase) Execute(email, password string) (string, error) {
	user, err := uc.UserRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("clave secreta de JWT no configurada")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
