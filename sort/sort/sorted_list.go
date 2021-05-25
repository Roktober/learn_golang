package sort

import (
	"math/rand"
)

type SortedListInt struct {
	Slice *[]int
}

func NewSortedListInt(slice *[]int) SortedListInt {
	sortedList := SortedListInt{slice}
	sortedList.sort()
	return sortedList
}

func (l *SortedListInt) GetMax() int {
	return (*l.Slice)[len(*(l.Slice))-1]
}

func (l *SortedListInt) GetMin() int {
	return (*l.Slice)[0]
}

func (l *SortedListInt) Append(el int) {
	if len(*l.Slice) == 0 {
		*l.Slice = append(*l.Slice, el)
		return
	}
	prevIndex := 0
	for index, num := range *l.Slice {
		if num >= el {
			l.insert(prevIndex, el)
			return
		}
		prevIndex = index
	}

	*l.Slice = append(*l.Slice, el)
}

func (l *SortedListInt) Delete(index int) {
	(*l.Slice)[index] = (*l.Slice)[len((*l.Slice))-1]
	(*l.Slice)[len((*l.Slice))-1] = 0
	*l.Slice = (*l.Slice)[:len(*l.Slice)-1]
}

func (l *SortedListInt) insert(index int, el int) {
	*l.Slice = append(*l.Slice, 0)
	copy((*l.Slice)[index+1:], (*l.Slice)[index:])
	(*l.Slice)[index] = el
}

func (l *SortedListInt) sort() {
	QuickSortPointer(l.Slice)
}

func QuickSortPointer(sourceArray *[]int) {
	if len(*sourceArray) < 2 {
		return
	}

	left, right := 0, len(*sourceArray)-1

	pivotIndex := rand.Int() % len(*sourceArray)

	(*sourceArray)[pivotIndex], (*sourceArray)[right] = (*sourceArray)[right], (*sourceArray)[pivotIndex]

	for i := 0; i < len(*sourceArray); i++ {
		if (*sourceArray)[i] < (*sourceArray)[right] {
			(*sourceArray)[i], (*sourceArray)[left] = (*sourceArray)[left], (*sourceArray)[i]
			left++
		}
	}

	(*sourceArray)[left], (*sourceArray)[right] = (*sourceArray)[right], (*sourceArray)[left]

	rightSlice := (*sourceArray)[:left]
	leftSlice := (*sourceArray)[left+1:]
	QuickSortPointer(&rightSlice)
	QuickSortPointer(&leftSlice)

	return
}
