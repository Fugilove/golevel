
package migrations

import (
	"gorm.io/gorm"
	"golevel/app/models"
)

func CreateUsersTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
