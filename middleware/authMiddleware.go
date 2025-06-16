package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Header"})
			c.Abort()
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
			c.Abort()
			return
		}

		tokenReceived := parts[1]

		claims := &model.CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenReceived, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRETKEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token2"})
			c.Abort()
			return
		}

		c.Set("id", claims.ID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Next()
	}
}
