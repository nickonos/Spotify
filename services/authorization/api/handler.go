package api

import "github.com/nickonos/Spotify/packages/broker"

type APIHandler struct {
	brk broker.Broker
}

func NewAPIHandler(brk broker.Broker) APIHandler {
	return APIHandler{
		brk,
	}
}

func (api *APIHandler) Subscribe() {

}
