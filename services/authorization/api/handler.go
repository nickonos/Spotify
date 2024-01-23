package api

import (
	"context"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/routes"
	"github.com/nickonos/Spotify/services/authorization/service"
)

type APIHandler struct {
	brk     broker.Broker
	service service.AuthorizationService
}

func NewAPIHandler(brk broker.Broker) (APIHandler, error) {
	srv, err := service.NewAuthorizationService()
	if err != nil {
		return APIHandler{}, err
	}

	return APIHandler{
		brk:     brk,
		service: srv,
	}, nil
}

func (api *APIHandler) Subscribe() {
	broker.Subscribe(&api.brk, func(ctx context.Context, message routes.Login) (routes.LoginResponse, error) {
		jwt, err := api.service.LoginUser(ctx, message.Code)
		if err != nil {
			return routes.LoginResponse{}, err
		}

		return routes.LoginResponse{
			JWT: jwt,
		}, nil
	})
}
