package doublylinkedlist

import (
	"errors"
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

	t.Run("Inserting value at index 0 of non-empty list", func(t *testing.T) {
		list := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{25, 10, 20, 30, 40, 50}
		list.InsertAll(instances...)
		err := list.InsertAt(0, 25)
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
		err := list.InsertAt(1, 25)
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
		err := list.InsertAt(list.Len()-1, 60)
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
		err := list.InsertAt(5, 60)
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

func TestRemoveBack(t *testing.T) {
	t.Run("Remove last element of empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		err := l.RemoveBack()
		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Remove last element of non-empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{10, 20, 30, 40}
		l.InsertAll(instances...)
		err := l.RemoveBack()
		assert.Nil(t, err)
		assert.Equal(t, 4, l.Len(), "Length must equal 4")

		for index, value := range l.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})
}

func TestRemoveAt(t *testing.T) {
	t.Run("Remove element at index 0 of empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		err := l.RemoveAt(0)
		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Remove element at index 0 of non-empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{20, 30, 40, 50}
		l.InsertAll(instances...)
		err := l.RemoveAt(0)
		assert.Nil(t, err)
		assert.Equal(t, 4, l.Len(), "Length must equal 4")

		for index, value := range l.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Remove element at middle index of non-empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{10, 20, 40, 50}
		l.InsertAll(instances...)
		err := l.RemoveAt(2)
		assert.Nil(t, err)
		assert.Equal(t, 4, l.Len(), "Length must equal 4")

		for index, value := range l.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Remove element at last index of non-empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		results := []int{10, 20, 30, 40}
		l.InsertAll(instances...)
		err := l.RemoveAt(4)
		assert.Nil(t, err)
		assert.Equal(t, 4, l.Len(), "Length must equal 4")

		for index, value := range l.Iterate() {
			assert.Equal(t, results[index], value)
		}
	})

	t.Run("Remove element at index that is out of bounds of non-empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		l.InsertAll(instances...)
		err := l.RemoveAt(5)
		assert.NotNil(t, err)
		assert.Equal(t, "index out of bounds", err.Error())
	})
}

func TestFront(t *testing.T) {
	testcases := []struct {
		name        string
		instances   []int
		result      int
		expectedErr error
	}{
		{
			"Empty list",
			[]int{},
			0,
			errors.New("empty list"),
		},
		{
			"Non-empty list",
			[]int{10, 20, 30, 40, 50},
			10,
			nil,
		},
	}

	for _, tc := range testcases {
		l := New[int]()
		l.InsertAll(tc.instances...)
		first, err := l.Front()
		assert.Equal(t, tc.result, first)
		assert.Equal(t, l.len, len(tc.instances))

		if err != nil {
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestBack(t *testing.T) {
	testcases := []struct {
		name        string
		instances   []int
		result      int
		expectedErr error
	}{
		{
			"Empty list",
			[]int{},
			0,
			errors.New("empty list"),
		},
		{
			"Non-empty list",
			[]int{10, 20, 30, 40, 50},
			50,
			nil,
		},
	}

	for _, tc := range testcases {
		l := New[int]()
		l.InsertAll(tc.instances...)
		first, err := l.Back()
		assert.Equal(t, tc.result, first)
		assert.Equal(t, l.len, len(tc.instances))

		if err != nil {
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestGet(t *testing.T) {
	t.Run("Get element at index 0 of empty list", func(t *testing.T) {
		t.Parallel()
		l := New[int]()
		_, err := l.Get(0)
		assert.NotNil(t, err)
		assert.Equal(t, "empty list", err.Error())
	})

	t.Run("Get first element of non-empty list", func(t *testing.T) {
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		l.InsertAll(instances...)
		first, err := l.Get(0)
		assert.Nil(t, err)
		assert.Equal(t, 10, first)
		assert.Equal(t, l.len, len(instances))
	})

	t.Run("Get middle element of non-empty list", func(t *testing.T) {
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		l.InsertAll(instances...)
		midEle, err := l.Get(2)
		assert.Nil(t, err)
		assert.Equal(t, 30, midEle)
		assert.Equal(t, l.len, len(instances))
	})

	t.Run("Get last element of non-empty list", func(t *testing.T) {
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		l.InsertAll(instances...)
		last, err := l.Get(len(instances) - 1)
		assert.Nil(t, err)
		assert.Equal(t, 50, last)
		assert.Equal(t, l.len, len(instances))
	})

	t.Run("Get element at index that is out of bounds", func(t *testing.T) {
		l := New[int]()
		instances := []int{10, 20, 30, 40, 50}
		l.InsertAll(instances...)
		_, err := l.Get(len(instances))
		assert.NotNil(t, err)
		assert.Equal(t, "index out of bounds", err.Error())
	})
}
