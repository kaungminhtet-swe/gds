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
