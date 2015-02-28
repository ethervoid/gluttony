package connector

import "fmt"

type Connector interface {
	Connect(host string, queues []string) error
	Consume(delivery chan []byte)
}

func New(connectorType string) (Connector, error) {
	switch connectorType {
	case "rabbitmq":
		return NewRabbitMQConnector(), nil
	default:
		return nil, fmt.Errorf("Uknown connector type")
	}
}
