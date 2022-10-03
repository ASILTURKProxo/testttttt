package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	// api.Post("/register", controllers.Register)
	Auth(api) //user login and register
	UsersRoutes(api)
	UserAbilityRoutes(api)
}
