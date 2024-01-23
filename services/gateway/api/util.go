package api

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}

func ValidateRequest(token string, required_role string) error {
	if token == "" {
		return fmt.Errorf("missing required Authorization header")
	}

	claims := JWTClaims{}
	parsed, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if !parsed.Valid {
		return fmt.Errorf("invalid token signature")
	}

	if required_role != "" && claims.Role != required_role {
		return fmt.Errorf("missing role: " + required_role)
	}

	return nil
}
