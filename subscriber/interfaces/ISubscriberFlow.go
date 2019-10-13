package interfaces

import stan "github.com/nats-io/stan.go"

type ISubscriberFlow interface {
	Handle(m *stan.Msg)
}
