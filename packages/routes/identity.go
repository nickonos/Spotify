package routes

import (
	"errors"

	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
)

type ServiceInstance struct {
	*logging.Logger
	broker *broker.Broker
	name   string
}

var exists = false

func NewInstance(name string) *ServiceInstance {
	instance := &ServiceInstance{
		broker: broker.NewMessageBroker(),
		Logger: logging.NewLogger(name),
		name:   name,
	}

	if exists {
		instance.Fatal(errors.New("ServiceInstance already exists, only one can exist per service"))
	}

	exists = true

	instance.Trace("Created ServiceInstance for: %s", name)

	return instance
}

func (instance *ServiceInstance) NewID() (int64, error) {
	var resp ResponseID
	err := broker.Request(instance.broker, GetID{}, &resp)

	return resp.Id, err
}

func (instance *ServiceInstance) Broker() *broker.Broker {
	return instance.broker
}

func (instance *ServiceInstance) Started() {
	instance.Trace("ServiceInstance Started: %s", instance.name)
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
