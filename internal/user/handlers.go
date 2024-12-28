package user

import (
	"fmt"
	"time"

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
	userService UserServiceIface
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/auth/user/register", uh.registerUser)
	app.Post("/auth/user/login", uh.loginUser)

	app.Use("/user", middleware.AuthMiddleware)
	app.Get("/user/checkauth", uh.amIAuthenticated)
}

// Am I Authenticated

func (uh *UserHandler) amIAuthenticated(c *fiber.Ctx) error {
	user, err := uh.userService.GetByID(c.Locals("userid").(string))
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

	err = uh.userService.CreateUser(req.Username, req.Email, req.Password, true)
	if err != nil {
		logger.Logger.Error("could not create user due to an error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "the user was not created"})
	}
	logger.Logger.Info("the user was created successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "the user was created successfully"})
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
	user, err := uh.userService.GetByUsername(claim.Username)
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
