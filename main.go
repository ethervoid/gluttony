package gluttony

import (
	"github.com/Sirupsen/logrus"
	"strixhq.com/gluttony/consumer"
	"strixhq.com/gluttony/task"
)

type Gluttony struct {
	host        string
	taskFactory *task.TaskFactory
}

func New(host string, taskFactory task.TaskFactory) *Gluttony {
	client := Gluttony{}
	client.host = host
	client.taskFactory = &taskFactory

	return &client
}

func (gluttony *Gluttony) Start(consumers int) {
	queues := []string{"strix"}
	consumer, err := consumer.New(
		gluttony.host,
		"rabbitmq",
		gluttony.taskFactory,
		queues,
	)

	if err != nil {
		logrus.Fatal("Error trying to create consumer")
	}

	consumer.Start()
}
