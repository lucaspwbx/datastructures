package stack

import "errors"

type Stack []interface{}

var NoItemsErr = errors.New("No items in stack")

func (s *Stack) IsEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (interface{}, error) {
	current := *s
	if r := current.IsEmpty(); r == true {
		return nil, NoItemsErr
	}
	popped := current[len(current)-1]
	current = current[:len(current)-1]
	return popped, nil
}

func (s *Stack) Peek() (interface{}, error) {
	current := *s
	if r := current.IsEmpty(); r == true {
		return nil, NoItemsErr
	}
	return current[len(current)-1], nil
}

func (s *Stack) Size() int {
	return len(*s)
}
