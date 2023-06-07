package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/MarioCarrion/todo-api/internal"
)

type Task struct {
	producer  *kafka.producer
	topicName string
}

type event struct {
	Type  string
	Value internal.Task
}

func NewTask(producer *kafka.Producer, topicName string) *Task {
	return &Task{
		topicName: topicName,
		producer:  producer,
	}
}

//Created -> Publishes a message to create a task
func (t *Task) Created(ctx context.Context, task internal.Task) error {
	return t.publish(ctx, "Task.Created", "tasks.events.created", task)
}

//Deleted -> Publishes a message to delete a task
func (t *Task) Deleted(ctx context.Context, task internal.Task) error {
	return t.publish(ctx, "Task.Deleted", "tasks.events.deleted", task)
}

//Updated -> Published a message to update a task
func (t *Task) Updated(ctx context.Context, task internal.Task) error {
	return t.publish(ctx, "Task.Updated", "tasks.events.updated", task)
}

func (t *Task) publish(ctx context.Context, spanName, routingKey string, e interface{}) error {
	var b bytes.Buffer

	evt := event{
		Type:  msgType,
		Value: task,
	}

	_ = json.NewEncoder(&b).Encode(evt)

	_ = t.producer.Producer(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &t.topicName,
			Partition: kafka.PartitionAny,
		},
		Value: b.Bytes(),
	}, nil)

	return nil
}
