package utils

import "github.com/gofiber/fiber"

func HandleEmptyUserOrPass(username string, password string) error {
	if username == "" || password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "username or password cannot be empty")
	}
	return nil
}

// Implement
func HandleSimilarUsers(username string) error {
	return nil
}
