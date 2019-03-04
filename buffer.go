package xcache

type Buffer interface {
	IsFull() bool
	Evict() *Block
	Add(key Key, value Value) *Block
	Remove(block *Block)
	Access(block *Block) *Block
	Size() int
	Keys() []Key
	Values() []Value
	Has(block *Block) bool
}

type LRUBuffer struct {
	head     *Block
	tail     *Block
	size     int
	capacity int
}

func (l LRUBuffer) Keys() []Key {
	keys := make([]Key, 0)

	for curr := l.head.next; curr != nil; curr = curr.next {
		keys = append(keys, curr.key)
	}

	return keys
}

func (l LRUBuffer) Values() []Value {
	values := make([]Value, 0)

	for curr := l.head.next; curr != nil; curr = curr.next {
		values = append(values, curr.val)
	}

	return values
}

func (l LRUBuffer) Has(block *Block) bool {
	for curr := l.head; curr != nil; curr = curr.next {
		if curr == block {
			return true
		}
	}

	return false
}

func (l *LRUBuffer) Access(block *Block) *Block {
	l.Remove(block)

	key := block.key
	value := block.val

	return l.Add(key, value)
}

func (l LRUBuffer) Size() int {
	return l.size
}

func (l *LRUBuffer) Add(key Key, value Value) *Block {
	newBlock := Block{
		prev: l.tail,
		key:  key,
		val:  value,
	}
	l.tail.next = &newBlock

	l.tail = &newBlock
	l.size++

	return &newBlock
}

func (l *LRUBuffer) Remove(block *Block) {
	if block == nil {
		return
	}

	prev := block.prev
	prev.next = block.next

	if prev.next != nil {
		prev.next.prev = prev
	}

	if l.tail == block {
		l.tail = block.prev
	}

	l.size--
}

func (l LRUBuffer) IsFull() bool {
	return l.size == l.capacity
}

func (l *LRUBuffer) Evict() *Block {
	block := l.head.next
	l.Remove(l.head.next)
	return block
}

type Block struct {
	prev *Block
	next *Block
	key  Key
	val  Value
}

func NewLRUBuffer(capacity int) Buffer {
	dummyBlock := Block{}

	return &LRUBuffer{
		head:     &dummyBlock,
		tail:     &dummyBlock,
		size:     0,
		capacity: capacity,
	}
}
