package utilities

import (
	"fmt"
	"go-games-api/models"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Minute * 60 * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func UserIdFromAuthHeader(c *gin.Context) int {
	userId := 0
	authorization := c.Request.Header["Authorization"]
	if len(authorization) == 0 {
		return 0
	}
	tokenString := strings.Split(authorization[0], " ")[1]
	if tokenString != "" {
		fmt.Println(tokenString)
	}
	return userId
}
