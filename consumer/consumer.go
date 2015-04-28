package consumer

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/ethervoid/gluttony/connector"
	"github.com/ethervoid/gluttony/task"
	"github.com/nu7hatch/gouuid"
)

type Consumer struct {
	uuid        *uuid.UUID
	connector   connector.Connector
	taskFactory task.TaskFactory
}

func New(connData *connector.ConnectorData, taskFactory task.TaskFactory) (*Consumer, error) {
	connector, err := connector.New(connData)
	if err != nil {
		logrus.Error("Can't create a connector for type: ", connData.Type)
		return nil, err
	}

	uuid, err := uuid.NewV5(uuid.NamespaceURL, []byte(time.Now().String()))
	if err != nil {
		logrus.Fatal("Can't generate an UUID for the consumer")
	}
	consumer := Consumer{uuid, connector, taskFactory}
	connerr := connector.Connect()
	if connerr != nil {
		logrus.Error("Error connecting to broker")
		return nil, connerr
	}

	return &consumer, nil
}

func (consumer *Consumer) Start() {
	logrus.Info("Starting consumer ", consumer.uuid.String())
	delivery := make(chan []byte)
	go consumer.connector.Consume(delivery)
	for {
		//TODO: Check if this could mix messages because is a slice (pointer)
		//instead a string (copy)
		message := <-delivery
		go consumer.executeTask(message)
	}
}

func (consumer *Consumer) Close() {
	consumer.connector.Close()
}

func (consumer *Consumer) executeTask(message []byte) {
	taskData, err := consumer.composeTaskData(message)
	if err != nil {
		logrus.Errorf("Can't compose task data: %s", err)
	}
	tsk := consumer.taskFactory.New(taskData)
	// Panic handling. We don't want to shutdown the consumer
	defer func() {
		if e, ok := recover().(error); ok {
			logrus.Errorf("Panic executing task: %v", e)
		}
	}()
	for i := 0; i < int(taskData.MaxRetries); i++ {
		if err := tsk.Execute(); err != nil {
			logrus.Errorf("Error executing task. Retrying..%s", err.Error())
			logrus.Infof("Current: %v --- Max: %v", taskData.CurrentRetries, taskData.MaxRetries)
			if int(taskData.CurrentRetries) == (int(taskData.MaxRetries) - 1) {
				logrus.Errorf("Max number of retries reached!")
			} else {
				taskData.CurrentRetries += 1
				time.Sleep(time.Duration(taskData.RetryTime) * time.Second)
			}
		} else {
			break
		}
	}
}

func (consumer *Consumer) composeTaskData(message []byte) (*task.TaskData, error) {
	logrus.Debugf("Received message : %s", string(message))
	taskData, err := task.Unmarshal(message)
	if err != nil {
		return nil, err
	} else {
		logrus.Debugf(
			"Message receive by consume %s with payload %s ",
			consumer.uuid.String(),
			taskData.String(),
		)

		return taskData, nil
	}
}
