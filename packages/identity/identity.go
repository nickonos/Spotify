package identity

import (
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/routes"
)

type IdentityHelper struct {
	broker *broker.Broker
}

func NewIdentityHelper(brk *broker.Broker) IdentityHelper {
	return IdentityHelper{
		broker: brk,
	}
}

func (instance *IdentityHelper) NewID() (int64, error) {
	var resp routes.ResponseID
	err := broker.Request(instance.broker, routes.GetID{}, &resp)

	return resp.Id, err
}
