package database

import (
	"root/mk/internal/models"
)

func Migrate() {
	if err := DB.DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}
}
