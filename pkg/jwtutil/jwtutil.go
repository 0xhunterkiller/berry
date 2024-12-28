package jwtutil

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

// GenerateJWT generates a JWT (JSON Web Token) using the provided claims.
//
// Parameters:
//   - claims (*jwt.MapClaims): A pointer to a MapClaims object that holds the claims to include in the JWT.
//     Example claims include "iss" (issuer), "sub" (subject), "iat" (issued at), "exp" (expiration), and custom claims like "userid".
//
// Returns:
// - string: The signed JWT as a string.
// - error: An error if the token generation or signing fails.
func GenerateJWT(claims *jwt.MapClaims) (string, error) {
	// Create a new JWT with the specified claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error generating jwt token > %w", err)
	}
	return signedToken, nil
}

// CheckAndGetJWT validates and parses a JWT (JSON Web Token) string.
//
// Parameters:
// - tokenstr (string): The JWT string to validate and parse.
// - validMethods ([]string): A list of valid signing methods (e.g., ["HS256"], etc.).
// - issuer (string): The expected issuer of the token.
// - claimMustHave (...string): A variadic list of claim keys that must be present in the token.
//
// Returns:
// - bool: Indicates if the token is valid.
// - *jwt.Token: The parsed token if valid.
// - error: An error if the token is invalid or if any required claims are missing.
func CheckAndGetJWT(tokenstr string, validMethods []string, issuer string, claimMustHave ...string) (bool, *jwt.Token, error) {
	// Parse the token with validation options
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	},
		jwt.WithValidMethods(validMethods),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(issuer))
	if err != nil {
		return false, nil, fmt.Errorf("error parsing jwt: %w", err)
	}

	// Ensure the token is valid
	if !token.Valid {
		return false, nil, fmt.Errorf("error validating jwt: %w", err)
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		for _, v := range claimMustHave {
			if claims[v] == nil {
				return false, nil, fmt.Errorf("%v claim is missing", v)
			}
		}
		return true, token, nil
	}
	return false, nil, fmt.Errorf("invalid claims")
}

func GetFromClaims(token *jwt.Token, key string) (string, bool) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		var value string
		if claims[key] == nil {
			return "", false
		}
		value = claims[key].(string)
		return value, true
	}
	return "", false
}
