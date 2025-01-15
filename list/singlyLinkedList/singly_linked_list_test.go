package singlyLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertAll(t *testing.T) {
	list := New[int]()

	list.InsertAll(10)

	assert.Equal(t, 10, list.Len())
}
