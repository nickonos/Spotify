package api

import (
	"github.com/nats-io/nats.go"
	"github.com/nickonos/Spotify/packages/broker"
	routes "github.com/nickonos/Spotify/packages/routes"
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
	broker.Subscribe(api.broker, func(msg routes.GetID, raw *nats.Msg) {

		id := api.service.GetID()

		broker.Respond(api.broker, routes.ResponseID{
			Id: id,
		}, raw)
	})

	api.service.KeepAlive()
}
