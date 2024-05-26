package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mrheaven778/go-astro-crud/db"
	"github.com/mrheaven778/go-astro-crud/models"
	"github.com/mrheaven778/go-astro-crud/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	userID, err := utils.GetUserFromContext(c)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	fmt.Println(userID)
	allUser := db.DB.Find(&users)
	if allUser.Error != nil {
		return c.Status(400).JSON(allUser.Error.Error())
	}
	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	userFind := db.DB.First(&user, id)
	if userFind.Error != nil {
		return c.Status(404).JSON("User not found")
	}
	return c.JSON(&user)
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update user")
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	userFind := db.DB.First(&user, id)
	if userFind.Error != nil {
		return c.Status(404).JSON("User not found")
	}
	deleteUser := db.DB.Delete(&user, id)
	if deleteUser.Error != nil {
		return c.Status(500).JSON(deleteUser.Error.Error())
	}
	return c.Status(204).JSON("User deleted")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
