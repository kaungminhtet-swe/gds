package binarytree

import (
	"sort"
	"testing"

	"github.com/kmh-io/gds/test"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	length := 1000
	inputs := test.GenerateRandomInts(length)
	btree := New[int](inputs[0])

	for _, val := range inputs[1:] {
		btree.Insert(val)
	}

	sort.Ints(inputs)
	for i, v := range btree.InOrder() {
		assert.Equal(t, inputs[i], v)
	}
}
