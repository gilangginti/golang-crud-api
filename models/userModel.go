package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v1()"`
	Username  string         `gorm:"type:varchar(100);unique"`
	Password  string         `gorm:"type:varchar(100)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
