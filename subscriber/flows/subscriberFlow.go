package flows

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

type SubscriberFlow struct {
	Waiter chan struct{}
}

func (SubscriberFlow) Handle(m *nats.Msg) {
	log.Printf("Received a message: %s\n", string(m.Data))
}
