package interfaces

import (
	nats "github.com/nats-io/nats.go"
)

type ISubscriberFlow interface {
	Handle(m *nats.Msg)
}
