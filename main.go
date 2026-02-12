package main

import (
	"log"

	"github.com/Athorobban/Sistem-Project-Management-Golang/config"
	"github.com/Athorobban/Sistem-Project-Management-Golang/controllers"
	"github.com/Athorobban/Sistem-Project-Management-Golang/database/seed"
	"github.com/Athorobban/Sistem-Project-Management-Golang/repositories"
	"github.com/Athorobban/Sistem-Project-Management-Golang/routes"
	"github.com/Athorobban/Sistem-Project-Management-Golang/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()
	// user setup
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewuserController(userService)

	// board setup
	boardRepo := repositories.NewBoardRepository()
	boardService := services.NewBoardService(boardRepo,userRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":" + port))
}