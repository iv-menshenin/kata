package main

type (
	entity struct {
		someData string
	}
	entities       []entity
	sortedEntities struct {
		entities entities
	}
)

func (c *sortedEntities) grow() {
	var (
		oldCap = cap(c.entities)
		newCap = 10
	)
	if oldCap < 100 {
		newCap = oldCap * 2
	} else if oldCap < 1000 {
		newCap = oldCap + (oldCap / 2)
	} else {
		newCap = newCap + 500
	}
	var entities = make([]entity, len(c.entities), newCap)
	copy(entities, c.entities)
	c.entities = entities
}

func (c *sortedEntities) Add(e entity) {
	if !(len(c.entities) < cap(c.entities)) {
		c.grow()
	}
	pos := c.entities.Mid(e)
	l := len(c.entities)
	if pos == l {
		c.entities = append(c.entities, e)
	} else {
		c.entities = append(c.entities, c.entities[l-1])
		copy(c.entities[pos+1:l], c.entities[pos:l-1])
		c.entities[pos] = e
	}
}

func (c entities) Mid(e entity) int {
	if len(c) == 0 {
		return 0
	}
	var (
		pos      = len(c) / 2
		down, up = 0, len(c) - 1
	)
	for {
		if pos == up && pos == down {
			if c[pos].someData < e.someData {
				return pos + 1
			}
			return pos
		}
		if c[pos].someData > e.someData && !(pos-1 < down) {
			up = pos - 1
		} else if c[pos].someData < e.someData && !(pos+1 > up) {
			down = pos + 1
		} else {
			return pos
		}
		pos = down + (up-down)/2
	}
}

func (c *sortedEntities) Find(someData string) int {
	if len(c.entities) == 0 {
		return -1
	}
	var (
		pos      = len(c.entities) / 2
		down, up = 0, len(c.entities) - 1
	)
	for {
		if c.entities[pos].someData == someData {
			return pos
		}
		if pos == up && pos == down {
			return -1
		}
		if c.entities[pos].someData > someData && !(pos-1 < down) {
			up = pos - 1
		} else if c.entities[pos].someData < someData && !(pos+1 > up) {
			down = pos + 1
		}
		pos = down + (up-down)/2
	}
}
