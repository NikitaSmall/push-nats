package main

import (
	"fmt"

	"github.com/nikitasmall/push-nats/publisher/commands"
	"github.com/nikitasmall/push-nats/publisher/connect"
)

func main() {
	publishCommand := commands.PublishCommand{
		Connect: connect.NewNats("nats"),
	}

	fmt.Println("start publisher!")
	work(publishCommand)
}
