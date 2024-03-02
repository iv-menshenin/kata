package fusioncol

type (
	Collection[T any] struct {
		count int
		last  *bucket[T]
		cache []bucket[T]
	}
	bucket[T any] struct {
		count int
		cont  []T
		prev  *bucket[T]
	}
)

func (c *Collection[T]) Get(i int) *T {
	if i > c.count-1 {
		panic(ErrOutOfBounds{i: i, l: c.count})
	}
	var cur = c.last
	var x = c.count
	for cur != nil && x > i+cur.count {
		x -= cur.count
		cur = cur.prev
	}
	return &cur.cont[i-(x-cur.count)]
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
		c.last = c.newBucket()
		return
	}
	n := c.newBucket()
	n.prev = c.last
	c.last = n
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
	c.last = c.last.prev
}
