package main

import (
	"GoMVCStarterKit/config"
	"GoMVCStarterKit/src/routes"
	"GoMVCStarterKit/src/utils/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.NewEnv()
	app := fiber.New(fiber.Config{
		AppName:       cfg.Application.Name,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	// Logger
	logger.RegisterLogger(cfg)
	// Routes
	routes.SetupRoutes(app)

	// Server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	go func() {
		<-ctx.Done()
		log.Info("Gracefully shutting down ...")

		if err := app.Shutdown(); err != nil {
			log.Errorf("app shutdown error :%v", err)
		} else {
			log.Info("app shutdown gracefully")
		}

	}()

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Application.Port)); err != nil {
		log.Fatalf("error starting app :%v", err)
	}

}
