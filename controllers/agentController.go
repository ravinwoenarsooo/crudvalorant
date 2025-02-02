package controllers

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ravinwoenarsooo/crudvalorant/database"
	"github.com/ravinwoenarsooo/crudvalorant/models/entity"
	"github.com/ravinwoenarsooo/crudvalorant/models/req"
)

func AgentControllerShowAll(c *fiber.Ctx) error {
	var agent []entity.Agent
	//Cek apakah ada data di database
	err := database.DB.Find(&agent).Error
	if err != nil {
		log.Println(err)
	}

	if len(agent) == 0 {
		return c.JSON(fiber.Map{
			"message": "No data on database",
		})
	}

	return c.JSON(agent)
}

func AgentControllerCreate(c *fiber.Ctx) error {
	agent := new(req.AgentReq)
	if err := c.BodyParser(agent); err != nil {
		return err
	}

	// Validasi input user pakai package validator/v10
	validate := validator.New()
	if err := validate.Struct(agent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   "Body not valid",
			"error": err.Error(),
		})
	}

	//Cek apakah roleID ada di database Roles
	// var role entity.Roles
	// if err := database.DB.First(&role, agent.Role_Id).Error; err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"msg": "RoleID not found",
	// 	})
	// }

	//Assign data input ke struct userReq
	newUser := entity.Agent{
		Agent_Name: agent.Name,
		Role_ID:    uint(agent.Role_Id),
	}

	//Bikin user baru
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to create new agent",
		})
	}

	var CreatedAgent entity.Agent
	if err := database.DB.Preload("Roles").First(&CreatedAgent, newUser.Agent_ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to fetch agent data",
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":  "Success create new agent",
		"data": CreatedAgent,
	})

}

func AgentControllerGetById(c *fiber.Ctx) error {
	var agent []entity.Agent
	id := c.Params("id")

	//Ngecek id input user ada atau tidak.
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id is required",
		})
		return nil
	}

	//Ngecek apakah id ada di database.
	if err := database.DB.Where("Agent_ID = ?", id).First(&agent).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data not found",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    agent,
	})
}

// func AgentControllerUpdate(c *fiber.Ctx) error {
// 	agent := new(req.AgentReq)
// 	if err := c.BodyParser(agent); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Failed to parse body",
// 			"error":   err.Error(),
// 		})
// 	}

// 	validate := validator.New()
// 	if err := validate.Struct(agent); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Body not valid",
// 			"error":   err.Error(),
// 		})
// 	}

// 	updatedUser := entity.Agent{
// 		Agent_Name: agent.Name,
// 		Role_ID: agent.Role,
// 	}

// 	if err := database.DB.Model(&entity.Agent{}).Where("id = ?", c.Params("id")).Updates(updatedUser).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to update agent",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Success update agent",
// 		"data":    updatedUser,
// 	})
// }

func AgentControllerDeleteById(c *fiber.Ctx) error {
	var agent []entity.Agent
	id := c.Params("id")

	//Ngecek id input user ada atau tidak.
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id is required",
		})
		return nil
	}

	if err := database.DB.Where("Agent_ID = ?", id).First(&agent).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Agent data not found",
		})
		return nil
	}

	//Ngecek apakah id ada di database.
	if err := database.DB.Where("Agent_ID = ?", id).Delete(&agent).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Delete Agent Failed",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Agent has been deleted",
	})
}
