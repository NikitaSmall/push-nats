package connect

import (
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
)

// NewNats creates a connection that
func NewNats(url string) *nats.Conn {
	connection, err := nats.Connect(url, nats.Name("API Options Example"), nats.Timeout(10*time.Second))
	if err != nil {
		log.Panic("can't connect to nats due to ", err)
	}

	return connection
}
