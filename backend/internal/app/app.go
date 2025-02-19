package app

import (
	"context"
	"log"
	"newsportal-backend/config"
	"newsportal-backend/internal/adapter/handler"
	"newsportal-backend/internal/adapter/repository"
	"newsportal-backend/internal/core/service"
	"newsportal-backend/lib/auth"
	"newsportal-backend/lib/middleware"
	"newsportal-backend/lib/pagination"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	cfgR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cfgR2)

	jwt := auth.NewJwt(cfg)
	_ = middleware.NewMiddleware(cfg)
	_ = pagination.NewPagination()

	// Repository
	authRepo := repository.NewAuthRepository(db.DB)

	// Service
	authService := service.NewAuthService(authRepo, cfg, jwt)

	// Handler
	authHandler := handler.NewAuthHandler(authService)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip} ${status} - ${latency} ${method} ${path}\n",
	}))

	api := app.Group("/api")
	api.Post("/login", authHandler.Login)

	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}

		err := app.Listen(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatalf("error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	log.Println("server shutdown of 5 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)
}
