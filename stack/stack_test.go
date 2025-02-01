package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := New[int]()
	stack.Push(10)
	assert.Equal(t, 1, stack.Size())
	value, err := stack.Seek()
	assert.Nil(t, err)
	assert.Equal(t, 10, value)

	stack.Push(20)
	assert.Equal(t, 2, stack.Size())
	value, err = stack.Seek()
	assert.Nil(t, err)
	assert.Equal(t, 20, value)

	stack.Push(30)
	assert.Equal(t, 3, stack.Size())
	value, err = stack.Seek()
	assert.Nil(t, err)
	assert.Equal(t, 30, value)
}

func TestStack_Pop(t *testing.T) {
	stack := New[int]()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	assert.Equal(t, 3, stack.Size())

	value, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 30, value)
	assert.Equal(t, 2, stack.Size())

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 20, value)
	assert.Equal(t, 1, stack.Size())

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 10, value)
	assert.Equal(t, 0, stack.Size())
}
