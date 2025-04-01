package sortedLinkedList

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"sort"
	"testing"
)

func generateRandoms(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rand.Intn(100)
	}
	return result
}

func isOrder(list *SortedLinkedList[int]) bool {
	firstElement, _ := list.Front()
	log.Println("firstElement", firstElement)
	for _, v := range list.Iterate() {
		if firstElement > v {
			return false
		}
		firstElement = v
	}
	return true
}

func Test_New(t *testing.T) {
	list := New[int]()
	assert.NotNil(t, list)
}

func TestInsert(t *testing.T) {
	instances := generateRandoms(100)
	list := New[int]()

	for _, v := range instances {
		list.Insert(v)
	}

	assert.Equal(t, 100, list.Len())
	assert.True(t, isOrder(list))
}

func TestFront(t *testing.T) {
	instances := generateRandoms(10)

	list := New[int]()
	err := list.InsertAll(instances...)
	assert.Nil(t, err)

	sort.Ints(instances)

	front, err := list.Front()
	assert.Nil(t, err)
	assert.Equal(t, instances[0], front)
}

func TestBack(t *testing.T) {
	instances := generateRandoms(10)

	list := New[int]()
	err := list.InsertAll(instances...)
	assert.Nil(t, err)

	sort.Ints(instances)

	back, err := list.Back()
	assert.Nil(t, err)
	assert.Equal(t, instances[len(instances)-1], back)
}

func TestRemoveFront(t *testing.T) {
	instances := generateRandoms(10)

	list := New[int]()
	err := list.InsertAll(instances...)
	assert.Nil(t, err)

	sort.Ints(instances)

	err = list.RemoveFront()
	assert.Nil(t, err)

	first, _ := list.Front()
	assert.Equal(t, instances[1], first)
	assert.Equal(t, len(instances)-1, list.Len())
}

func TestRemoveBack(t *testing.T) {
	instances := generateRandoms(10)

	list := New[int]()
	err := list.InsertAll(instances...)
	assert.Nil(t, err)

	sort.Ints(instances)

	err = list.RemoveBack()
	assert.Nil(t, err)

	last, _ := list.Back()
	assert.Equal(t, instances[len(instances)-2], last)
	assert.Equal(t, len(instances)-1, list.Len())
}

func TestGet(t *testing.T) {
	instances := generateRandoms(10)

	list := New[int]()
	err := list.InsertAll(instances...)
	assert.Nil(t, err)

	sort.Ints(instances)

	for i := 0; i < len(instances); i++ {
		value, err := list.Get(i)
		assert.Nil(t, err)
		assert.Equal(t, instances[i], value)
	}
}
