package main

import (
	"fmt"
	"time"

	"github.com/nikitasmall/push-nats/subscriber/connect"
	"github.com/nikitasmall/push-nats/subscriber/flows"
)

func main() {
	time.Sleep(2 * time.Second)
	ns := connect.NewNats("nats://nats:4222")

	waiter := make(chan struct{})
	ns.Subscribe("test", flows.SubscriberFlow{waiter}.Handle)
	fmt.Println("start subscriber!")
	<-waiter
}
