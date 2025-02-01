package singlyLinkedList

import (
	"errors"
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

func TestInsertAt(t *testing.T) {
	testcases := []struct {
		name      string
		instances []int
		result    []int
		value     int
		index     int
		wantedErr error
	}{
		{
			"Inserting value at index 0 of empty list",
			[]int{},
			[]int{},
			10,
			0,
			errors.New("empty list"),
		},
		{
			"Inserting value at index 0 of non-empty list",
			[]int{20, 30, 40},
			[]int{10, 20, 30, 40},
			10,
			0,
			nil,
		},
		{
			"Inserting value in the middle of non-empty list",
			[]int{10, 20, 30},
			[]int{10, 40, 20, 30},
			40,
			1,
			nil,
		},
		{
			"Inserting value at last index of non-empty list",
			[]int{10, 20, 30},
			[]int{10, 20, 40, 30},
			40,
			2,
			nil,
		},
		{
			"Inserting value at out of bounds index of non-empty list",
			[]int{10, 20, 30},
			[]int{},
			40,
			4,
			errors.New("index out of bounds"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			list := New[int]()
			if len(tc.instances) > 0 {
				list.InsertAll(tc.instances...)
			}
			err := list.InsertAt(tc.index, tc.value)

			if tc.wantedErr != nil {
				assert.Equal(t, tc.wantedErr.Error(), err.Error())
			} else {
				for i, value := range list.Iterate() {
					assert.Equal(t, tc.result[i], value)
				}
				assert.Equal(t, len(tc.result), list.Len())
			}
		})
	}
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
	testcases := []struct {
		name      string
		instances []int
		result    []int
		index     int
		wantedErr error
	}{
		{
			name:      "Remove first value of empty list",
			instances: []int{},
			index:     0,
			result:    []int{},
			wantedErr: errors.New("empty list"),
		},
		{
			name:      "Remove first value of non-empty list",
			instances: []int{10, 20, 30, 40, 50},
			index:     0,
			result:    []int{20, 30, 40, 50},
			wantedErr: nil,
		},
		{
			name:      "Remove middle value of non-empty list",
			instances: []int{10, 20, 30, 40, 50},
			index:     2,
			result:    []int{10, 20, 40, 50},
			wantedErr: nil,
		},
		{
			name:      "Remove last value of non-empty list",
			instances: []int{10, 20, 30, 40, 50},
			index:     4,
			result:    []int{10, 20, 30, 40},
			wantedErr: nil,
		},
		{
			name:      "Remove out of bounds value of non-empty list",
			instances: []int{10, 20, 30, 40, 50},
			index:     5,
			result:    []int{},
			wantedErr: errors.New("index out of bounds"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			list := New[int]()
			list.InsertAll(tc.instances...)
			err := list.RemoveAt(tc.index)

			if tc.wantedErr != nil {
				assert.Equal(t, tc.wantedErr.Error(), err.Error())
			} else {
				for i, value := range list.Iterate() {
					assert.Equal(t, tc.result[i], value)
				}
				assert.Equal(t, len(tc.result), list.Len())
			}
		})
	}
}

func TestGet(t *testing.T) {
	testcases := []struct {
		name      string
		instances []int
		index     int
		result    int
		wantedErr error
	}{
		{
			name:      "Get value at index 0 of empty list",
			instances: []int{},
			index:     0,
			result:    0,
			wantedErr: errors.New("empty list"),
		},
		{
			name:      "Get value out of bounds index of non-empty list",
			instances: []int{10, 20, 30},
			index:     4,
			result:    0,
			wantedErr: errors.New("index out of bounds"),
		},
		{
			name:      "Get frist of non-empty list",
			instances: []int{10, 20, 30, 40},
			index:     0,
			result:    10,
			wantedErr: nil,
		},
		{
			name:      "Get last value of non-empty list",
			instances: []int{10, 20, 30, 40},
			index:     3,
			result:    40,
			wantedErr: nil,
		},
		{
			name:      "Get middle value of non-empty",
			instances: []int{10, 20, 30, 40},
			index:     2,
			result:    30,
			wantedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			list := New[int]()
			list.InsertAll(tc.instances...)

			value, err := list.Get(tc.index)
			if tc.wantedErr != nil {
				assert.Equal(t, tc.wantedErr.Error(), err.Error())
			} else {
				assert.Equal(t, tc.result, value)
			}
		})
	}
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
		results := []int{10, 20, 30, 40, 50, 60}
		list1 := New[int]()
		list1.InsertAll(10, 20, 30)

		list2 := New[int]()
		list2.InsertAll(40, 50, 60)

		list1.Append(list2)
		assert.Equal(t, 6, list1.Len())

		for i, value := range list1.Iterate() {
			assert.Equal(t, results[i], value)
		}
	})
}
