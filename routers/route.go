package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/ravinwoenarsooo/crudvalorant/controllers"
)

func RouterApp(c *fiber.App) {

	c.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	c.Get("/api/showall", controllers.AgentControllerShowAll)
	c.Get("/api/getUserById", controllers.AgentControllerGetById)
	c.Post("/api/create", controllers.AgentControllerCreate)
	c.Put("/api/updateUser/:id", controllers.AgentControllerUpdate)
	c.Delete("/api/delete/:id", controllers.AgentControllerDeleteById)
}
