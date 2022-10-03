package routes

import (
	c "e-vet/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Auth(api fiber.Router) {
	api = api.Group("/auth")
	api.Post("/register", c.Register)
	api.Post("/login", c.Login)
}

func UsersRoutes(api fiber.Router) {
	// api = api.Group("/user")
}
