package interfaces

type IPublishCommand interface {
	Publish(topic string, message []byte)
}
