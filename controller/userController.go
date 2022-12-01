package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang-crud-api/initializers"
	"golang-crud-api/models"
	"golang-crud-api/utils"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {

	//Request Body
	var body struct {
		Username string
		Password string
	}

	//IF Body Null
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Failed to read body",
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	//Hashed Password
	hash, err := utils.HashPassword(body.Password)

	//If password error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Hash Password",
			"status":  http.StatusBadRequest,
			"data":    nil,
		})
		return
	}

	//Created User
	user := models.User{Username: body.Username, Password: hash}
	result := initializers.DB.Create(&user)

	//If failed create user
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
			"status":  http.StatusBadRequest,
			"data":    nil,
		})
		return
	}
	//Response
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "Success Create User",
		"status":  http.StatusOK,
	})

}

func Login(c *gin.Context) {
	//Request Body
	var body struct {
		Username string
		Password string
	}
	//IF Body Null
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Failed to read body",
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"status":  http.StatusBadRequest,
			"message": "Invalid username",
		})
		return
	}
	comparePass := utils.CheckPasswordHash(body.Password, user.Password)
	if comparePass != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"status":  http.StatusBadRequest,
			"message": "Invalid password",
		})
		return
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	mySigningKey := []byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"status":  http.StatusBadRequest,
			"message": "Failed create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"token":   tokenString,
		"status":  http.StatusOK,
		"message": "Success Login",
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"message": "Success You have token",
	})
}
