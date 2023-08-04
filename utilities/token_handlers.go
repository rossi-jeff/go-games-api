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
		return userId
	}
	tokenString := strings.Split(authorization[0], " ")[1]
	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			fmt.Println(err)
			return userId
		}
		if !token.Valid {
			fmt.Println("invalid")
			return userId
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			sub := claims["sub"].(float64)
			userId = int(sub)
		}
	}
	return userId
}
