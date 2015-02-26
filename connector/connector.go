package connector

import "github.com/Sirupsen/logrus"

type Connector interface {
	Connect(host string)
	Queue(payload []byte)
	Dequeue() []byte
	GetQueueInfo(queue string)
	Consume(queue string, delivery chan []byte)
}

func New(connectorType string) Connector {
	switch connectorType {
	case "rabbitmq":
		return NewRabbitMQConnector()
	default:
		logrus.Error("Unknown connector type")
		return nil
	}
}
