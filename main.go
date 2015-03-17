package gluttony

import (
	"sync"

	"github.com/Sirupsen/logrus"
	"strixhq.com/gluttony/connector"
	"strixhq.com/gluttony/consumer"
	"strixhq.com/gluttony/task"
)

type gluttony struct {
	connData    *connector.ConnectorData
	taskFactory task.TaskFactory
}

func New(connData *connector.ConnectorData) *gluttony {
	if len(connData.Queues) == 0 {
		logrus.Fatal("You must pass at least one queue")
	}

	if connData.Host == "" {
		logrus.Fatal("Host can't be empty")
	}

	if connData.Type == "" {
		logrus.Fatal("Connector type can't be empty")
	}

	client := gluttony{}
	client.connData = connData

	return &client
}

func (gluttony *gluttony) RegisterJobsFactory(taskFactory task.TaskFactory) {
	gluttony.taskFactory = taskFactory
}

func (gluttony *gluttony) Start(consumers int) {
	var wg sync.WaitGroup
	wg.Add(consumers)
	for i := 0; i < consumers; i++ {
		go func() {
			consumer, err := consumer.New(
				gluttony.connData,
				gluttony.taskFactory,
			)

			if err != nil {
				logrus.Fatal("Error trying to create consumer")
			}

			defer wg.Done()
			defer consumer.Close()
			consumer.Start()
		}()
	}
	wg.Wait()
}
