package main

import (
	"fmt"
	"time"

	"github.com/nikitasmall/push-nats/publisher/commands"
	"github.com/nikitasmall/push-nats/publisher/connect"
)

func main() {
	waiter := make(chan struct{})
	time.Sleep(5 * time.Second)
	publishCommand := commands.PublishCommand{
		Connect: connect.NewNats("nats://nats:4222"),
	}

	for !publishCommand.Connect.IsConnected() {
		fmt.Println("connecting...")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("start publisher!")

	fmt.Println("sending!")
	work(publishCommand)
	<-waiter
}
