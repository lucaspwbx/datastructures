package queue

import "errors"

type Queue []interface{}

var NoItemsErr = errors.New("No items on queue")

func (q *Queue) Enqueue(element interface{}) error {
	*q = append(*q, element)
	return nil
}

func (q *Queue) Dequeue() (interface{}, error) {
	current := *q
	if current.Empty() {
		return nil, NoItemsErr
	}
	dequeued := current[0]
	*q = current[1:len(current)]
	return dequeued, nil
}

func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, NoItemsErr
	}
	elem := *q
	return elem[0], nil
}

func (q *Queue) Back() (interface{}, error) {
	if q.Empty() {
		return nil, NoItemsErr
	}
	elem := *q
	return elem[elem.Size()-1], nil
}

func (q *Queue) Empty() bool {
	if q.Size() == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return len(*q)
}
