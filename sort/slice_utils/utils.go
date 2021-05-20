package slice_utils

func DeleteByIndex(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func InsertByIndex(slice []int, el int, index int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = el
	return slice
}
