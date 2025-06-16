package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"github.com/itsharshitk/1_ToDoCRUD/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var req model.User
	var foundUser model.User

	db := config.GetDB()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", req.Email).First(&foundUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Already Exists"})
		return
	}

	req.Password = utils.HashPassword(req.Password)

	if err := db.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":      req.ID,
		"Name":    req.Name,
		"Email":   req.Email,
		"Message": "User Created Successfully",
	})

}

func Login(c *gin.Context) {
	var req model.User
	var foundUser model.User

	db := config.GetDB()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", req.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User Not Registered"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	token, err := utils.GenerateToken(foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token Generation Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"user": gin.H{
			"id":    foundUser.ID,
			"name":  foundUser.Name,
			"email": foundUser.Email,
		},
		"token": token,
	})
}
