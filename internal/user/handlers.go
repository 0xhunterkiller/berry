package user

import (
	"fmt"
	"time"

	"github.com/0xhunterkiller/berry/internal/helpers"
	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/pkg/jwtutil"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service UserServiceIface
}

func NewUserHandler(service UserServiceIface) UserHandlerIface {
	return &UserHandler{service: service}
}

func (uh *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/auth/user/register", uh.registerUser)
	app.Post("/auth/user/login", uh.loginUser)

	app.Use("/user", middleware.AuthMiddleware)
	app.Get("/user/checkauth", uh.amIAuthenticated)
	app.Patch("/user/email", uh.updateEmail)
	app.Patch("/user/password", uh.updatePassword)
	app.Patch("/user/deactivate", uh.deactivateUser)
	app.Patch("/user/activate", uh.activateUser)
	app.Delete("/user", uh.deleteUser)
}

// Am I Authenticated
func (uh *UserHandler) amIAuthenticated(c *fiber.Ctx) error {
	if !c.Locals("chocolatedip").(bool) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
	}

	user, err := uh.service.getByID(c.Locals("userid").(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"timestamp": time.Now(), "message": fmt.Sprintf("you are authenticated as %v", user.Username)})
}

// Register User
type registerRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (uh *UserHandler) registerUser(c *fiber.Ctx) error {
	var req registerRequest

	err := c.BodyParser(&req)
	if err != nil {
		logger.Logger.Error("the request body was invalid: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "the request body is invalid"})
	}

	v := validator.New()
	if err := v.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}

	id, err := uh.service.createUser(req.Username, req.Email, req.Password, true)
	if err != nil {
		logger.Logger.Error("could not create user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "the user was not created"})
	}
	logger.Logger.Info("the user was created successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

// Login User
type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func checkPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (uh *UserHandler) loginUser(c *fiber.Ctx) error {
	var claim loginRequest

	if err := c.BodyParser(&claim); err != nil {
		logger.Logger.Error("could not login user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	var user *models.UserModel
	user, err := uh.service.getByUsername(claim.Username)
	if err != nil {
		logger.Logger.Error("could not login user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect username or password"})
	}

	v := validator.New()
	if err := v.Struct(&claim); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}

	jwtClaim := jwt.MapClaims{
		"iss":    "berry-authn",
		"sub":    user.Username,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour).Unix(),
		"userid": user.ID,
	}

	if checkPassword(user.Password, claim.Password) {
		jwt, err := jwtutil.GenerateJWT(&jwtClaim)
		if err != nil {
			logger.Logger.Error("could not login user due to an error: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to authenticate user"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"jwt": jwt, "user": fiber.Map{"username": user.Username, "ts": time.Now().Unix()}})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect username or password"})
}

// update password
type updateEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (uh *UserHandler) updateEmail(c *fiber.Ctx) error {
	id, ok := helpers.CheckAuthentication(c)
	if !ok {
		return helpers.ForbiddenMsg(c)
	}

	var uer updateEmailRequest

	if err := c.BodyParser(&uer); err != nil {
		logger.Logger.Error("could not update user email due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	v := validator.New()
	if err := v.Struct(&uer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}

	err := uh.service.updateEmail(id, uer.Email)
	if err != nil {
		logger.Logger.Error("an error occured while updating email: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "an error occured"})
	}

	return nil
}

// update password
type updatePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

func (uh *UserHandler) updatePassword(c *fiber.Ctx) error {
	id, ok := helpers.CheckAuthentication(c)
	if !ok {
		return helpers.ForbiddenMsg(c)
	}

	var upr updatePasswordRequest

	if err := c.BodyParser(&upr); err != nil {
		logger.Logger.Error("could not update user password due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	v := validator.New()
	if err := v.Struct(&upr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}

	err := uh.service.updatePassword(id, upr.Password)
	if err != nil {
		logger.Logger.Error("an error occured while updating password: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "an error occured"})
	}

	return nil
}

func (uh *UserHandler) deactivateUser(c *fiber.Ctx) error {
	id, ok := helpers.CheckAuthentication(c)
	if !ok {
		return helpers.ForbiddenMsg(c)
	}

	err := uh.service.deactivateUser(id)
	if err != nil {
		logger.Logger.Error("an error occured while deactivating user: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "an error occured"})
	}
	return nil
}

func (uh *UserHandler) activateUser(c *fiber.Ctx) error {
	id, ok := helpers.CheckAuthentication(c)
	if !ok {
		return helpers.ForbiddenMsg(c)
	}

	err := uh.service.activateUser(id)
	if err != nil {
		logger.Logger.Error("an error occured while activating user: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "an error occured"})
	}
	return nil
}

func (uh *UserHandler) deleteUser(c *fiber.Ctx) error {
	id, ok := helpers.CheckAuthentication(c)
	if !ok {
		return helpers.ForbiddenMsg(c)
	}

	err := uh.service.deleteUser(id)
	if err != nil {
		logger.Logger.Error("an error occured while deleting user: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "an error occured"})
	}
	return nil
}

type UserHandlerIface interface {
	RegisterRoutes(app *fiber.App)
	amIAuthenticated(c *fiber.Ctx) error
	registerUser(c *fiber.Ctx) error
	loginUser(c *fiber.Ctx) error
	updateEmail(c *fiber.Ctx) error
	updatePassword(c *fiber.Ctx) error
	deactivateUser(c *fiber.Ctx) error
	activateUser(c *fiber.Ctx) error
	deleteUser(c *fiber.Ctx) error
}

var _ UserHandlerIface = &UserHandler{}
