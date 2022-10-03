package routes

import (
	c "e-vet/Controllers"

	"github.com/gofiber/fiber/v2"
)

func UserAbilityRoutes(api fiber.Router) {
	api = api.Group("/ability")
	api.Post("/store", c.UserAbilityCreateController)
}
