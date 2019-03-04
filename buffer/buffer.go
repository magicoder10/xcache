package buffer

import "xcache/entity"

type Buffer interface {
	IsFull() bool
	Evict() *Page
	Add(key entity.Key, value entity.Value) *Page
	Remove(Page *Page)
	Access(Page *Page) *Page
	Size() int
	Keys() []entity.Key
	Values() []entity.Value
	Has(Page *Page) bool
}

type LRU struct {
	head     *Page
	tail     *Page
	size     int
	capacity int
}

func (l LRU) Keys() []entity.Key {
	keys := make([]entity.Key, 0)

	for curr := l.head.next; curr != nil; curr = curr.next {
		keys = append(keys, curr.Key)
	}

	return keys
}

func (l LRU) Values() []entity.Value {
	values := make([]entity.Value, 0)

	for curr := l.head.next; curr != nil; curr = curr.next {
		values = append(values, curr.Val)
	}

	return values
}

func (l LRU) Has(Page *Page) bool {
	for curr := l.head; curr != nil; curr = curr.next {
		if curr == Page {
			return true
		}
	}

	return false
}

func (l *LRU) Access(Page *Page) *Page {
	l.Remove(Page)

	key := Page.Key
	value := Page.Val

	return l.Add(key, value)
}

func (l LRU) Size() int {
	return l.size
}

func (l *LRU) Add(key entity.Key, value entity.Value) *Page {
	newPage := Page{
		prev: l.tail,
		Key:  key,
		Val:  value,
	}
	l.tail.next = &newPage

	l.tail = &newPage
	l.size++

	return &newPage
}

func (l *LRU) Remove(Page *Page) {
	if Page == nil {
		return
	}

	prev := Page.prev
	prev.next = Page.next

	if prev.next != nil {
		prev.next.prev = prev
	}

	if l.tail == Page {
		l.tail = Page.prev
	}

	l.size--
}

func (l LRU) IsFull() bool {
	return l.size == l.capacity
}

func (l *LRU) Evict() *Page {
	Page := l.head.next
	l.Remove(l.head.next)
	return Page
}

type Page struct {
	prev *Page
	next *Page
	Key  entity.Key
	Val  entity.Value
}

func NewLRU(capacity int) Buffer {
	dummyPage := Page{}

	return &LRU{
		head:     &dummyPage,
		tail:     &dummyPage,
		size:     0,
		capacity: capacity,
	}
}
