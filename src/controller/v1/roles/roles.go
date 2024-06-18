package roles

import (
	"GoMVCStarterKit/src/entity"
	"GoMVCStarterKit/src/utils/logger"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"xorm.io/xorm"
)

type rolesController struct {
	db *xorm.Session
}

func NewRolesController(db *xorm.Session) *rolesController {
	return &rolesController{db: db}
}

func (r *rolesController) FindAll(c *fiber.Ctx) error {
	var (
		roles []entity.Roles
	)
	if err := r.db.Find(&roles); err != nil {
		logger.Error(fmt.Sprintf("failed find roles got err :%v", err))
		return c.Status(http.StatusInternalServerError).JSON("false")
	}
	logger.Info("success find roles")
	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": roles,
	})
}
