package main

import (
	"log"
	"time"

	"golang-fiber-example/app/handlers"
	"golang-fiber-example/app/middleware"
	"golang-fiber-example/app/routes"
	"golang-fiber-example/app/services"
	"golang-fiber-example/config"

	// "golang-fiber-example/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 설정 로드
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Fiber 애플리케이션 생성
	app := fiber.New()

	// 미들웨어 등록
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	// 라우트 설정
	routes.SetupHomeRoutes(app)
	routes.SetupUserRoutes(app)

	// 서비스 생성
	userService := services.NewUserService()

	// 핸들러 생성
	homeHandler := handlers.NewHomeHandler(cfg, utils.GenerateRandomString)
	userHandler := handlers.NewUserHandler(userService)

	// 핸들러 등록
	routes.RegisterHomeHandler(app, homeHandler)
	routes.RegisterUserHandler(app, userHandler)

	// 서버 시작
	go func() {
		// 서비스가 준비될 때까지 대기
		time.Sleep(2 * time.Second)

		// 서버 시작
		err := app.Listen(":" + cfg.Server.Port)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Server is running on port %s\n", cfg.Server.Port)

	// 프로그램 종료까지 대기
	select {}
}
