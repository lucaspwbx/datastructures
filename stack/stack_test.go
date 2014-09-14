package stack

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	stackTests := []struct {
		Stack  Stack
		Result bool
	}{
		{Stack{1, 2}, false},
		{Stack{}, true},
	}

	//standard test
	for _, v := range stackTests {
		if returned := v.Stack.IsEmpty(); returned != v.Result {
			t.Errorf("Error, got %v, should get %v", returned, v.Result)
		}
	}

	//with testify (assert)
	for _, v := range stackTests {
		assert.Equal(t, v.Stack.IsEmpty(), v.Result)
	}
}

func TestPush(t *testing.T) {
	pushTests := []struct {
		Stack  Stack
		Item   int
		Result Stack
	}{
		{Stack{}, 1, Stack{1}},
		{Stack{1, 2}, 3, Stack{1, 2, 3}},
	}

	//standard test
	for _, v := range pushTests {
		v.Stack.Push(v.Item)
		if test := reflect.DeepEqual(v.Stack, v.Result); test != true {
			t.Errorf("Error")
		}
	}

	//with testify
	for _, v := range pushTests {
		v.Stack.Push(v.Item)
		assert.Equal(t, reflect.DeepEqual(v.Stack, v.Result), true)
	}
}

func TestPop(t *testing.T) {
	popTests := []struct {
		Stack  Stack
		Popped interface{}
		Result Stack
		Error  error
	}{
		{Stack{1, 2}, 2, Stack{1}, nil},
		{Stack{}, nil, Stack{}, NoItemsErr},
		{Stack{"go", "rules"}, "rules", Stack{"go"}, nil},
	}

	//standard tests
	for _, v := range popTests {
		//testing popped item
		popped, err := v.Stack.Pop()
		if popped != v.Popped && err != v.Error {
			t.Errorf("Error")
		}
		if err == nil {
			if r := reflect.DeepEqual(v.Stack, v.Result); r == true {
				t.Errorf("Error, stacks should be different")
			}
		}
	}

	//with testify
	for _, v := range popTests {
		popped, err := v.Stack.Pop()
		assert.Equal(t, popped, v.Popped)
		assert.Equal(t, err, v.Error)
		if err == nil {
			assert.NotEqual(t, v.Stack, v.Result)
		}
	}
}

func TestPeek(t *testing.T) {
	peekTests := []struct {
		Stack  Stack
		Peeked interface{}
		Result Stack
		Error  error
	}{
		{Stack{1, 2}, 2, Stack{1, 2}, nil},
		{Stack{"go", "rules"}, "rules", Stack{"go", "rules"}, nil},
		{Stack{}, nil, Stack{}, NoItemsErr},
	}

	//with testify
	for _, v := range peekTests {
		peeked, err := v.Stack.Peek()
		assert.Equal(t, peeked, v.Peeked)
		assert.Equal(t, err, v.Error)
		assert.Equal(t, v.Stack, v.Result)
	}
}

func TestSize(t *testing.T) {
	sizeTests := []struct {
		Stack  Stack
		Result int
	}{
		{Stack{1, 2}, 2},
		{Stack{}, 0},
	}

	for _, v := range sizeTests {
		assert.Equal(t, v.Stack.Size(), v.Result)
	}
}
