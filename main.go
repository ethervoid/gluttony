package gluttony

import (
	"strixhq.com/gluttony/connector"
	"strixhq.com/gluttony/task"
)

type Gluttony struct {
	conn *connector.Connector
}

func New(host string, taskFactory task.TaskFactory) *Gluttony {
	client := Gluttony{}
	connector.New("rabbitmq")
	//Excample of how to get a task from caller project
	email := taskFactory.New("email_notification")
	email.Execute()

	return &client
}
