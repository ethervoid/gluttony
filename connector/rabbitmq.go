package connector

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

type rabbitMQConnector struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQConnector() *rabbitMQConnector {
	return &rabbitMQConnector{}
}

func (connector *rabbitMQConnector) Connect(host string) {
	logrus.Info("Trying to connect to RabbitMQ to ", host)
	connection, err := amqp.Dial(host)
	if err != nil {
		logrus.Fatal("Can't connect to RabbitMQ in ", host)
	}
	connector.conn = connection

	channel, err := connection.Channel()
	if err != nil {
		logrus.Fatal("Can't create a channel to RabbitMQ in ", host)
	}
	connector.channel = channel

}

func (connector *rabbitMQConnector) GetQueueInfo(queueName string) {
	queue, error := connector.channel.QueueInspect(queueName)
	if error != nil {
		fmt.Println(error)
		logrus.Error("Can't check the queue state for ", queueName)
	}

	fmt.Println(queue)
}

func (connector *rabbitMQConnector) Queue(payload []byte) {
	logrus.Debug("Queuing...")
}

func (connector *rabbitMQConnector) Dequeue() []byte {
	logrus.Debug("Dequeuing...")
	return nil
}

func (connector *rabbitMQConnector) Consume(queue string, delivery chan []byte) {
	logrus.Info("Start consuming...")
	firehose, err := connector.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		logrus.Error("Can't consume messages: ", err)
	}

	for {
		message := <-firehose
		//TODO Procesar mensaje
		delivery <- message.Body
	}

}
