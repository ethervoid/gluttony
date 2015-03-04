package connector

import "fmt"

type ConnectorData struct {
	// Args could be used to pass specific connector parameters
	// for example in RabbitMQ you could pass if want to use TLS,SASL or Plain
	// authentication
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Queues   []string
	Args     map[string]interface{}
}

type Connector interface {
	Connect() error
	Consume(delivery chan []byte)
}

func New(connData *ConnectorData) (Connector, error) {
	switch connData.Type {
	case "amqp":
		return NewRabbitMQConnector(connData), nil
	default:
		return nil, fmt.Errorf("Uknown connector type")
	}
}
