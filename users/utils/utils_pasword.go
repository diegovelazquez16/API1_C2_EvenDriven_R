package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword toma una contraseña en texto plano y devuelve un hash seguro.
func HashPassword(password string) (string, error) {
	// Genera un hash bcrypt de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Devuelve el hash como una cadena de texto
	return string(hashedPassword), nil
}

// CheckPasswordHash verifica si la contraseña coincide con el hash almacenado
func CheckPasswordHash(password, hash string) bool {
	// Compara el hash de la contraseña con el almacenado
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
