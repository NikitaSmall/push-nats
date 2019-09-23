package commands

import (
	"log"

	nats "github.com/nats-io/nats.go"
	"github.com/nikitasmall/push-nats/publisher/interfaces"
)

type PublishCommand struct {
	Connect *nats.Conn
}

func (pc PublishCommand) Publish(topic string, message []byte) {
	if err := pc.Connect.Publish(topic, message); err != nil {
		log.Printf("error during publish command: %s", err)
	}
}

var _ interfaces.IPublishCommand = (*PublishCommand)(nil)
