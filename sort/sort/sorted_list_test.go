package sort_test

import (
	"fmt"
	"math/rand"
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

// Testing utils

func generateData(b *testing.B, size int) []int {
	var randomizer = rand.New(rand.NewSource(42))
	var data = make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = randomizer.Int()
	}
	return data
}

func BenchmarkSortedListInt_GetMax_100(b *testing.B) {
	b.StopTimer()
	inputArr := generateData(b, 100)
	sortedList := sort.NewSortedListInt(&inputArr)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sortedList.GetMax()
	}
	_ = sortedList
}

func BenchmarkSortedListInt_GetMax_1000000(b *testing.B) {
	b.StopTimer()
	inputArr := generateData(b, 1000000)
	sortedList := sort.NewSortedListInt(&inputArr)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sortedList.GetMax()
	}
	_ = sortedList
}

func BenchmarkSortedListInt_GetMin_100(b *testing.B) {
	b.StopTimer()
	inputArr := generateData(b, 100)
	sortedList := sort.NewSortedListInt(&inputArr)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sortedList.GetMin()
	}
	_ = sortedList
}

func BenchmarkSortedListInt_GetMin_1000000(b *testing.B) {
	b.StopTimer()
	inputArr := generateData(b, 1000000)
	sortedList := sort.NewSortedListInt(&inputArr)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sortedList.GetMin()
	}
	_ = sortedList
}

/*
BenchmarkSortedListInt_GetMax_100-12            1000000000               0.4805 ns/op          0 B/op          0 allocs/op
BenchmarkSortedListInt_GetMax_1000000-12        1000000000               0.4846 ns/op          0 B/op          0 allocs/op
BenchmarkSortedListInt_GetMin_100-12            1000000000               0.3650 ns/op          0 B/op          0 allocs/op
BenchmarkSortedListInt_GetMin_1000000-12        1000000000               0.3690 ns/op          0 B/op          0 allocs/op
*/
