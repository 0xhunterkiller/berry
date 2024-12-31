package permission

import (
	"github.com/0xhunterkiller/berry/internal/helpers"
	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	service PermissionServiceIface
}

func NewPermissionHandler(service PermissionServiceIface) PermissionHandlerIface {
	return &PermissionHandler{service: service}
}

func (ph *PermissionHandler) RegisterRoutes(app *fiber.App) {
	app.Use("/permission", middleware.AuthMiddleware)
	app.Post("/permission/create", ph.CreatePermission)
	app.Get("/permission/delete", ph.DeletePermission)
}

type createReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (ph *PermissionHandler) CreatePermission(c *fiber.Ctx) error {
	if _, ok := helpers.CheckAuthentication(c); !ok {
		return helpers.ForbiddenMsg(c)
	}

	var req *createReq
	err := c.BodyParser(&req)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error while parsing your request"})
	}

	v := validator.New()
	err = v.Struct(req)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error while parsing your request"})
	}

	id, err := ph.service.createPermission(req.Name, req.Description)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while creating permission"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (ph *PermissionHandler) DeletePermission(c *fiber.Ctx) error {
	if _, ok := helpers.CheckAuthentication(c); !ok {
		return helpers.ForbiddenMsg(c)
	}
	delID := c.Queries()["id"]
	if delID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' param not found"})
	}
	err := ph.service.deletePermission(delID)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting permission"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "deleted permission"})
}

type PermissionHandlerIface interface {
	RegisterRoutes(app *fiber.App)
	CreatePermission(c *fiber.Ctx) error
	DeletePermission(c *fiber.Ctx) error
}

var _ PermissionHandlerIface = &PermissionHandler{}
