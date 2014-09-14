package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(2)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue, Queue{2})
}

func TestDequeue(t *testing.T) {
	dequeueTests := []struct {
		Queue  Queue
		Item   interface{}
		Result Queue
		Error  error
	}{
		{Queue{1, 2}, 1, Queue{1}, nil},
		{Queue{}, nil, Queue{}, NoItemsErr},
	}
	for _, v := range dequeueTests {
		item, err := v.Queue.Dequeue()
		assert.Equal(t, v.Item, item)
		assert.Equal(t, v.Error, err)
		if err != nil {
			assert.Equal(t, v.Queue, v.Result)
		} else {
			assert.NotEqual(t, v.Queue, v.Result)
		}
	}
}

func TestFront(t *testing.T) {
	frontTests := []struct {
		Queue Queue
		Item  interface{}
		Error error
	}{
		{Queue{1, 2}, 1, nil},
		{Queue{}, nil, NoItemsErr},
	}
	for _, v := range frontTests {
		item, err := v.Queue.Front()
		assert.Equal(t, v.Item, item)
		assert.Equal(t, v.Error, err)
	}
}

func TestBack(t *testing.T) {
	backTests := []struct {
		Queue Queue
		Item  interface{}
		Error error
	}{
		{Queue{1, 2}, 2, nil},
		{Queue{}, nil, NoItemsErr},
	}
	for _, v := range backTests {
		item, err := v.Queue.Back()
		assert.Equal(t, v.Item, item)
		assert.Equal(t, v.Error, err)
	}
}

func TestEmpty(t *testing.T) {
	emptyTests := []struct {
		Queue  Queue
		Result bool
	}{
		{Queue{1, 2}, false},
		{Queue{}, true},
	}
	for _, v := range emptyTests {
		assert.Equal(t, v.Queue.Empty(), v.Result)
	}
}

func TestSize(t *testing.T) {
	queue := Queue{2}
	assert.Equal(t, queue.Size(), 1)
}
