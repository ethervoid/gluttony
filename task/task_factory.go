package task

type TaskFactory interface {
	New(taskData *TaskData) Task
}
