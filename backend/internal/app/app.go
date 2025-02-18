package app

import (
	"context"
	"log"
	"newsportal-backend/config"
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
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	cfgR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cfgR2)

	_ = auth.NewJwt(cfg)
	_ = middleware.NewMiddleware(cfg)
	_ = pagination.NewPagination()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] %{ip} &{status} - %{latency} %{method} %{path}\n",
	}))

	_ = app.Group("/api")

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
