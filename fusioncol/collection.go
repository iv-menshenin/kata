package fusioncol

type (
	Collection[T any] struct {
		count int
		first *bucket[T]
		last  *bucket[T]
		cache []bucket[T]
	}
	bucket[T any] struct {
		count int
		cont  []T
		next  *bucket[T]
	}
)

func (c *Collection[T]) Get(i int) *T {
	if i > c.count-1 {
		panic(ErrOutOfBounds{i: i, l: c.count})
	}
	var cur = c.first
	var x = i
	for cur != nil && x >= cur.count {
		x -= cur.count
		cur = cur.next
	}
	return &cur.cont[x]
}

func (c *Collection[T]) Pop() T {
	if c.count == 0 {
		panic(ErrOutOfBounds{l: c.count})
	}
	l := len(c.last.cont)
	for l == 0 {
		c.removeLast()
		l = len(c.last.cont)
	}
	v := c.last.cont[l-1]
	c.last.cont = c.last.cont[:l-1]
	c.count--
	c.last.count--
	return v
}

func (c *Collection[T]) Push(elem T) *T {
	return c.Append(elem)
}

func (c *Collection[T]) Append(elem T) *T {
	if !c.capable() {
		c.extend()
	}
	l := len(c.last.cont)
	c.last.cont = c.last.cont[:l+1]
	c.last.cont[l] = elem
	c.last.count++
	c.count++
	return &c.last.cont[l]
}

func (c *Collection[T]) Len() int {
	return c.count
}

func (c *Collection[T]) capable() bool {
	if c.last == nil {
		return false
	}
	return len(c.last.cont) < cap(c.last.cont)
}

func (c *Collection[T]) extend() {
	if c.last == nil {
		c.first = c.newBucket()
		c.last = c.first
		return
	}
	c.last.next = c.newBucket()
	c.last = c.last.next
}

const (
	bucketsCache  = 32
	firstBucketSz = 16
	maxBucketSz   = 32768
)

func (c *Collection[T]) newBucket() *bucket[T] {
	if len(c.cache) < cap(c.cache) {
		c.cache = c.cache[:len(c.cache)+1]
	} else {
		c.cache = make([]bucket[T], 1, bucketsCache)
	}
	var b = &c.cache[len(c.cache)-1]
	b.cont = make([]T, 0, c.sz())
	return b
}

func (c *Collection[T]) sz() int {
	if c.last == nil {
		return firstBucketSz
	}
	var sz = cap(c.last.cont) * 2
	if sz > 512 {
		sz -= sz / 4
	}
	if sz > maxBucketSz {
		sz = maxBucketSz
	}
	return sz
}

func (c *Collection[T]) removeLast() {
	if c.last == nil {
		return
	}
	if c.last.count > 0 {
		panic("remove nonempty bucket")
	}
	if c.first.next == nil {
		c.last = nil
		c.first = nil
		if c.count > 0 {
			panic("remove first bucket in nonempty collection")
		}
	}
	var cur = c.first
	for {
		if cur.next.next == nil {
			cur.next = nil
			c.last = cur
			break
		}
		cur = cur.next
	}
	c.last.next = nil
}
