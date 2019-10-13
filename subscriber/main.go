package main

import (
	"fmt"
	"time"

	"github.com/nikitasmall/push-nats/subscriber/connect"
	"github.com/nikitasmall/push-nats/subscriber/flows"
)

func main() {
	time.Sleep(2 * time.Second)
	// ns := connect.NewNats("nats://nats:4222") // nats
	nss := connect.NewNatsStream("test-cluster", "subscriber", "nats://nats:4222") // streams

	waiter := make(chan struct{})
	nss.Subscribe("test", flows.SubscriberFlow{Waiter: waiter}.Handle)
	fmt.Println("start subscriber!")
	<-waiter
}
