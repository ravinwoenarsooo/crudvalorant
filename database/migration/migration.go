package migration

import (
	"fmt"

	"github.com/ravinwoenarsooo/crudvalorant/database"
	"github.com/ravinwoenarsooo/crudvalorant/models/entity"
)

// Ini buat migrate table ke database.
func RunMigrate() {
	err := database.DB.AutoMigrate(&entity.Agent{}, &entity.Roles{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB Migration success")
	}
}
