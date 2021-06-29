package ordered

type listMapItem struct {
	value *MapItemStringInt
	next  *listMapItem
}

func (l *listMapItem) Value() *MapItemStringInt {
	return l.value
}

func (l *listMapItem) Next() *listMapItem {
	return l.next
}

type linkedList struct {
	size int
	head *listMapItem
	tail *listMapItem
}

func (l *linkedList) Size() int {
	return l.size
}

func (l *linkedList) Put(value *MapItemStringInt) {
	newItem := &listMapItem{value: value}
	if l.head == nil {
		l.head = newItem
		l.tail = newItem
		return
	}
	l.tail.next = newItem
	l.tail = newItem
}

func (l *linkedList) Head() *listMapItem {
	return l.head
}

func (l *linkedList) Tail() *listMapItem {
	return l.tail
}
