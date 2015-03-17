package connector

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

type rabbitMQConnector struct {
	connData     *ConnectorData
	consumerId   string
	deliveryChan chan []byte
	conn         *amqp.Connection
	channel      *amqp.Channel
}

func NewRabbitMQConnector(connData *ConnectorData) *rabbitMQConnector {
	rabbitMQ := &rabbitMQConnector{}
	rabbitMQ.connData = connData

	return rabbitMQ
}

func (connector *rabbitMQConnector) Connect() error {
	uri := connector.composeUri()
	logrus.Debug("Trying to connect to RabbitMQ to ", uri.String())
	connection, err := amqp.Dial(uri.String())
	if err != nil {
		logrus.Error("Can't connect to RabbitMQ in ", uri.String())
		return err
	}
	connector.conn = connection
	connector.checkForClose()

	channel, err := connection.Channel()
	if err != nil {
		logrus.Error("Can't create a channel to RabbitMQ in ", uri.String())
		return err
	}
	connector.channel = channel

	return nil
}

func (connector *rabbitMQConnector) reconnect() {
	for {
		logrus.Debug("Trying to reconnect...")
		err := connector.Connect()
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
	for _, queue := range connector.connData.Queues {
		logrus.Debug("Listening in queue ", queue)
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

func (connector *rabbitMQConnector) composeUri() amqp.URI {
	uri := amqp.URI{}
	uri.Scheme = connector.connData.Type
	uri.Host = connector.connData.Host
	uri.Port = connector.connData.Port
	uri.Username = connector.connData.User
	uri.Password = connector.connData.Password
	uri.Vhost = connector.connData.Args["vhost"].(string)

	return uri
}

func (connector *rabbitMQConnector) Close() error {
	if err := connector.conn.Close(); err != nil {
		return err
	}

	return nil
}
