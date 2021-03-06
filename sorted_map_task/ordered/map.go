package ordered

type PairContainer interface {
	Put(string, int)
	Get(string) int
	KeyExist(key string) bool
}

type PairList []MapItemStringInt

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

type MapItemStringInt struct {
	Key   string
	Value int
}

type MapStringInt struct {
	OrderedValues linkedList
	Size          int
	Cap           int
	Bucket        []*MapItemStringInt
}

func (i *MapStringInt) Put(key string, value int) {
	keyPresent := i.KeyExist(key)
	index := i.getIndexByValue(key)
	if keyPresent {
		i.Bucket[index].Value = value
	} else {
		item := &MapItemStringInt{Key: key, Value: value}
		i.Bucket[index] = item
		i.OrderedValues.Put(item)
		i.Size = i.Size + 1
	}
	i.reindex()
}

func (i *MapStringInt) Get(key string) int {
	index := calculateHashString(key) % i.Cap
	return i.Bucket[index].Value
}

func (i *MapStringInt) KeyExist(key string) bool {
	index := calculateHashString(key) % i.Cap
	el := i.Bucket[index]
	if el != nil {
		return true
	}
	return false
}

func (i *MapStringInt) getIndexByValue(key string) int {
	index := calculateHashString(key) % i.Cap
	el := i.Bucket[index]
	if el != nil {
		return index
	}
	return index
}

func calculateHashString(key string) int {
	var hash int
	for _, b := range key {
		hash += int(b)
	}

	return hash * len(key)
}

func (i *MapStringInt) reindex() {
	if float32(i.Size)/float32(i.Cap) <= 0.75 {
		return
	}

	newCap := i.Cap * 3 / 2
	newBucket := make([]*MapItemStringInt, newCap)
	for _, el := range i.Bucket {
		if el != nil {
			index := calculateHashString(el.Key) % newCap
			newBucket[index] = el
		}
	}
	i.Cap = newCap
	i.Bucket = newBucket
}

func NewOrderedMap(cap int) *MapStringInt {
	return &MapStringInt{*new(linkedList), 0, cap, make([]*MapItemStringInt, cap)}
}
