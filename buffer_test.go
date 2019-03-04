package xcache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU_Add(t *testing.T) {
	buffer := NewLRUBuffer(3)
	assert.Equal(t, 0, buffer.Size())

	buffer.Add(1, 10)
	assert.Equal(t, 1, buffer.Size())
}

func TestLRU_IsFull(t *testing.T) {
	buffer := NewLRUBuffer(2)

	assert.False(t, buffer.IsFull())
	assert.Equal(t, 0, buffer.Size())

	buffer.Add(1, 10)
	assert.False(t, buffer.IsFull())
	assert.Equal(t, 1, buffer.Size())

	buffer.Add(2, 10)
	assert.True(t, buffer.IsFull())
	assert.Equal(t, 2, buffer.Size())
}

func TestLRU_Evict(t *testing.T) {
	buffer := NewLRUBuffer(2)

	assert.Equal(t, 0, buffer.Size())

	buffer.Add(1, 10)
	assert.Equal(t, 1, buffer.Size())

	buffer.Add(2, 11)
	assert.Equal(t, 2, buffer.Size())

	block := buffer.Evict()
	assert.Equal(t, 1, buffer.Size())
	assert.Equal(t, 1, block.key)
}

func TestLRU_Remove(t *testing.T) {
	buffer := NewLRUBuffer(3)

	assert.False(t, buffer.IsFull())
	assert.Equal(t, 0, buffer.Size())

	buffer.Add(1, 10)
	assert.Equal(t, 1, buffer.Size())

	block := buffer.Add(2, 11)
	assert.Equal(t, 2, buffer.Size())

	buffer.Add(3, 12)
	assert.Equal(t, 3, buffer.Size())

	assert.True(t, buffer.Has(block))
	buffer.Remove(block)
	assert.Equal(t, 2, buffer.Size())
	assert.False(t, buffer.Has(block))
}
