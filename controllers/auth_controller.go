package controllers

import (
	"fmt"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Register
// @Description  register a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.CredentialsPayload		true	"User Credentials"
// @Success      201  {object} models.UserJson
// @Router       /api/auth/register [post]
func Register(c *gin.Context) {
	params := payloads.CredentialsPayload{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hashedPassword, err2 = utilities.HashPassword(params.Password)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	user := models.User{
		UserName: params.UserName,
		PassWord: hashedPassword,
	}
	user.CreatedAt = now
	user.UpdatedAt = now

	initializers.DB.Save(&user)

	if user.Id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Save User"})
		return
	}

	user.PassWord = ""

	// response
	c.JSON(http.StatusCreated, user.Json())
}

// @Summary      Login
// @Description  Login to get an auth token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.CredentialsPayload		true	"User Credentials"
// @Success      200  {object} payloads.LoginResponse
// @Router       /api/auth/login [post]
func Login(c *gin.Context) {
	params := payloads.CredentialsPayload{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error with payload")
		return
	}

	user := models.User{}
	initializers.DB.Where("UserName = ?", params.UserName).First(&user)

	if !utilities.PasswordsMatch(user.PassWord, params.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		fmt.Println("error matching")
		return
	}

	tokenString, err := utilities.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		fmt.Println("error token")
		return
	}

	response := payloads.LoginResponse{
		UserName: user.UserName,
		Token:    tokenString,
	}

	// response
	c.JSON(http.StatusOK, response)
}
