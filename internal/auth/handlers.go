package auth

import (
	"time"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/internal/user"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userService user.UserServiceIface
}

func NewAuthHandler(userService *user.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (ah *AuthHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/auth/user/register", ah.registerUser)
	app.Post("/auth/user/login", ah.loginUser)
}

// Register User
type registerRequest struct {
	username string `validate:"required"`
	email    string `validate:"required,email"`
	password string `validate:"required"`
}

func (ah *AuthHandler) registerUser(c *fiber.Ctx) error {
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

	err = ah.userService.CreateUser(req.username, req.email, req.password, true)
	if err != nil {
		logger.Logger.Error("could not create user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "the user was not created"})
	}
	logger.Logger.Info("the user was created successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "the user was created successfully"})
}

// Login User

type loginRequest struct {
	username string `validate:"required"`
	password string `validate:"required"`
}

func checkPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (ah *AuthHandler) loginUser(c *fiber.Ctx) error {
	var claim loginRequest

	if err := c.BodyParser(&claim); err != nil {
		logger.Logger.Error("could not login user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	var user *models.UserModel
	user, err := ah.userService.GetByUsername(claim.username)
	if err != nil {
		logger.Logger.Error("could not login user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect username or password"})
	}

	v := validator.New()
	if err := v.Struct(&claim); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}

	if checkPassword(user.Password, claim.password) {
		jwt, err := generateJWT(user)
		if err != nil {
			logger.Logger.Error("could not login user due to an error: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to authenticate user"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"jwt": jwt, "user": fiber.Map{"username": user.Username, "ts": time.Now().Unix()}})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect username or password"})
}
