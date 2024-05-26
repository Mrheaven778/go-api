// auth.go
package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserFromContext(c *fiber.Ctx) (string, error) {
	userToken := c.Locals("user")
	if userToken == nil {
		return "", fmt.Errorf("user not found in context")
	}
	token, ok := userToken.(*jwt.Token)
	if !ok {
		return "", fmt.Errorf("unexpected token type")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("unexpected claims type")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("user_id not found in claims")
	}
	return userID, nil
}
