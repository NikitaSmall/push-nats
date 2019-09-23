package connect

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

// NewNats creates a connection that
func NewNats(url string) *nats.Conn {
	connection, err := nats.Connect(url)
	if err != nil {
		log.Panic("can't connect to nats due to ", err)
	}

	return connection
}
