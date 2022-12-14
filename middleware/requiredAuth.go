package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang-crud-api/initializers"
	"golang-crud-api/models"
	"golang-crud-api/utils"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	//Get Cookies
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error get token",
		})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check Expire
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token expired!",
			})
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])
		ui := user.ID.String()
		if utils.IsValidUUID(ui) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid id from user",
			})
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
