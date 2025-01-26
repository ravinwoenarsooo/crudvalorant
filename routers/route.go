package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ravinwoenarsooo/crudvalorant/controllers"
)

func RouterApp(c *fiber.App) {
	c.Get("/api/showall", controllers.AgentControllerShowAll)
	c.Get("/api/getUserById/:id", controllers.AgentControllerGetById)
	c.Post("/api/create", controllers.AgentControllerCreate)
	c.Put("/api/updateUser/:id", controllers.AgentControllerUpdate)
	c.Delete("/api/delete/:id", controllers.AgentControllerDeleteById)
}
