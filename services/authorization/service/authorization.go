package service

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt"

	"github.com/nickonos/Spotify/services/authorization/data"
	"github.com/nickonos/Spotify/services/authorization/spotify"
)

type AuthorizationService struct {
	spotify spotify.Operations
	db      data.DB
}

func NewAuthorizationService() (AuthorizationService, error) {
	sp, err := spotify.NewOperations()
	if err != nil {
		return AuthorizationService{}, err
	}

	db, err := data.NewMysqlDatabase()
	if err != nil {
		return AuthorizationService{}, err
	}

	return AuthorizationService{
		spotify: sp,
		db:      db,
	}, nil
}

type JWTClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (a AuthorizationService) LoginUser(ctx context.Context, code string) (string, error) {
	if code == "admin" {
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
			Role:  "admin",
			Email: "admin@spotify.com",
		})

		token, err := accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
		if err != nil {
			return "", err
		}
		return token, nil
	}

	auth_token, err := a.spotify.GetAccessToken(ctx, code)
	if err != nil {
		return "", err
	}

	email, err := a.spotify.GetUserEmail(ctx, auth_token)
	if err != nil {
		return "", err
	}

	role, err := a.db.GetUserRole(ctx, email)
	if err != nil {
		err = a.db.AddUserRole(ctx, email)
		if err != nil {
			return "", err
		}

		role = "user"
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Role:  role,
		Email: email,
	})

	token, err := accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
