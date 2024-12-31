package manager

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type ManagerHandler struct {
	service ManagerServiceIface
}

func NewManagerHandler(service ManagerServiceIface) ManagerHandlerIface {
	return &ManagerHandler{service: service}
}

func (mh *ManagerHandler) RegisterRoutes(app *fiber.App) {
	app.Use("/manage", middleware.AuthMiddleware)

	app.Get("/manage/create-user-role", mh.createUserRole)
	app.Get("/manage/delete-user-role", mh.deleteUserRole)

	app.Get("/manage/create-role-permission", mh.createRolePermission)
	app.Get("/manage/delete-role-permission", mh.deleteRolePermission)

	app.Get("/manage/create-interaction", mh.createInteraction)
	app.Get("/manage/delete-interaction", mh.deleteInteraction)

	app.Get("/manage/create-permission-interaction", mh.createPermissionInteraction)
	app.Get("/manage/delete-permission-interaction", mh.deletePermissionInteraction)
}

func (mh *ManagerHandler) createUserRole(c *fiber.Ctx) error {
	qmap := c.Queries()
	var user, role string
	var ok bool
	var err error
	if user, ok = qmap["user"]; ok {
		err = fmt.Errorf("could not find user in query args")
	}
	if role, ok = qmap["role"]; ok {
		err = fmt.Errorf("could not find role in query args")
	}
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error while creating user-role: %v", err.Error())})
	}
	id, err := mh.service.createUserRole(user, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("error while creating user-role: %v", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (mh *ManagerHandler) deleteUserRole(c *fiber.Ctx) error {
	if id, ok := c.Queries()["id"]; ok {
		err := mh.service.deleteUserRole(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting user-role"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("deleted %v", id)})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' not in params"})
}

func (mh *ManagerHandler) createRolePermission(c *fiber.Ctx) error {
	qmap := c.Queries()
	var role, permission string
	var ok bool
	var err error
	if role, ok = qmap["role"]; ok {
		err = fmt.Errorf("could not find role in query args")
	}
	if permission, ok = qmap["permission"]; ok {
		err = fmt.Errorf("could not find permission in query args")
	}
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error while creating role-permission: %v", err.Error())})
	}
	id, err := mh.service.createRolePermission(role, permission)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("error while creating role-permission: %v", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (mh *ManagerHandler) deleteRolePermission(c *fiber.Ctx) error {
	if id, ok := c.Queries()["id"]; ok {
		err := mh.service.deleteRolePermission(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting role-permission"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("deleted %v", id)})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' not in params"})
}

func (mh *ManagerHandler) createInteraction(c *fiber.Ctx) error {
	qmap := c.Queries()
	var resource, action string
	var ok bool
	var err error
	if resource, ok = qmap["resource"]; ok {
		err = fmt.Errorf("could not find resource in query args")
	}
	if action, ok = qmap["action"]; ok {
		err = fmt.Errorf("could not find action in query args")
	}
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error while creating interaction: %v", err.Error())})
	}
	id, err := mh.service.createInteraction(resource, action)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("error while creating interaction: %v", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (mh *ManagerHandler) deleteInteraction(c *fiber.Ctx) error {
	if id, ok := c.Queries()["id"]; ok {
		err := mh.service.deleteInteraction(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting interaction"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("deleted %v", id)})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' not in params"})
}

func (mh *ManagerHandler) createPermissionInteraction(c *fiber.Ctx) error {
	qmap := c.Queries()
	var permission, interaction string
	var ok bool
	var err error
	if permission, ok = qmap["permission"]; ok {
		err = fmt.Errorf("could not find permission in query args")
	}
	if interaction, ok = qmap["interaction"]; ok {
		err = fmt.Errorf("could not find interaction in query args")
	}
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error while creating permission-interaction: %v", err.Error())})
	}
	id, err := mh.service.createPermissionInteraction(permission, interaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("error while creating permission-interaction: %v", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (mh *ManagerHandler) deletePermissionInteraction(c *fiber.Ctx) error {
	if id, ok := c.Queries()["id"]; ok {
		err := mh.service.deletePermissionInteraction(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting permission-interaction"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("deleted %v", id)})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' not in params"})
}

type ManagerHandlerIface interface {
	RegisterRoutes(app *fiber.App)
	createUserRole(c *fiber.Ctx) error
	deleteUserRole(c *fiber.Ctx) error

	createRolePermission(c *fiber.Ctx) error
	deleteRolePermission(c *fiber.Ctx) error

	createInteraction(c *fiber.Ctx) error
	deleteInteraction(c *fiber.Ctx) error

	createPermissionInteraction(c *fiber.Ctx) error
	deletePermissionInteraction(c *fiber.Ctx) error
}

var _ ManagerHandlerIface = &ManagerHandler{}
