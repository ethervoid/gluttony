package task

type TaskFactory interface {
	New(taskType string) Task
}
