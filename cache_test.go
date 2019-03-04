package xcache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLRUCache(t *testing.T) {

	buffer := NewLRUBuffer(3)
	cache := NewMemCache(buffer)

	cache.Set(1, 4)
	value, err := cache.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 4, value)
	assert.Equal(t, []Key{1}, buffer.Keys())
	assert.Equal(t, []Value{4}, buffer.Values())

	_, err = cache.Get(2)
	assert.Error(t, err)

	cache.Set(2, 10)
	value, err = cache.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, 10, value)
	assert.Equal(t, []Key{1, 2}, buffer.Keys())
	assert.Equal(t, []Value{4, 10}, buffer.Values())

	cache.Set(3, 12)
	value, err = cache.Get(3)
	assert.Nil(t, err)
	assert.Equal(t, 12, value)
	assert.Equal(t, []Key{1, 2, 3}, buffer.Keys())
	assert.Equal(t, []Value{4, 10, 12}, buffer.Values())

	cache.Set(4, 13)
	value, err = cache.Get(4)
	assert.Nil(t, err)
	assert.Equal(t, 13, value)
	assert.Equal(t, []Key{2, 3, 4}, buffer.Keys())
	assert.Equal(t, []Value{10, 12, 13}, buffer.Values())

	value, err = cache.Get(1)
	assert.NotNil(t, err)

	cache.Set(5, 14)
	value, err = cache.Get(5)
	assert.Nil(t, err)
	assert.Equal(t, 14, value)
	assert.Equal(t, []Key{3, 4, 5}, buffer.Keys())
	assert.Equal(t, []Value{12, 13, 14}, buffer.Values())

	value, err = cache.Get(2)
	assert.NotNil(t, err)

	value, err = cache.Get(3)
	assert.Nil(t, err)
	assert.Equal(t, 12, value)
	assert.Equal(t, []Key{4, 5, 3}, buffer.Keys())
	assert.Equal(t, []Value{13, 14, 12}, buffer.Values())

	value, err = cache.Get(4)
	assert.Nil(t, err)
	assert.Equal(t, 13, value)
	assert.Equal(t, []Key{5, 3, 4}, buffer.Keys())
	assert.Equal(t, []Value{14, 12, 13}, buffer.Values())

	cache.Set(6, 15)
	value, err = cache.Get(6)
	assert.Nil(t, err)
	assert.Equal(t, 15, value)
	assert.Equal(t, []Key{3, 4, 6}, buffer.Keys())
	assert.Equal(t, []Value{12, 13, 15}, buffer.Values())

	value, err = cache.Get(5)
	assert.NotNil(t, err)

	cache.Set(3, 2)
	assert.Equal(t, []Key{3, 4, 6}, buffer.Keys())
	assert.Equal(t, []Value{2, 13, 15}, buffer.Values())
}
