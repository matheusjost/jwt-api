package initializers

import "github.com/matheusjost/jwt-api/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
