package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ravinwoenarsooo/crudvalorant/database"
	"github.com/ravinwoenarsooo/crudvalorant/database/migration"
	"github.com/ravinwoenarsooo/crudvalorant/routers"
)

func main() {
	database.ConnectDatabase()
	migration.RunMigrate()
	app := fiber.New()

	routers.RouterApp(app)

	app.Listen(":8000")
}
