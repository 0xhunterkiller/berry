package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

func generateJWT(user *models.UserModel) (string, error) {
	claims := jwt.MapClaims{
		"iss":    "tranquil-authn",
		"sub":    user.Username,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour).Unix(),
		"userid": user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error generating jwt token > %w", err)
	}
	return signedToken, nil
}

func validateJWT(tokenstr string) (bool, string, error) {
	var userid string = ""
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer("tranquil-authn"))
	if err != nil {
		return false, userid, fmt.Errorf("error parsing jwt: %w", err)
	}

	if !token.Valid {
		return false, userid, fmt.Errorf("error validating jwt: %w", err)
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["userid"] == nil {
			return false, userid, fmt.Errorf("userid claim is missing")
		}
		userid = claims["userid"].(string)
		return true, userid, nil
	}
	return false, userid, fmt.Errorf("invalid claims format")
}

func unauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return unauthorizedResponse(c)
	}

	ok, userid, err := validateJWT(tokenString)
	if !ok || err != nil {
		return unauthorizedResponse(c)
	}
	c.Locals("userID", userid)
	return c.Next()
}
