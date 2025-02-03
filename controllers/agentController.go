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
	var agent []entity.Agents
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

	if err := database.DB.Preload("Role").Find(&agent).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed fetch agent role data",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    agent,
	})
}

func AgentControllerCreate(c *fiber.Ctx) error {
	agent := new(req.AgentsReq)
	if err := c.BodyParser(agent); err != nil {
		return err
	}

	//Validasi input user pakai package validator/v10
	validate := validator.New()
	if err := validate.Struct(agent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   "Body not valid",
			"error": err.Error(),
		})
	}

	//Masukin data struct ini ke struct user
	newAgent := entity.Agents{
		Agent_Name: agent.Name,
		Role_Id:    agent.Role,
	}
	if err := database.DB.Create(&newAgent).Error; err != nil {
		//Kalau go-native status(500)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create new agent",
			"error":   err.Error(),
		})
	}

	var newCreatedAgent entity.Agents
	if err := database.DB.Preload("Role").First(&newCreatedAgent, newAgent.Agent_Id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get new agent",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success create new agent",
		"data":    newCreatedAgent,
	})
}

func AgentControllerGetById(c *fiber.Ctx) error {
	var agent []entity.Agents
	id := c.Query("id")

	//Ngecek id input user ada atau tidak.
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id is required",
		})
		return nil
	}

	//Ngecek apakah id ada di database.
	if err := database.DB.Where("Agent_Id = ?", id).First(&agent).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data not found",
			"error":   err.Error(),
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    agent,
	})
}

func AgentControllerUpdate(c *fiber.Ctx) error {
	agent := new(req.AgentsReq)
	if err := c.BodyParser(agent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(agent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Body not valid",
			"error":   err.Error(),
		})
	}

	updatedUser := entity.Agents{
		Agent_Name: agent.Name,
		Role_Id:    agent.Role,
	}

	if err := database.DB.Model(&entity.Agents{}).Where("Agent_Id = ?", c.Params("id")).Updates(updatedUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update agent",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success update agent",
		"data":    updatedUser,
	})
}

func AgentControllerDeleteById(c *fiber.Ctx) error {
	var agent []entity.Agents
	id := c.Params("id")

	//Ngecek id input user ada atau tidak.
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id is required",
		})
		return nil
	}

	if err := database.DB.Where("id = ?", id).First(&agent).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Agent data not found",
		})
		return nil
	}

	//Ngecek apakah id ada di database.
	if err := database.DB.Where("id = ?", id).Delete(&agent).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Delete Agent Failed",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Agent has been deleted",
	})
}

func ShowAllAgentByRoles(c *fiber.Ctx) error {
	var agent []entity.Agents
	id := c.Query("id")

	err := database.DB.Find(&agent).Error
	if err != nil {
		log.Println(err)
	}

	if len(agent) == 0 {
		return c.JSON(fiber.Map{
			"message": "No data on database",
		})
	}

	if err := database.DB.Preload("Role").Where("Role_Id = ?", id).Find(&agent).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed fetch agent data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    agent,
	})
}
