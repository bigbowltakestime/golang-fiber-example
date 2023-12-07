package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetHome는 홈 페이지를 반환하는 핸들러 함수입니다.
func GetHome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the home page!")
}

// PostMessage는 메시지를 받아서 처리하는 핸들러 함수입니다.
func PostMessage(c *fiber.Ctx) error {
	var request struct {
		Message string `json:"message"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// 받은 메시지를 처리하는 로직을 추가할 수 있습니다.

	return c.SendString("Message received: " + request.Message)
}
