package orderedMap

// You need to implement a hash-map structure
// so that its iterator returns keys in the order of their addition.

type (
	keyExistence struct {
		keyName string
		removed bool
	}
	OrderedMap struct {
		data     map[string]string
		ordering []keyExistence
		keyPosit map[string]int
	}
)

func (m *OrderedMap) Set(key, value string) {
	if _, ok := m.data[key]; !ok {
		m.keyPosit[key] = len(m.ordering)
		m.ordering = append(m.ordering, keyExistence{keyName: key})
	}
	m.data[key] = value
}

func (m *OrderedMap) Get(key string) (string, bool) {
	value, ok := m.data[key]
	return value, ok
}

func (m *OrderedMap) Delete(key string) {
	pos, ok := m.keyPosit[key]
	if !ok {
		return
	}
	m.ordering[pos].removed = true
	delete(m.data, key)
	delete(m.keyPosit, key)
	if m.isNeedToReclaimSpace() {
		m.reclaimSpace()
	}
}

func (m *OrderedMap) isNeedToReclaimSpace() bool {
	deleted := float64(len(m.ordering) - len(m.data))
	load := float64(len(m.ordering))
	return deleted/load > 0.25
}

func (m *OrderedMap) reclaimSpace() {
	var newOrdered = make([]keyExistence, 0, len(m.data))
	for _, key := range m.ordering {
		if !key.removed {
			newOrdered = append(newOrdered, key)
		}
	}
	var newPosit = make(map[string]int, len(newOrdered))
	for i, key := range newOrdered {
		newPosit[key.keyName] = i
	}
	m.ordering = newOrdered
	m.keyPosit = newPosit
}

func (m *OrderedMap) Range() <-chan string {
	var ch = make(chan string)
	go func() {
		defer close(ch)
		for _, key := range m.ordering {
			if !key.removed {
				ch <- key.keyName
			}
		}
	}()
	return ch
}

func New() *OrderedMap {
	return &OrderedMap{
		data:     make(map[string]string),
		keyPosit: make(map[string]int),
	}
}
