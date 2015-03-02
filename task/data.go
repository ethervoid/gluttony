package task

import (
	"encoding/json"
	"strconv"

	"github.com/Sirupsen/logrus"
)

type TaskData struct {
	Id    string                 `json:"id"`
	Retry float64                `json:"retry"`
	Args  map[string]interface{} `json:"args"`
}

func Unmarshal(data []byte) (*TaskData, error) {
	task := &TaskData{}

	err := json.Unmarshal(data, task)
	if err != nil {
		return nil, err
	}

	logrus.Info("Args: ", task.Args)
	logrus.Info("Unmarshal: ", task.String())
	return task, nil
}

func (taskData *TaskData) String() string {
	output := "{id: " + taskData.Id + ", retry: " + strconv.FormatFloat(taskData.Retry, 'f', 0, 64) + ", args: ["
	for _, arg := range taskData.Args {
		switch t := arg.(type) {
		case string:
			output = output + string(t) + ", "
		case int:
			output = output + strconv.Itoa(t) + ", "
		case float64:
			output = output + strconv.FormatFloat(t, 'f', 2, 64) + ", "
		}
	}

	output = output + "]"

	return output
}
