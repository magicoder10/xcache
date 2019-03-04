package cache

import (
	"cache/buffer"
	"cache/entity"
	"errors"
)

type Cache interface {
	Set(key entity.Key, value entity.Value)
	Get(key entity.Key) (entity.Value, error)
}

type MemCache struct {
	buffer buffer.Buffer
	Pages  map[entity.Key]*buffer.Page
}

func NewMemCache(buf buffer.Buffer) Cache {
	return MemCache{
		buffer: buf,
		Pages:  make(map[entity.Key]*buffer.Page),
	}
}

func (c MemCache) Set(key entity.Key, value entity.Value) {
	if Page, ok := c.Pages[key]; ok {
		Page.Key = key
		Page.Val = value
		return
	}

	if c.buffer.IsFull() {
		Page := c.buffer.Evict()
		delete(c.Pages, Page.Key)
	}

	Page := c.buffer.Add(key, value)
	c.Pages[key] = Page
}

func (c MemCache) Get(key entity.Key) (entity.Value, error) {
	Page, ok := c.Pages[key]

	if !ok {
		return nil, errors.New("not in cache")
	}

	c.Pages[key] = c.buffer.Access(Page)
	return Page.Val, nil
}
