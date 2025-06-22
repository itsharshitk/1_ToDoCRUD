package utils

import (
	"log"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"golang.org/x/crypto/bcrypt"
)

var Validate *validator.Validate

func InitValidations() {
	Validate = validator.New()

	Validate.RegisterValidation("customPassVal", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 6 || len(password) > 20 {
			return false
		}

		upper := regexp.MustCompile(`[A-Z]`)
		lower := regexp.MustCompile(`[a-z]`)
		number := regexp.MustCompile(`[0-9]`)
		symbol := regexp.MustCompile(`[\W_]`) // Non-word (includes !@#$, etc.)

		return upper.MatchString(password) && lower.MatchString(password) && symbol.MatchString(password) && number.MatchString(password)
	})
}

func GetValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return err.Field() + " is too short"
	case "max":
		return err.Field() + " is too long"
	case "customPassVal":
		return "Password must be 6 to 20 characters long, include at least one uppercase letter, one lowercase letter, one number, and one special character"
	default:
		return "Invalid value for " + err.Field()
	}
}

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
