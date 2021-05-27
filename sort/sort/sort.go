package sort

import (
	"math/rand"
	"premetiv_sort/slice_utils"
)

func OnFly(sourceArray []int, number int) []int {
	if len(sourceArray) == 0 {
		return append(sourceArray, number)
	}
	prevIndex := 0
	for index, el := range sourceArray {
		if el >= number {
			return slice_utils.InsertByIndex(sourceArray, number, prevIndex)
		}
		prevIndex = index
	}

	return append(sourceArray, number)
}
func QuickSort(sourceArray []int) []int {
	if len(sourceArray) < 2 {
		return sourceArray
	}

	left, right := 0, len(sourceArray)-1

	pivotIndex := rand.Int() % len(sourceArray)

	sourceArray[pivotIndex], sourceArray[right] = sourceArray[right], sourceArray[pivotIndex]

	for i := 0; i < len(sourceArray); i++ {
		if sourceArray[i] < sourceArray[right] {
			sourceArray[i], sourceArray[left] = sourceArray[left], sourceArray[i]
			left++
		}
	}

	sourceArray[left], sourceArray[right] = sourceArray[right], sourceArray[left]

	QuickSort(sourceArray[:left])
	QuickSort(sourceArray[left+1:])

	return sourceArray
}
