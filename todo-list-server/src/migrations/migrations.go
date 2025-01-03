package migrations

import (
	"github.com/kenzohfs/m/src/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Workspace{})
	db.AutoMigrate(&models.Task{})
}
