package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrheaven778/go-astro-crud/db"
	"github.com/mrheaven778/go-astro-crud/models"
	"github.com/mrheaven778/go-astro-crud/validators"
	"strconv"
)

func GetTasksHandler(c *fiber.Ctx) error {
	var tasks []models.Task
	allTasks := db.DB.Find(&tasks)
	if allTasks.Error != nil {
		return c.Status(500).JSON(allTasks.Error.Error())
	}
	return c.JSON(tasks)
}

func CreateTaskHandler(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	validationErros := validators.ValidateTask(&task)
	if validationErros != "" {
		return c.Status(400).JSON(validationErros)
	}
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(400).JSON("User id must be a number")
	}
	task.UserID = uint(userId)
	createTask := db.DB.Create(&task)
	if createTask.Error != nil {
		return c.Status(500).JSON(createTask.Error.Error())
	}
	return c.Status(201).JSON(task)
}

func GetTaskHandler(c *fiber.Ctx) error {
	c.Params("id")
	var task models.Task
	taskId := c.Params("id")
	taskQuery := db.DB.First(&task, taskId)
	if taskQuery.Error != nil {
		return c.Status(500).JSON(taskQuery.Error.Error())
	}
	return c.JSON(task)
}

func DeleteTaskHandler(c *fiber.Ctx) error {
	return c.SendString("Get task")
}
