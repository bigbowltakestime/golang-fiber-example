package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetUserByID는 사용자 ID에 해당하는 정보를 반환하는 핸들러 함수입니다.
func GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	// userID를 사용하여 사용자 정보를 조회하는 로직을 추가할 수 있습니다.

	return c.SendString(fmt.Sprintf("User ID: %s", userID))
}

// CreateUser는 새로운 사용자를 생성하는 핸들러 함수입니다.
func CreateUser(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		// 추가 필드들을 정의할 수 있습니다.
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// 받은 정보를 사용하여 사용자를 생성하는 로직을 추가할 수 있습니다.

	return c.SendString("User created: " + request.Username)
}

// UpdateUser는 사용자 정보를 업데이트하는 핸들러 함수입니다.
func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	// userID를 사용하여 업데이트할 사용자 정보를 찾아서 처리하는 로직을 추가할 수 있습니다.

	return c.SendString(fmt.Sprintf("User updated: %s", userID))
}

// DeleteUser는 사용자를 삭제하는 핸들러 함수입니다.
func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	// userID를 사용하여 삭제할 사용자 정보를 찾아서 처리하는 로직을 추가할 수 있습니다.

	return c.SendString(fmt.Sprintf("User deleted: %s", userID))
}
