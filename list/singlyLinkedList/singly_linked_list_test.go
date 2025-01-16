package singlyLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertAll(t *testing.T) {
	instances := []int{10, 20, 30}

	list := New[int]()
	list.InsertAll(instances...)

	for index, value := range list.Iterate() {
		assert.Equal(t, instances[index], value)
	}
}

func TestFront(t *testing.T) {
	list := New[int]()

	value, err := list.Front()
	assert.NotNil(t, err)
	assert.Equal(t, "empty list", err.Error())
	assert.Equal(t, 0, value)

	list.InsertAll(10, 20, 30)
	value, err = list.Front()
	assert.Nil(t, err)
	assert.Equal(t, 10, value)
}

func TestBack(t *testing.T) {
	list := New[int]()

	value, err := list.Back()
	assert.NotNil(t, err)
	assert.Equal(t, "empty list", err.Error())
	assert.Equal(t, 0, value)

	list.InsertAll(10, 20, 30)
	value, err = list.Back()
	assert.Nil(t, err)
	assert.Equal(t, 30, value)
}

func TestInsertFront(t *testing.T) {
	list := New[int]()
	length := 10

	for i := 0; i < length; i++ {
		list.InsertFront(i)
	}

	for index, value := range list.Iterate() {
		assert.Equal(t, length-1-index, value)
	}

	assert.Equal(t, length, list.Len(), "Length must equal ")
}

func TestInsertBack(t *testing.T) {
	list := New[int]()
	length := 10

	for i := 0; i < length; i++ {
		list.InsertBack(i)
	}

	for index, value := range list.Iterate() {
		assert.Equal(t, index, value)
	}

	assert.Equal(t, length, list.Len(), "Length must equal ")
}
