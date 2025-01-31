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

func TestInsertIndex(t *testing.T) {
	t.Run("Inserting value at index 0 of empty list", func(t *testing.T) {
		list := New[int]()
		err := list.InsertAtIndex(0, 10)
		assert.Nil(t, err)
		assert.Equal(t, 1, list.Len(), "Length must equal")
	})

	t.Run("Inserting value at index 1 of empty list", func(t *testing.T) {
		list := New[int]()
		err := list.InsertAtIndex(1, 10)
		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Inserting value at index 0 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{25, 10, 20, 30, 40, 50}
		list.InsertAll(instances...)
		err := list.InsertAtIndex(0, 25)
		assert.Nil(t, err)
		assert.Equal(t, 6, list.Len(), "Length must equal")

		for index, value := range list.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Inserting value at index 1 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{10, 25, 20, 30, 40, 50}
		list.InsertAll(instances...)
		err := list.InsertAtIndex(1, 25)
		assert.Nil(t, err)
		assert.Equal(t, 6, list.Len(), "Length must equal")

		for index, value := range list.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Inserting value at last index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{10, 20, 30, 40, 60, 50}
		list.InsertAll(instances...)
		err := list.InsertAtIndex(list.Len()-1, 60)
		assert.Nil(t, err)
		assert.Equal(t, 6, list.Len(), "Length must equal")

		for index, value := range list.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Inserting value at out of range index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		list.InsertAll(instances...)
		err := list.InsertAtIndex(5, 60)
		assert.NotNil(t, err)
		assert.Equal(t, "index out of range", err.Error())
	})
}

func TestRemoveFront(t *testing.T) {
	t.Run("Remove first element of empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		err := l.RemoveFront()
		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Remove first element of non-empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{20, 30, 40, 50}
		l.InsertAll(instances...)
		err := l.RemoveFront()
		assert.Nil(t, err)
		assert.Equal(t, 4, l.Len(), "Length must equal 4")

		for index, value := range l.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})
}
