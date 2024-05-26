package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrheaven778/go-astro-crud/config"
	"github.com/mrheaven778/go-astro-crud/db"
	"github.com/mrheaven778/go-astro-crud/models"
	"github.com/mrheaven778/go-astro-crud/validators"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var SecretKey = config.GetSecretKey()

func LoginHandler(c *fiber.Ctx) error {
	var user = models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var userInDB models.User
	userExists := db.DB.Where("email = ?", user.Email).First(&userInDB)
	if userExists.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	err := CheckPassword(userInDB.Password, user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid password"})
	}

	t, err := GenerateJWT(strconv.Itoa(int(userInDB.ID)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}

func GenerateJWT(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func RegisterHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	errMessage := validators.ValidateUser(&user)
	if errMessage != "" {
		return c.Status(400).JSON(fiber.Map{"message": errMessage})
	}

	var existingUser models.User
	db.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Email already in use"})
	}

	// Check if user already exists by username
	var existingUserByUsername models.User
	db.DB.Where("username = ?", user.Username).First(&existingUserByUsername)
	if existingUserByUsername.ID != 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Username already in use"})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.SendString("Error hashing password")
	}
	user.Password = hash
	createUser := db.DB.Create(&user)
	if createUser.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": createUser.Error.Error()})
	}
	t, err := GenerateJWT(strconv.Itoa(int(user.ID)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(201).JSON(fiber.Map{"user": &user, "token": t})
}
