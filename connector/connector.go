package connector

import "github.com/Sirupsen/logrus"

type Connector interface {
	Queue(payload []byte)
	Dequeue() []byte
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
