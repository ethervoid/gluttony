package gluttony

import (
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
	consumer := consumer.New(gluttony.host, "rabbitmq", gluttony.taskFactory)
	consumer.Start("strix")
}
