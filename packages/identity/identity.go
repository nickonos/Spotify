package identity

import (
	"errors"

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
	var resp broker.Response[routes.ResponseID]
	err := broker.Request(instance.broker, routes.GetID{}, &resp)
	if err != nil {
		return 0, err
	}

	if resp.Err != "" {
		return 0, errors.New(resp.Err)
	}

	return resp.Data.Id, nil
}
