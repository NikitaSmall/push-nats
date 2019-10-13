package interfaces

type IConnect interface {
	Publish(topic string, payload []byte) error
}
