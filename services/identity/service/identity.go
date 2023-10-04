package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/logging"
)

type IdentityService struct {
	logger   *logging.Logger
	broker   *broker.Broker
	identity int
	node     *snowflake.Node
	current  uint64
}

func NewIdentiyService(b *broker.Broker) (*IdentityService, error) {
	for x := 0; x < 1024; x++ {
		var updatedAt int64
		err := broker.GetKeyValue(b, strconv.Itoa(x), &updatedAt)
		if err != nil && updatedAt > 0 {
			continue
		}

		revision, err := broker.CreateKeyValue(b, strconv.Itoa(x), time.Now().Unix())
		if err != nil {
			continue
		}

		node, err := snowflake.NewNode(int64(x))
		if err != nil {
			return nil, err
		}

		return &IdentityService{
			node:     node,
			identity: x,
			current:  revision,
			logger:   logging.NewLogger("Identity"),
			broker:   b,
		}, nil
	}

	return nil, errors.New("no free node id")
}

func (s *IdentityService) KeepAlive() {
	s.logger.Trace("Started Identity service")

	// Start a goroutine to refresh the TTL of the node id every 5 minutes
	go func() {
		for {
			revision, err := broker.UpdateKeyValue(s.broker, strconv.Itoa(s.identity), time.Now().Unix(), s.current)
			if err != nil {
				s.logger.Fatal(err)
			}
			s.current = revision
			time.Sleep(time.Minute * 5)
		}
	}()
}

func (s *IdentityService) GetID() int64 {
	return s.node.Generate().Int64()
}
