package task

import (
	"encoding/json"
	"strconv"
)

const DEFAULT_MAX_RETRIES = 10
const DEFAULT_RETRY_TIME = 30

type TaskData struct {
	Id             string  `json:"id"`
	RetryTime      float64 `json:"retry,omitempty"`
	MaxRetries     float64 `json:"max_retries,omitempty"`
	CurrentRetries int
	Args           map[string]interface{} `json:"args"`
}

func Unmarshal(data []byte) (*TaskData, error) {
	task := &TaskData{}
	task.CurrentRetries = 0
	task.RetryTime = DEFAULT_RETRY_TIME
	task.MaxRetries = DEFAULT_MAX_RETRIES

	err := json.Unmarshal(data, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (taskData *TaskData) String() string {
	output := "{id: " + taskData.Id + ", retry: " + strconv.FormatFloat(taskData.RetryTime, 'f', 0, 64) + ", args: ["
	for _, arg := range taskData.Args {
		switch t := arg.(type) {
		case string:
			output = output + string(t) + ", "
		case float64:
			output = output + strconv.FormatFloat(t, 'f', 2, 64) + ", "
		}
	}

	output = output + "]"

	return output
}
