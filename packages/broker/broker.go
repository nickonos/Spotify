package broker

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/goccy/go-json"

	"github.com/nats-io/nats.go"
	log "github.com/nickonos/Spotify/packages/logging"
)

type Broker struct {
	nc           *nats.Conn
	sessionStore nats.ObjectStore
	keyValue     nats.KeyValue
}

type Response[D any] struct {
	Data D
	Err  string
}

type Message interface {
	Subject() string
}

func newBroker(nc *nats.Conn, js nats.JetStreamContext) *Broker {
	logger := log.NewLogger("broker")

	// Ensure that the bucket for key value storage exists
	store, err := js.CreateObjectStore(&nats.ObjectStoreConfig{
		Bucket:      "account_sessions",
		Description: "",
		TTL:         time.Hour * 2,
	})

	if err != nil {
		logger.Fatal(err)
	}

	kv, new := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket:  "node_ids",
		Storage: nats.MemoryStorage,
		TTL:     time.Minute * 10,
	})

	if new != nil {
		logger.Fatal(err)
	}

	return &Broker{
		nc:           nc,
		sessionStore: store,
		keyValue:     kv,
	}
}

func NewMessageBroker() *Broker {
	logger := log.NewLogger("broker")

	// Get NATS_URL from environment
	natsURL := os.Getenv("NATS_URL")

	// If NATS_URL is not set, use the default
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	// Connect to a server
	nc, err := nats.Connect(natsURL)

	// As the broker is a critical component, we panic if we cannot connect
	if err != nil {
		logger.Fatal(err)
	}

	// JetStream context
	jsOpts := []nats.JSOpt{
		nats.MaxWait(50 * time.Second),
	}
	js, err := nc.JetStream(jsOpts...)

	if err != nil {
		logger.Fatal(err)
	}

	return newBroker(nc, js)
}

func CreateKeyValue[T any](b *Broker, key string, value T) (uint64, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(value)
	if err != nil {
		return 0, err
	}

	return b.keyValue.Create(key, buffer.Bytes())
}

func GetKeyValue[T any](b *Broker, key string, dst *T) error {
	value, err := b.keyValue.Get(key)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(bytes.NewReader(value.Value()))
	err = decoder.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func UpdateKeyValue[T any](b *Broker, key string, value T, last uint64) (uint64, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(value)
	if err != nil {
		return 0, err
	}

	return b.keyValue.Update(key, buffer.Bytes(), last)
}

func Request[M Message, R any](broker *Broker, message M, dst *Response[R]) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	msg, err := broker.nc.Request(fmt.Sprintf("request.%s", message.Subject()), data, time.Second*5)

	if err != nil {
		return err
	}

	var response Response[R]

	err = json.Unmarshal(msg.Data, &response)

	if err != nil {
		return err
	}

	*dst = response

	return nil
}

func Respond[Res Response[D], D any](broker *Broker, message Res, raw *nats.Msg) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	return raw.Respond(data)
}

func Subscribe[Req Message, D any](broker *Broker, cb func(message Req) (D, error)) error {
	logger := log.NewLogger("broker")

	// Create instance of M to get the subject
	var r Req
	subject := r.Subject()

	logger.Trace("Subscribing to " + subject)

	_, err := broker.nc.QueueSubscribe(fmt.Sprintf("request.%s", subject), "job_workers", func(msg *nats.Msg) {
		var message Req

		err := json.Unmarshal(msg.Data, &message)

		if err != nil {
			return
		}

		res, err := cb(message)
		resp := Response[D]{
			Data: res,
		}

		if err != nil {
			resp.Err = err.Error()
		}

		err = Respond(broker, resp, msg)
		if err != nil {
			logger.Trace(err.Error())
		}
	})

	return err
}
