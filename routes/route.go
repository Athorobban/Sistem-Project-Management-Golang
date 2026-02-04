package routes

import (
	"log"

	"github.com/Athorobban/Sistem-Project-Management-Golang/config"
	"github.com/Athorobban/Sistem-Project-Management-Golang/controllers"
	"github.com/Athorobban/Sistem-Project-Management-Golang/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App,uc *controllers.UserController){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	//JWT Protected Routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func (c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c,"Error unauthorized",err.Error())
		},
	}))

	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination)
	userGroup.Get("/:id", uc.GetUser)
	userGroup.Put("/:id",uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)
}