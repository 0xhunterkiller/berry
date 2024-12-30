package action

import (
	"time"

	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type ActionHandler struct {
	service ActionServiceIface
}

func NewActionHandler(service ActionServiceIface) ActionHandlerIface {
	return &ActionHandler{service: service}
}

func (ah *ActionHandler) RegisterRoutes(app *fiber.App) {
	app.Use("/action", middleware.AuthMiddleware)
	app.Post("/action/create", ah.CreateAction)
	app.Get("/action/delete", ah.DeleteAction)
}

type createReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (ah *ActionHandler) CreateAction(c *fiber.Ctx) error {
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

	err = ah.service.createAction(req.Name, req.Description)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while creating action"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "created action"})
}

func (ah *ActionHandler) DeleteAction(c *fiber.Ctx) error {
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
	err := ah.service.deleteAction(delID)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting action"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "deleted action"})
}

type ActionHandlerIface interface {
	RegisterRoutes(app *fiber.App)
	CreateAction(c *fiber.Ctx) error
	DeleteAction(c *fiber.Ctx) error
}

var _ ActionHandlerIface = &ActionHandler{}
