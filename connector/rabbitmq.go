package connector

import "github.com/Sirupsen/logrus"

type rabbitMQConnector struct {
}

func NewRabbitMQConnector() *rabbitMQConnector {
	return &rabbitMQConnector{}
}

func (connector *rabbitMQConnector) Queue(payload []byte) {
	logrus.Info("Queuing...")
}

func (connector *rabbitMQConnector) Dequeue() []byte {
	logrus.Info("Dequeuing...")
	return nil
}
