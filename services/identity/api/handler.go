package api

import (
	"context"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/routes"
	"github.com/nickonos/Spotify/services/identity/service"
)

type APIHandler struct {
	service *service.IdentityService
	broker  *broker.Broker
}

func NewAPIHandler(s *service.IdentityService, b *broker.Broker) APIHandler {
	return APIHandler{
		service: s,
		broker:  b,
	}
}

func (api *APIHandler) Subscribe() {
	broker.Subscribe(api.broker, func(ctx context.Context, msg routes.GetID) (routes.ResponseID, error) {

		id := api.service.GetID()

		return routes.ResponseID{
			Id: id,
		}, nil
	})

	api.service.UpdateTTL()
}
