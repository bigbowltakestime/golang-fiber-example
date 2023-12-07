package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware는 인증을 수행하는 미들웨어입니다.
func AuthMiddleware(c *fiber.Ctx) error {
	// 여기에서 인증 로직을 추가할 수 있습니다.

	// 인증이 성공하면 다음 핸들러로 진행합니다.
	return c.Next()
}
