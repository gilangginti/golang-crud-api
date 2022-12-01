package main

import (
	"github.com/gin-gonic/gin"
	"golang-crud-api/controller"
	"golang-crud-api/initializers"
	"golang-crud-api/migration"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	migration.UserMigration()
}

func main() {
	r := gin.Default()
	defer r.Run()

	//Router

	//User Rotuer
	r.POST("/user", controller.SignUp)

}
