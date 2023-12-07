package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/project/app/handlers"
)

// SetupUserRoutes는 사용자 관련 라우트를 설정하는 함수입니다.
func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/:id", handlers.GetUserByID)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)
	// 다른 사용자 관련 라우트들을 추가할 수 있습니다.
}
