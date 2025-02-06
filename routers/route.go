package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/ravinwoenarsooo/crudvalorant/controllers"
)

func RouterApp(c *fiber.App) {
	c.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	
	c.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	c.Get("/api/showall", controllers.AgentControllerShowAll)
	c.Get("/api/getUserById", controllers.AgentControllerGetById)
	c.Get("/api/showAllAgentbyRoles", controllers.ShowAllAgentByRoles)
	c.Post("/api/create", controllers.AgentControllerCreate)
	c.Put("/api/updateUser", controllers.AgentControllerUpdate)
	c.Delete("/api/delete", controllers.AgentControllerDeleteById)
}
