package xcache

import (
	"cache/buffer"
	"cache/cache"
)

type ReplacementPolicy int

const (
	LRU ReplacementPolicy = iota
)

type Builder struct {
	capacity int
	policy   ReplacementPolicy
}

func (b *Builder) Build() cache.Cache {
	switch b.policy {
	case LRU:
		buf := buffer.NewLRU(b.capacity)
		return cache.NewMemCache(buf)
	default:
		panic("xcache: unknown cache replacement policy")
	}
}

func (b *Builder) LRU() *Builder {
	b.policy = LRU
	return b
}

func (b *Builder) Capacity(capacity int) *Builder {
	b.capacity = capacity
	return b
}

func NewBuilder() *Builder {
	return &Builder{}
}
