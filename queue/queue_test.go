package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Operations(t *testing.T) {
	tests := []struct {
		name     string
		ops      func(q Queue[int])
		wantSize int
	}{
		{
			name: "enqueue operations",
			ops: func(q Queue[int]) {
				q.Enqueue(10)
				q.Enqueue(20)
				q.Enqueue(30)
			},
			wantSize: 3,
		},
		{
			name: "dequeue operations",
			ops: func(q Queue[int]) {
				q.Enqueue(10)
				q.Enqueue(20)
				v1, _ := q.Dequeue()
				assert.Equal(t, 10, v1)
				v2, _ := q.Dequeue()
				assert.Equal(t, 20, v2)
			},
			wantSize: 0,
		},
		{
			name: "empty queue operations",
			ops: func(q Queue[int]) {
				_, err := q.Dequeue()
				assert.NotNil(t, err)
				_, err = q.Peek()
				assert.NotNil(t, err)
			},
			wantSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New[int]()
			tt.ops(q)
			assert.Equal(t, tt.wantSize, q.Size())
		})
	}
}

func TestQueue_FIFO(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "FIFO order test",
			input:    []string{"first", "second", "third"},
			expected: []string{"first", "second", "third"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New[string]()
			for _, v := range tt.input {
				q.Enqueue(v)
			}

			for _, want := range tt.expected {
				got, err := q.Dequeue()
				assert.Nil(t, err)
				assert.Equal(t, want, got)
			}
		})
	}
}
