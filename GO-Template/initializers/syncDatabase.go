package initializers

import "go-template/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
