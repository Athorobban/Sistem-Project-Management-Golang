package seed

import (
	"log"

	"github.com/Athorobban/Sistem-Project-Management-Golang/config"
	"github.com/Athorobban/Sistem-Project-Management-Golang/models"
	"github.com/Athorobban/Sistem-Project-Management-Golang/utils"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name: "Super admin",
		Email: "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed too seed admin", err)
	} else {
		log.Println("Admin user seeded")
	}
}