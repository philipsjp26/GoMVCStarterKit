package routes

import (
	"GoMVCStarterKit/src/controller/v1/roles"
	"GoMVCStarterKit/src/utils/logger"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"xorm.io/xorm"
)

func SetupRoutes(router fiber.Router, db *xorm.Session) {
	apiV1 := router.Group("/api").Group("/v1")

	// controller
	rolesController := roles.NewRolesController(db)

	apiV1.Get("/roles", rolesController.FindAll)

	router.Get("/health", func(c *fiber.Ctx) error {
		logger.Info("test")
		return c.Status(http.StatusOK).JSON(true)
	})
}
