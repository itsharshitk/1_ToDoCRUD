package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	if pass == "" {
		log.Fatal("No Password Found")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {
		log.Fatal("Password Encryption Failed: ", err.Error())
	}
	return string(hash)
}

func GenerateToken(user model.User) (string, error) {
	claims := model.CustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 23).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))

	return tokenString, err
}
