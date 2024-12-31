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
	app.Use(middleware.AuthMiddleware)

	app.Get("/manage/create-user-role", mh.createUserRole)
	app.Get("/manage/delete-user-role", mh.deleteUserRole)

	app.Get("/manage/create-role-permission", mh.createRolePermission)
	app.Get("/manage/delete-role-permission", mh.deleteRolePermission)

	app.Get("/manage/create-resource-action", mh.createResourceAction)
	app.Get("/manage/delete-resource-action", mh.deleteResourceAction)

	app.Get("/manage/create-permission-resource-action", mh.createPermissionResourceAction)
	app.Get("/manage/delete-permission-resource-action", mh.deletePermissionResourceAction)
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

func (mh *ManagerHandler) createResourceAction(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error while creating resource-action: %v", err.Error())})
	}
	id, err := mh.service.createResourceAction(resource, action)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("error while creating resource-action: %v", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (mh *ManagerHandler) deleteResourceAction(c *fiber.Ctx) error {
	if id, ok := c.Queries()["id"]; ok {
		err := mh.service.deleteResourceAction(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting resource-action"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("deleted %v", id)})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "'id' not in params"})
}

func (mh *ManagerHandler) createPermissionResourceAction(c *fiber.Ctx) error {
	qmap := c.Queries()
	var permission, resourceAction string
	var ok bool
	var err error
	if permission, ok = qmap["permission"]; ok {
		err = fmt.Errorf("could not find permission in query args")
	}
	if resourceAction, ok = qmap["resourceaction"]; ok {
		err = fmt.Errorf("could not find resourceaction in query args")
	}
	if err != nil {
		logger.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error while creating permission-resource-action: %v", err.Error())})
	}
	id, err := mh.service.createPermissionResourceAction(permission, resourceAction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("error while creating permission-resource-action: %v", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (mh *ManagerHandler) deletePermissionResourceAction(c *fiber.Ctx) error {
	if id, ok := c.Queries()["id"]; ok {
		err := mh.service.deletePermissionResourceAction(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error while deleting permission-resource-action"})
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

	createResourceAction(c *fiber.Ctx) error
	deleteResourceAction(c *fiber.Ctx) error

	createPermissionResourceAction(c *fiber.Ctx) error
	deletePermissionResourceAction(c *fiber.Ctx) error
}

var _ ManagerHandlerIface = &ManagerHandler{}
