package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware는 요청 로깅을 수행하는 미들웨어입니다.
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// 다음 핸들러로 진행합니다.
	if err := c.Next(); err != nil {
		return err
	}

	// 요청이 완료된 후 로깅을 수행합니다.
	log.Printf("Request %s %s - %v", c.Method(), c.OriginalURL(), time.Since(start))

	return nil
}
