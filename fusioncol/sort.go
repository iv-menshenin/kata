package fusioncol

type Cmp[T any] struct {
	collection *Collection[T]
	comparator func(*T, *T) bool
}

func Sortable[T any](collection *Collection[T], less func(*T, *T) bool) Cmp[T] {
	return Cmp[T]{
		collection: collection,
		comparator: less,
	}
}

func (c Cmp[T]) Len() int {
	return c.collection.Len()
}

func (c Cmp[T]) Less(i, j int) bool {
	return c.comparator(c.collection.Get(i), c.collection.Get(j))
}

func (c Cmp[T]) Swap(i, j int) {
	a, b := c.collection.Get(i), c.collection.Get(j)
	*a, *b = *b, *a
}
