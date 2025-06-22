package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"github.com/itsharshitk/1_ToDoCRUD/utils"
	"golang.org/x/crypto/bcrypt"
)

// Signup godoc
// @Summary      Register a new user
// @Description  Creates a new user with name, email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User Data"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  model.APIResponse
// @Router       /signup [post]
func Signup(c *gin.Context) {
	var req model.User
	var foundUser model.User

	db := config.Db

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
			Details: err.Error(),
		})
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errs := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errs[err.Field()] = utils.GetValidationMessage(err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": errs})
		return
	}

	if err := db.Where("email = ?", req.Email).First(&foundUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "User Already Exists",
		})
		return
	}

	req.Password = utils.HashPassword(req.Password)

	if err := db.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Can't create user",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{
		ID:      req.ID,
		Name:    req.Name,
		Email:   req.Email,
		Message: "User Created Successfully",
	})

}

// Login godoc
// @Summary      Login existing user
// @Description  Authenticates a user and returns JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      model.LoginRequest  true  "Login credentials"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  model.APIResponse
// @Router       /login [post]
func Login(c *gin.Context) {
	var req model.LoginRequest
	var foundUser model.User

	db := config.Db

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &model.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Something Went Wrong",
			Details: err.Error(),
		})
		return
	}

	if err := db.Where("email = ?", req.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "User Not Registered",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, &model.APIResponse{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Password",
		})
		return
	}

	token, err := utils.GenerateToken(foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Token Generation Failed",
		})
		return
	}

	c.JSON(http.StatusOK, &model.UserResponse{
		ID:      foundUser.ID,
		Name:    foundUser.Name,
		Email:   foundUser.Email,
		Message: "Login Successful",
		Token:   token,
	},
	)
}
