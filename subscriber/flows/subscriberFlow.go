package flows

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

type SubscriberFlow struct {
	Waiter chan struct{}
}

func (SubscriberFlow) Handle(m *stan.Msg) {
	log.Printf("Received a message: %s\n", string(m.Data))
}
