package resource

import (
	"github.com/0xhunterkiller/berry/internal/helpers"
	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ResourceHandler struct {
	service ResourceServiceIface
}

func NewResourceHandler(service ResourceServiceIface) ResourceHandlerIface {
	return &ResourceHandler{service: service}
}
func (reh *ResourceHandler) RegisterRoutes(app *fiber.App) {
	app.Use("/resource", middleware.AuthMiddleware)
	app.Post("/resource/create", reh.CreateResource)
	app.Get("/resource/delete", reh.DeleteResource)
}

type createReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (reh *ResourceHandler) CreateResource(c *fiber.Ctx) error {
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

	id, err := reh.service.createResource(req.Name, req.Description)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while creating resource"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (reh *ResourceHandler) DeleteResource(c *fiber.Ctx) error {
	if _, ok := helpers.CheckAuthentication(c); !ok {
		return helpers.ForbiddenMsg(c)
	}

	delID := c.Queries()["id"]
	if delID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' param not found"})
	}
	err := reh.service.deleteResource(delID)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting resource"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "deleted resource"})
}

type ResourceHandlerIface interface {
	RegisterRoutes(app *fiber.App)
	CreateResource(c *fiber.Ctx) error
	DeleteResource(c *fiber.Ctx) error
}

var _ ResourceHandlerIface = &ResourceHandler{}
