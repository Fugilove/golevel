package migrations

import (
	"golevel/app/models"

	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
