package tree

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var root *Node[int]

func TestMain(m *testing.M) {
	root = NewNode(10)
	l2 := root.SetLeft(5)
	r3 := root.SetRight(20)

	l2.SetLeft(3)
	l2.SetRight(9)

	r3.SetLeft(15)
	r3.SetRight(25)

	code := m.Run()

	root = nil

	os.Exit(code)
}

func TestInOrder(t *testing.T) {
	assert.NotNil(t, root, "Root should not be nil")
	tree := New(root)

	expected := []int{3, 5, 9, 10, 15, 20, 25}

	for i, actual := range tree.Iterate() {
		assert.Equal(t, expected[i], actual)
	}
}

func TestPreOrder(t *testing.T) {
	assert.NotNil(t, root, "Root should not be nil")
	tree := New(root)
	tree.SetTraverser(RecursivePreOrder)

	expected := []int{10, 5, 3, 9, 20, 15, 25}

	for i, actual := range tree.Iterate() {
		assert.Equal(t, expected[i], actual)
	}
}
