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

func New(host string, connectorType string, taskFactory *task.TaskFactory, queues []string) (*Consumer, error) {
	connector := connector.New(connectorType)
	consumer := Consumer{connector, taskFactory}
	connector.Connect(host, queues)

	return &consumer, nil
}

func (consumer *Consumer) Start() {
	logrus.Info("Starting consumer...")
	delivery := make(chan []byte)
	go consumer.connector.Consume(delivery)
	for {
		message := <-delivery
		logrus.Info("Message receive: ", string(message))
	}
}
