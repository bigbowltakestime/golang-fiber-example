package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/project/app/handlers"
)

// SetupHomeRoutes는 홈 관련 라우트를 설정하는 함수입니다.
func SetupHomeRoutes(app *fiber.App) {
	home := app.Group("/home")
	home.Get("/", handlers.GetHome)
	home.Post("/message", handlers.PostMessage)
	// 다른 홈 관련 라우트들을 추가할 수 있습니다.
}
