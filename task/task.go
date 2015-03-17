package task

import "github.com/Sirupsen/logrus"

type Task interface {
	Execute() error
	GetTaskData() *TaskData
}

func Retry(task Task) {
	taskData := task.GetTaskData()
	logrus.Infof("Current retries %v --- Max retries %v", taskData.CurrentRetries, taskData.MaxRetries)
	taskData.CurrentRetries += 1
	if taskData.CurrentRetries <= int(taskData.MaxRetries) {
		task.Execute()
	}
}
