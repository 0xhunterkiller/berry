package resource

import (
	"time"

	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
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

	err = reh.service.createResource(req.Name, req.Description)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while creating resource"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "created resource"})
}

func (reh *ResourceHandler) DeleteResource(c *fiber.Ctx) error {
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
