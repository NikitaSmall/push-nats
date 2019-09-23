package main

import "github.com/nikitasmall/push-nats/publisher/commands"

// TODO: provide an actual `controller`
func work(publishCommand commands.PublishCommand) {
	publishCommand.Publish("test", []byte("hello world"))
}
