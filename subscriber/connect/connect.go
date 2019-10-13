package connect

import (
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

// NewNats creates a connection that
func NewNats(url string) *nats.Conn {
	connection, err := nats.Connect(url, nats.Name("API Options Example"), nats.Timeout(10*time.Second))
	if err != nil {
		log.Panic("can't connect to nats due to ", err)
	}

	return connection
}

// NewNatsStream connects
func NewNatsStream(clusterID, clientID, url string) stan.Conn {
	connection, err := stan.Connect(clusterID, clientID, stan.NatsConn(NewNats(url)))
	if err != nil {
		log.Panic("can't connect to nats streaming due to ", err)
	}

	return connection
}
