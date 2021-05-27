package array_protector

import (
	"premetiv_sort/slice_utils"
	"premetiv_sort/sort"
)

func ProtectArray(array []int, number int) []int {
	var cleared bool
	array, cleared = clearSlice(array, number)
	if !cleared {
		array = sort.OnFly(array, number)
	}
	return array
}

func ProtectArrayQuickSort(array []int, number int) []int {
	var cleared bool
	array, cleared = clearSlice(array, number)
	if !cleared {
		array = append(array, number)
		array = sort.QuickSort(array)
	}
	return array
}

func clearSlice(slice []int, number int) ([]int, bool) {
	cleared := false
	antiNumber := number * -1
	for i := 0; i < len(slice); i++ {
		if slice[i] == antiNumber {
			slice = slice_utils.DeleteByIndex(slice, i)
			cleared = true
			break
		}
	}
	return slice, cleared
}
