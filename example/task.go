package task

import (
	"github.com/Sirupsen/logrus"
	"strixhq.com/gluttony/task"
)

type TaskFactory struct {
}

func (factory TaskFactory) New(taskData *task.TaskData) task.Task {
	switch taskData.Id {
	case "alert_email_notification":
		return NewAlertEmailNotification(taskData)
	default:
		logrus.Error("Task not found")
		return nil
	}
}
