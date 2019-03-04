package xcache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	cache := NewBuilder().
		Capacity(3).
		LRU().
		Build()

	cache.Set(1, 4)
	value, err := cache.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 4, value)

	_, err = cache.Get(2)
	assert.Error(t, err)

	cache.Set(2, 10)
	value, err = cache.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, 10, value)

	cache.Set(3, 12)
	value, err = cache.Get(3)
	assert.Nil(t, err)
	assert.Equal(t, 12, value)

	cache.Set(4, 13)
	value, err = cache.Get(4)
	assert.Nil(t, err)
	assert.Equal(t, 13, value)

	value, err = cache.Get(1)
	assert.NotNil(t, err)
}
