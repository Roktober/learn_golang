package slice_utils

func DeleteByIndex(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func DeleteOnAppendWithCopy(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func DeleteOnAppend(slice []int, index int) []int {
	return slice[:index+copy(slice[index:], slice[index+1:])]
}

func InsertByIndex(slice []int, el int, index int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = el
	return slice
}

func InsertOnAppend(slice []int, el int, index int) []int {
	return append(slice[:index], append([]int{el}, slice[index:]...)...)
}
