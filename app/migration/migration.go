package migration

import (
	"praktikum_23/models"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	// auto migrate untuk table book
}
