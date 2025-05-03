package test

import (
	"math/rand"

	"github.com/kmh-io/gds/list"
)

func GenerateRandomInts(length int) []int {
	result := make([]int, length, length)
	for i := range length {
		result[i] = rand.Intn(100)
	}
	return result
}

func IsOrderList(l list.List[int]) bool {
	firstElement, _ := l.Front()

	for _, v := range l.Iterate() {
		if firstElement > v {
			return false
		}
		firstElement = v
	}

	return true
}
