package initializers

import "example.com/m/v2/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
