package routes

import (
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
)

type ServiceInstance struct {
	*logging.Logger
	broker *broker.Broker
	name   string
}

type GetID struct {
}

func (id GetID) Subject() string {
	return "response.Id"
}

type ResponseID struct {
	Id int64 `json:"id"`
}

func (id ResponseID) Subject() string {
	return "response.Id"
}
