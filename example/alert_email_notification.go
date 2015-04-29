package task

import (
	"github.com/Sirupsen/logrus"
	"strixhq.com/gluttony/task"
)

type alertEmailNotification struct {
	data *task.TaskData
}

func NewAlertEmailNotification(taskData *task.TaskData) *alertEmailNotification {
	logrus.Info(taskData)
	task := &alertEmailNotification{}
	task.data = taskData

	return task
}

func (alertTask *alertEmailNotification) Execute() error {
	//Logic to send the email notification
}

func (alertTask *alertEmailNotification) GetTaskData() *task.TaskData {
	return alertTask.data
}
