package connector

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

type rabbitMQConnector struct {
	host         string
	queues       []string
	consumerId   string
	deliveryChan chan []byte
	conn         *amqp.Connection
	channel      *amqp.Channel
}

func NewRabbitMQConnector() *rabbitMQConnector {
	return &rabbitMQConnector{}
}

func (connector *rabbitMQConnector) Connect(host string, queues []string) error {
	connector.host = host
	connector.queues = queues
	logrus.Info("Trying to connect to RabbitMQ to ", host)
	connection, err := amqp.Dial(host)
	if err != nil {
		logrus.Error("Can't connect to RabbitMQ in ", host)
		return err
	}
	connector.conn = connection
	connector.checkForClose()

	channel, err := connection.Channel()
	if err != nil {
		logrus.Error("Can't create a channel to RabbitMQ in ", host)
		return err
	}
	connector.channel = channel

	return nil
}

func (connector *rabbitMQConnector) reconnect() {
	for {
		logrus.Info("Trying to reconnect...")
		err := connector.Connect(connector.host, connector.queues)
		if err != nil {
			//TODO Change to incremental time
			time.Sleep(5 * time.Second)
			continue
		}
		connector.Consume(connector.deliveryChan)
		break
	}
}

func (connector *rabbitMQConnector) Consume(delivery chan []byte) {
	connector.deliveryChan = delivery
	for _, queue := range connector.queues {
		logrus.Info("Start consuming in queue ", queue)
		firehose, err := connector.channel.Consume(
			queue, "", true, false, false, false, nil,
		)
		if err != nil {
			logrus.Error("Can't consume messages: ", err)
		}

		for {
			message := <-firehose
			if connector.consumerId == "" {
				connector.consumerId = message.ConsumerTag
			}
			if message.ConsumerTag != "" {
				delivery <- message.Body
			}
		}
	}
}

func (connector *rabbitMQConnector) checkForClose() {
	go func() {
		closeEvent := <-connector.conn.NotifyClose(make(chan *amqp.Error))
		connector.channel.Cancel(connector.consumerId, false)
		logrus.Error("Connection error: ", closeEvent)
		connector.reconnect()
	}()
}
