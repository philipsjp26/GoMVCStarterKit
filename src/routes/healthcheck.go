package routes

import (
	"GoMVCStarterKit/src/utils/logger"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		logger.Info("test")
		return c.Status(http.StatusOK).JSON(true)
	})
}
