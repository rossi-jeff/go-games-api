package controllers

import (
	"fmt"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Register
// @Description  register a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.CredentialsPayload		true	"User Credentials"
// @Success      201  {object} models.User
// @Router       /api/auth/register [post]
func Register(c *gin.Context) {
	params := payloads.CredentialsPayload{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hashedPassword, err2 = HashPassword(params.Password)

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
	c.JSON(http.StatusCreated, user)
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

	if !PasswordsMatch(user.PassWord, params.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		fmt.Println("error matching")
		return
	}

	tokenString, err := GenerateToken(user)
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

func HashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func PasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Minute * 60 * 24 * 7).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
