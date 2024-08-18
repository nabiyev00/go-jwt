package initializers

import "github.com/nabiyev00/go-jwt/models"

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
