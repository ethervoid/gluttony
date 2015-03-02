package task

import (
	"bytes"
	"encoding/json"
)

type TaskData struct {
	id    string        `json:"id"`
	retry int           `json:"retry_time"`
	args  []interface{} `json:"args"`
}

func Unmarshal(data []byte) (*TaskData, error) {
	task := &TaskData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&task); err != nil {
		return nil, err
	}
	return task, nil
}

func (taskData *TaskData) String() string {
	output := "{id: " + taskData.id + ", retry: " + string(taskData.retry) + ", args: ["
	for _, arg := range taskData.args {
		if str, ok := arg.(string); ok {
			output = output + string(str) + ", "
		}
	}

	output = output + "]"

	return output
}

func (taskData *TaskData) Id() string {
	return taskData.id
}
