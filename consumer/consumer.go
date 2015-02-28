package consumer

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/nu7hatch/gouuid"
	"strixhq.com/gluttony/connector"
	"strixhq.com/gluttony/task"
)

type Consumer struct {
	uuid        *uuid.UUID
	connector   connector.Connector
	taskFactory *task.TaskFactory
}

func New(
	host string, connectorType string,
	taskFactory *task.TaskFactory, queues []string) (*Consumer, error) {

	connector, err := connector.New(connectorType)
	if err != nil {
		logrus.Error("Can't create a connector for type: ", connectorType)
		return nil, err
	}

	uuid, err := uuid.NewV5(uuid.NamespaceURL, []byte(time.Now().String()))
	if err != nil {
		logrus.Fatal("Can't generate an UUID for the consumer")
	}
	consumer := Consumer{uuid, connector, taskFactory}
	connector.Connect(host, queues)

	return &consumer, nil
}

func (consumer *Consumer) Start() {
	logrus.Info("Starting consumer ", consumer.uuid.String())
	delivery := make(chan []byte)
	go consumer.connector.Consume(delivery)
	for {
		message := <-delivery
		logrus.Debugf("Message receive by consume %s with payload %s ", consumer.uuid.String(), string(message))
	}
}
