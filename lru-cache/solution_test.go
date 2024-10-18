package lru_cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type op func(*testing.T, *LRUCache)

func put(k, v int) op {
	return func(t *testing.T, c *LRUCache) {
		c.Put(k, v)
	}
}

func get(k, wantV int) op {
	return func(t *testing.T, c *LRUCache) {
		gotV := c.Get(k)
		assert.Equal(t, wantV, gotV)
	}
}

func Test_LRUCache_Get_Put(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []op
	}{
		{
			name:     "cap 1 site case 1",
			capacity: 1,
			operations: []op{
				get(6, -1),
				get(8, -1),
				put(12, 1),
				get(2, -1),
				put(15, 11),
				put(5, 2),
				put(1, 15),
				put(4, 2),
				get(4, 2),
				put(15, 15),
			},
		},
		{
			name:     "cap 1 case 1",
			capacity: 1,
			operations: []op{
				put(1, 10),
				get(1, 10),
			},
		},
		{
			name:     "cap 1 case 2",
			capacity: 1,
			operations: []op{
				put(1, 10),
				get(2, -1),
			},
		},
		{
			name:     "cap 1 case 3",
			capacity: 1,
			operations: []op{
				put(1, 10),
				put(2, 20),
				get(1, -1),
				get(2, 20),
			},
		},
		{
			name:     "cap 2 case 1",
			capacity: 2,
			operations: []op{
				put(1, 10),
				get(1, 10),
				put(2, 20),
				put(3, 30),
				get(1, -1),
				get(2, 20),
				get(3, 30),
			},
		},
		{
			name:     "cap 2 case 2",
			capacity: 2,
			operations: []op{
				put(1, 10),
				put(2, 20),
				get(1, 10),
				put(3, 30),
				get(1, 10),
				get(2, -1),
				get(3, 30),
			},
		},
		{
			name:     "cap 2 site case 1",
			capacity: 2,
			operations: []op{
				put(1, 1),
				put(2, 2),
				get(1, 1),
				put(3, 3),
				get(2, -1),
				put(4, 4),
				get(1, -1),
				get(3, 3),
				get(4, 4),
			},
		},
		{
			name:     "cap 3 case 1",
			capacity: 3,
			operations: []op{
				put(1, 10),
				put(2, 20),
				get(1, 10),
				put(3, 30),
				get(1, 10),
				get(2, 20),
				get(3, 30),
			},
		},
		{
			name:     "cap 3 case 1",
			capacity: 3,
			operations: []op{
				put(1, 10),
				put(2, 20),
				get(1, 10),
				put(3, 30),
				put(4, 40),
				get(1, 10),
				get(2, -1),
				get(3, 30),
				get(4, 40),
			},
		},
		{
			name:     "cap 3 case 1",
			capacity: 3,
			operations: []op{
				put(1, 10),
				put(2, 20),
				get(1, 10),
				put(3, 30),
				put(4, 40),
				put(3, 300),
				get(1, 10),
				get(2, -1),
				get(3, 300),
				get(4, 40),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor(tt.capacity)
			for _, op := range tt.operations {
				op(t, &c)
			}
		})
	}

}
