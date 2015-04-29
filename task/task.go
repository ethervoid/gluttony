package task

type Task interface {
	Execute() error
	GetTaskData() *TaskData
}

func Retry(task Task) {
	taskData := task.GetTaskData()
	taskData.CurrentRetries += 1
	if taskData.CurrentRetries <= int(taskData.MaxRetries) {
		task.Execute()
	}
}
