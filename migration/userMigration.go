package migration

import (
	"golang-crud-api/initializers"
	"golang-crud-api/models"
)

func UserMigration() {
	initializers.DB.AutoMigrate(&models.User{})
}
