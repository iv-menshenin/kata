package fusioncol

import "context"

func (c *Collection[T]) Iterator(ctx context.Context, backward bool, buf int) <-chan *T {
	ch := make(chan *T, buf)
	if backward {
		go c.sendBackward(ctx, ch)
	} else {
		go c.sendForward(ctx, ch)
	}
	return ch
}

func (c *Collection[T]) sendBackward(ctx context.Context, ch chan<- *T) {
	defer close(ch)
	var (
		cur = c.last
		idx = cur.count
	)
	for {
		if idx--; idx < 0 {
			cur = cur.prev
			if cur == nil {
				break
			}
			idx = cur.count - 1
		}
		select {
		case <-ctx.Done():
			return
		case ch <- &cur.cont[idx]:
			// next
		}
	}
}

func (c *Collection[T]) sendForward(ctx context.Context, ch chan<- *T) {
	defer close(ch)
	var q = make([]*bucket[T], 0)
	for cur := c.last; cur != nil; cur = cur.prev {
		q = append(q, cur)
	}
	if len(q) == 0 {
		return
	}
	var idx int
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- &q[len(q)-1].cont[idx]:
			idx++
			if idx >= q[len(q)-1].count {
				if q = q[:len(q)-1]; len(q) == 0 {
					return
				}
				idx = 0
			}
		}
	}
}
