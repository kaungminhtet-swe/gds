package singlyLinkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAll(t *testing.T) {
	instances := []int{10, 20, 30, 40, 50}

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

func TestInsertIndex(t *testing.T) {
	t.Run("Inserting value at index 0 of empty list", func(t *testing.T) {
		list := New[int]()
		err := list.InsertAt(0, 10)
		assert.Nil(t, err)
		assert.Equal(t, 1, list.Len(), "Length must equal")
	})

	t.Run("Inserting value at index 1 of empty list", func(t *testing.T) {
		list := New[int]()
		err := list.InsertAt(1, 10)
		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Inserting value at index 1 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30}

		list.InsertAll(instances...)
		err := list.InsertAt(1, 40)

		assert.Nil(t, err)
		assert.Equal(t, 4, list.Len(), "Length must equal 4")

		results := []int{10, 40, 20, 30}
		for index, value := range list.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Inserting value at index 0 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30}

		list.InsertAll(instances...)
		err := list.InsertAt(0, 40)

		assert.Nil(t, err)
		assert.Equal(t, 4, list.Len(), "Length must equal 4")

		results := []int{40, 10, 20, 30}
		for index, value := range list.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Inserting value at last index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30}

		list.InsertAll(instances...)
		err := list.InsertAt(list.Len(), 40)

		assert.Nil(t, err)
		assert.Equal(t, 4, list.Len(), "Length must equal 4")

		results := []int{10, 20, 30, 40}
		for index, value := range list.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Inserting value at out of range index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30}

		list.InsertAll(instances...)
		err := list.InsertAt(4, 40)

		assert.NotNil(t, err)
		assert.Equal(t, "index out of range", err.Error())
	})
}

func TestRemoveFront(t *testing.T) {
	t.Run("Remove front value of empty list", func(t *testing.T) {
		list := New[int]()

		err := list.RemoveFront()

		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Remover front value of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30}

		list.InsertAll(instances...)
		err := list.RemoveFront()
		assert.Nil(t, err)
		assert.Equal(t, 2, list.Len(), "Length must equal 2")

		result := []int{20, 30}
		for i, value := range list.Iterate() {
			assert.Equal(t, result[i], value)
		}
	})
}

func TestRemoveBack(t *testing.T) {
	t.Run("Remove back value of empty list", func(t *testing.T) {
		list := New[int]()

		err := list.RemoveBack()

		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Remover back value of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40}

		list.InsertAll(instances...)
		err := list.RemoveBack()
		assert.Nil(t, err)
		assert.Equal(t, 3, list.Len(), "Length must equal 3")

		result := []int{10, 20, 30}
		for i, value := range list.Iterate() {
			assert.Equal(t, result[i], value)
		}
	})

}

func TestRemoveAt(t *testing.T) {
	t.Run("Remove value at index 0 of empty list", func(t *testing.T) {
		list := New[int]()

		err := list.RemoveAt(0)

		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Remove value at index 0 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40}

		list.InsertAll(instances...)
		err := list.RemoveAt(0)
		assert.Nil(t, err)
		assert.Equal(t, 3, list.Len(), "Length must equal 3")

		result := []int{20, 30, 40}
		for i, value := range list.Iterate() {
			assert.Equal(t, result[i], value)
		}
	})

	t.Run("Remove value at index 2 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40}

		list.InsertAll(instances...)
		err := list.RemoveAt(2)
		assert.Nil(t, err)
		assert.Equal(t, 3, list.Len(), "Length must equal 3")

		result := []int{10, 20, 40}
		for i, value := range list.Iterate() {
			assert.Equal(t, result[i], value)
		}
	})

	t.Run("Remove value at last index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40}

		list.InsertAll(instances...)
		err := list.RemoveAt(list.Len() - 1)
		assert.Nil(t, err)
		assert.Equal(t, 3, list.Len(), "Length must equal 3")

		result := []int{10, 20, 30}
		for i, value := range list.Iterate() {
			assert.Equal(t, result[i], value)
		}
	})

	t.Run("Remove value at out of range index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40}

		list.InsertAll(instances...)
		err := list.RemoveAt(4)
		assert.NotNil(t, err)
		assert.Equal(t, "index out of range", err.Error())
	})
}

func TestGet(t *testing.T) {
	t.Run("Get value at index 0 of empty list", func(t *testing.T) {
		list := New[int]()
		value, err := list.Get(0)
		assert.Equal(t, "empty list", err.Error())
		assert.Equal(t, 0, value)
	})

	t.Run("Get value out of range index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30}

		list.InsertAll(instances...)
		value, err := list.Get(4)
		assert.NotNil(t, err)
		assert.Equal(t, "index out of range", err.Error())
		assert.Equal(t, 0, value)
	})

	t.Run("Get value at range index of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40}
		list.InsertAll(instances...)

		for i, ins := range instances {
			value, err := list.Get(i)
			assert.Nil(t, err)
			assert.Equal(t, ins, value)
		}
	})
}

func TestAppend(t *testing.T) {
	t.Run("Append non-empty list to empty list", func(t *testing.T) {
		instances := []int{10, 20, 30, 40, 50, 60}
		list1 := New[int]()
		list1.InsertAll(instances...)

		list2 := New[int]()
		list2.Append(list1)

		assert.Equal(t, 6, list1.Len())
		for i, value := range list2.Iterate() {
			assert.Equal(t, instances[i], value)
		}
	})

	t.Run("Append non-empty list to non-empty list", func(t *testing.T) {
		instances := []int{10, 20, 30, 40, 50, 60}
		list1 := New[int]()
		list1.InsertAll(10, 20, 30)

		list2 := New[int]()
		list2.InsertAll(40, 50, 60)

		list1.Append(list2)
		assert.Equal(t, 6, list1.Len())

		for i, value := range list1.Iterate() {
			assert.Equal(t, instances[i], value)
		}
	})
}
