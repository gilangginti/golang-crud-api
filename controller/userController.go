package controller

import (
	"github.com/gin-gonic/gin"
	"golang-crud-api/initializers"
	"golang-crud-api/models"
	"golang-crud-api/utils"
	"net/http"
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
			"message": "Failed to create user",
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
