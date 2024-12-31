package action

import (
	"github.com/0xhunterkiller/berry/internal/helpers"
	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/go-playground/validator/v10"
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

	id, err := ah.service.createAction(req.Name, req.Description)
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while creating action"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (ah *ActionHandler) DeleteAction(c *fiber.Ctx) error {
	if _, ok := helpers.CheckAuthentication(c); !ok {
		return helpers.ForbiddenMsg(c)
	}

	delID := c.Queries()["id"]
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
