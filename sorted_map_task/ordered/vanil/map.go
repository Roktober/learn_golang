package vanil

type OrderedMap struct {
	OrderedItems []string
	internalMap  map[string]int
}

func (m *OrderedMap) Cap() int {
	return cap(m.OrderedItems)
}

func (m *OrderedMap) Len() int {
	return len(m.OrderedItems)
}

func (m *OrderedMap) Put(key string, value int) {
	if !m.KeyExist(key) {
		m.appendToOrderedSlice(key)
	}
	m.internalMap[key] = value
}

func (m *OrderedMap) Get(key string) int {
	return m.internalMap[key]
}

func (m *OrderedMap) KeyExist(key string) bool {
	_, ok := m.internalMap[key]
	return ok
}

func (m *OrderedMap) appendToOrderedSlice(key string) {
	m.OrderedItems = append(m.OrderedItems, key)
}

func NewOrderedMap(cap int) *OrderedMap {
	return &OrderedMap{make([]string, 0, cap), make(map[string]int, cap)}
}
