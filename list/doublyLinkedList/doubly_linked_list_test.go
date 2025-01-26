package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAll(t *testing.T) {
	instances := []int{10, 20, 30, 40, 50}
	list := New[int]()
	list.InsertAll(instances...)

	assert.Equal(t, 5, list.Len())

	for index, value := range list.Iterate() {
		assert.Equal(t, instances[index], value)
	}
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
