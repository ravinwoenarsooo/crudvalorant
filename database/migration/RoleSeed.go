package migration

import (
	"fmt"

	"github.com/ravinwoenarsooo/crudvalorant/models/entity"
	"gorm.io/gorm"
)

func RoleSeed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.Roles{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		roles := []entity.Roles{
			{Role_Name: "Duelist"},
			{Role_Name: "Initiator"},
			{Role_Name: "Controller"},
			{Role_Name: "Sentinel"},
		}

		//Pakai transaction biar kalau ada 1 salah langsung cancel semua
		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&roles).Error; err != nil {
				return err
			}
			fmt.Println("Seeding roles success")
			return nil
		}); err != nil {
			return err
		}
	} else {
		fmt.Println("4 Roles already seeded")
	}
	return nil
}
