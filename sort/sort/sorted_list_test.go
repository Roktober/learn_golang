package sort_test

import (
	"fmt"
	"premetiv_sort/sort"
	"testing"
)

type sortTest struct {
	input    []int
	expected []int
}

var sortTests = []sortTest{
	{[]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}, []int{-2000, -100, -5, -4, 1, 2, 3, 4, 5, 10}},
}

func TestQuickSortPointer(t *testing.T) {
	for _, tt := range sortTests {
		actual := make([]int, len(tt.input))
		copy(actual, tt.input)
		sort.QuickSortPointer(&actual)
		if !testSliceEq(actual, tt.expected) {
			t.Errorf("Actual != Expected \n %v != %v", actual, tt.expected)
		}
	}
}

func TestNewSortedListInt_Append(t *testing.T) {
	for _, tt := range sortTests {
		actual := make([]int, 0)
		sortedList := sort.NewSortedListInt(&actual)
		for _, el := range tt.input {
			sortedList.Append(el)
		}
		if !testSliceEq(actual, tt.expected) {
			t.Errorf("Actual != Expected \n %v != %v", actual, tt.expected)
		}
	}
}
func testSliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			fmt.Println(a[i])
			fmt.Println(b[i])
			return false
		}
	}

	return true
}

func BenchmarkQuickSortPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := inputArr[:]
		b.StartTimer()
		sort.QuickSortPointer(&arrSlice)
	}
}

func BenchmarkNewSortedListInt_Append(b *testing.B) {
	b.StopTimer()
	actual := make([]int, 0)
	sortedList := sort.NewSortedListInt(&actual)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, el := range [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000} {
			sortedList.Append(el)
		}
	}
}

func BenchmarkNewSortedListInt_Delete(b *testing.B) {
	b.StopTimer()
	inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
	arrSlice := inputArr[:]
	sortedList := sort.NewSortedListInt(&arrSlice)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sortedList.Delete(5)
	}
	_ = sortedList
}
