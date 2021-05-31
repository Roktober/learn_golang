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
	if !m.KeyPresent(key) {
		m.appendToOrderedSlice(key)
		m.expandOrderedSlice()
	}
	m.internalMap[key] = value
}

func (m *OrderedMap) Get(key string) int {
	return m.internalMap[key]
}

func (m *OrderedMap) KeyPresent(key string) bool {
	_, keyPresent := m.internalMap[key]
	return keyPresent
}

func (m *OrderedMap) appendToOrderedSlice(key string) {
	m.OrderedItems = append(m.OrderedItems, key)
}

func (m *OrderedMap) expandOrderedSlice() {
	if float32(len(m.OrderedItems))/float32(cap(m.OrderedItems)) <= 0.75 {
		return
	}
	newSlice := make([]string, len(m.OrderedItems), len(m.OrderedItems)*3/2)
	copy(newSlice, m.OrderedItems)

	m.OrderedItems = newSlice
}

func NewOrderedMap(cap int) *OrderedMap {
	return &OrderedMap{make([]string, 0, cap), make(map[string]int, cap)}
}
