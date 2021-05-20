package array_protector

import (
	"fmt"
	"testing"
)

type protectArrayTest struct {
	sourceArray []int
	input       []int
	expected    []int
}

var protectArrayTests = []protectArrayTest{
	{make([]int, 0), []int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}, []int{-2000, -100, 1, 2, 3, 10}},
}

func TestProtectArray(t *testing.T) {
	for _, tt := range protectArrayTests {
		actual := make([]int, len(tt.sourceArray))
		copy(actual, tt.sourceArray)
		for _, num := range tt.input {
			actual = ProtectArray(actual, num)
		}
		if !testSliceEq(actual, tt.expected) {
			t.Errorf("Actual != Expected \n %v != %v", actual, tt.expected)
		}
	}
}

func TestProtectArrayQuickSort(t *testing.T) {
	for _, tt := range protectArrayTests {
		actual := make([]int, len(tt.sourceArray))
		copy(actual, tt.sourceArray)
		for _, num := range tt.input {
			actual = ProtectArrayQuickSort(actual, num)
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

func BenchmarkProtectArrayQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := make([]int, 0)
		for i := 0; i < len(inputArr); i++ {
			arrSlice = ProtectArrayQuickSort(arrSlice, inputArr[i])
		}
	}
}

func BenchmarkProtectArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := make([]int, 0)
		for i := 0; i < len(inputArr); i++ {
			arrSlice = ProtectArray(arrSlice, inputArr[i])
		}
	}
}

/*
goos: windows
goarch: amd64
pkg: premetiv_sort/array_protector
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkProtectArrayQuickSort-12        1458886               823.3 ns/op
BenchmarkProtectArray-12                 5747326               210.1 ns/op
PASS
ok      premetiv_sort/array_protector   3.478s

*/
