package xcache

import "errors"

type Key interface{}
type Value interface{}

type Cache interface {
	Set(key Key, value Value)
	Get(key Key) (Value, error)
}

type MemCache struct {
	buffer Buffer
	blocks map[Key]*Block
}

func NewMemCache(buffer Buffer) Cache {
	return MemCache{
		buffer: buffer,
		blocks: make(map[Key]*Block),
	}
}

func (c MemCache) Set(key Key, value Value) {
	if block, ok := c.blocks[key]; ok {
		block.key = key
		block.val = value
		return
	}

	if c.buffer.IsFull() {
		block := c.buffer.Evict()
		delete(c.blocks, block.key)
	}

	block := c.buffer.Add(key, value)
	c.blocks[key] = block
}

func (c MemCache) Get(key Key) (Value, error) {
	block, ok := c.blocks[key]

	if !ok {
		return nil, errors.New("not in cache")
	}

	c.blocks[key] = c.buffer.Access(block)
	return block.val, nil
}

type Pair struct {
	Key
	Value
}
