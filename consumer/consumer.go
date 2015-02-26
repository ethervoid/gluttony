package consumer

import (
	"github.com/Sirupsen/logrus"
	"strixhq.com/gluttony/connector"
	"strixhq.com/gluttony/task"
)

type Consumer struct {
	connector   connector.Connector
	taskFactory *task.TaskFactory
}

func New(host string, connectorType string, taskFactory *task.TaskFactory) *Consumer {
	connector := connector.New(connectorType)
	consumer := Consumer{connector, taskFactory}
	connector.Connect(host)

	return &consumer
}

func (consumer *Consumer) Start(queue string) {
	logrus.Info("Starting consumer...")
	delivery := make(chan []byte)
	go consumer.connector.Consume(queue, delivery)
	for {
		message := <-delivery
		logrus.Info("Message receive: ", string(message))
	}
}
