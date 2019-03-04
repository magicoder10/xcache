package cache

import (
	"github.com/byliuyang/xcache/buffer"
	"github.com/byliuyang/xcache/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	buf := buffer.NewLRU(3)
	cache := NewMemCache(buf)

	cache.Set(1, 4)
	value, err := cache.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 4, value)
	assert.Equal(t, []entity.Key{1}, buf.Keys())
	assert.Equal(t, []entity.Value{4}, buf.Values())

	_, err = cache.Get(2)
	assert.Error(t, err)

	cache.Set(2, 10)
	value, err = cache.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, 10, value)
	assert.Equal(t, []entity.Key{1, 2}, buf.Keys())
	assert.Equal(t, []entity.Value{4, 10}, buf.Values())

	cache.Set(3, 12)
	value, err = cache.Get(3)
	assert.Nil(t, err)
	assert.Equal(t, 12, value)
	assert.Equal(t, []entity.Key{1, 2, 3}, buf.Keys())
	assert.Equal(t, []entity.Value{4, 10, 12}, buf.Values())

	cache.Set(4, 13)
	value, err = cache.Get(4)
	assert.Nil(t, err)
	assert.Equal(t, 13, value)
	assert.Equal(t, []entity.Key{2, 3, 4}, buf.Keys())
	assert.Equal(t, []entity.Value{10, 12, 13}, buf.Values())

	value, err = cache.Get(1)
	assert.NotNil(t, err)

	cache.Set(5, 14)
	value, err = cache.Get(5)
	assert.Nil(t, err)
	assert.Equal(t, 14, value)
	assert.Equal(t, []entity.Key{3, 4, 5}, buf.Keys())
	assert.Equal(t, []entity.Value{12, 13, 14}, buf.Values())

	value, err = cache.Get(2)
	assert.NotNil(t, err)

	value, err = cache.Get(3)
	assert.Nil(t, err)
	assert.Equal(t, 12, value)
	assert.Equal(t, []entity.Key{4, 5, 3}, buf.Keys())
	assert.Equal(t, []entity.Value{13, 14, 12}, buf.Values())

	value, err = cache.Get(4)
	assert.Nil(t, err)
	assert.Equal(t, 13, value)
	assert.Equal(t, []entity.Key{5, 3, 4}, buf.Keys())
	assert.Equal(t, []entity.Value{14, 12, 13}, buf.Values())

	cache.Set(6, 15)
	value, err = cache.Get(6)
	assert.Nil(t, err)
	assert.Equal(t, 15, value)
	assert.Equal(t, []entity.Key{3, 4, 6}, buf.Keys())
	assert.Equal(t, []entity.Value{12, 13, 15}, buf.Values())

	value, err = cache.Get(5)
	assert.NotNil(t, err)

	cache.Set(3, 2)
	assert.Equal(t, []entity.Key{3, 4, 6}, buf.Keys())
	assert.Equal(t, []entity.Value{2, 13, 15}, buf.Values())
}
