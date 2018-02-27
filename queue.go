package queue

import (
	"github.com/satori/go.uuid"
)

// Backend defines interface to manage queues. ChannelBackend is default.
type Backend interface {
	Put(queueName string, value interface{}) error
	Get(queueName string, value interface{}) error
	RemoveQueue(queueName string) error
}

var b Backend = NewChannelBackend()

// Use sets backend to manage queues. ChannelBackend is default.
func Use(value Backend) {
	b = value
}

// Put adds value to the end of a queue.
func Put(queueName string, value interface{}) error {
	return b.Put(queueName, value)
}

// Get removes the first element from a queue and put it in the value pointed to by v
func Get(queueName string, v interface{}) error {
	return b.Get(queueName, v)
}

func RemoveQueue(queueName string) error {
	return b.RemoveQueue(queueName)
}

func uniqueId() (string, error) {
	var err error

	for i := 0; i < 10; i++ {
		r, err := uuid.NewV4()

		if err == nil {
			return r.String(), nil
		}
	}

	return "", err
}
