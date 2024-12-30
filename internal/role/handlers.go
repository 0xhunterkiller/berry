package role

import (
	"time"

	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	service RoleServiceIface
}

func NewRoleHandler(service RoleServiceIface) RoleHandlerIface {
	return &RoleHandler{service: service}
}

func (rh *RoleHandler) RegisterRoutes(app *fiber.App) {
	app.Use("/role", middleware.AuthMiddleware)
	app.Post("/role/create", rh.CreateRole)
	app.Get("/role/delete", rh.DeleteRole)
}

type createReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (rh *RoleHandler) CreateRole(c *fiber.Ctx) error {
	if !c.Locals("chocolatedip").(bool) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
	}

	userID := c.Locals("userid").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
	}

	var req *createReq
	err := c.BodyParser(&req)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error while parsing your request"})
	}

	err = rh.service.createRole(req.Name, req.Description)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while creating role"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "created role"})
}

func (rh *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	if !c.Locals("chocolatedip").(bool) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
	}

	userID := c.Locals("userid").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
	}

	delID := c.Params("id")
	if delID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' param not found"})
	}
	err := rh.service.deleteRole(delID)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting role"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "deleted role"})
}

type RoleHandlerIface interface {
	RegisterRoutes(app *fiber.App)
	CreateRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error
}

var _ RoleHandlerIface = &RoleHandler{}
