package gluttony

import (
	"github.com/Sirupsen/logrus"
	"strixhq.com/gluttony/consumer"
	"strixhq.com/gluttony/task"
)

type gluttony struct {
	host          string
	connectorType string
	queues        []string
	taskFactory   *task.TaskFactory
}

func New(host string, connectorType string, queues []string) *gluttony {
	if len(queues) == 0 {
		logrus.Fatal("You must pass at least one queue")
	}

	if host == "" {
		logrus.Fatal("Host can't be empty")
	}

	if connectorType == "" {
		logrus.Fatal("Connector type can't be empty")
	}

	client := gluttony{}
	client.host = host
	client.connectorType = connectorType
	client.queues = queues

	return &client
}

func (gluttony *gluttony) RegisterJobsFactory(taskFactory task.TaskFactory) {
	gluttony.taskFactory = &taskFactory
}

func (gluttony *gluttony) Start(consumers int) {
	consumer, err := consumer.New(
		gluttony.host,
		gluttony.connectorType,
		gluttony.taskFactory,
		gluttony.queues,
	)

	if err != nil {
		logrus.Fatal("Error trying to create consumer")
	}

	consumer.Start()
}
