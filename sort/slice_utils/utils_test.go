package slice_utils_test

import "testing"
import "premetiv_sort/slice_utils"

func BenchmarkInsertByIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := inputArr[:]
		slice_utils.InsertByIndex(arrSlice, 1, 0)
		slice_utils.InsertByIndex(arrSlice, 1, 5)
		slice_utils.InsertByIndex(arrSlice, 1, 10)
	}
}

func BenchmarkInsertByAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := inputArr[:]
		slice_utils.InsertOnAppend(arrSlice, 1, 0)
		slice_utils.InsertOnAppend(arrSlice, 1, 5)
		slice_utils.InsertOnAppend(arrSlice, 1, 10)
	}
}

func BenchmarkDeleteOnAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := inputArr[:]
		slice_utils.DeleteOnAppend(arrSlice, 0)
		slice_utils.DeleteOnAppend(arrSlice, 5)
		slice_utils.DeleteOnAppend(arrSlice, 9)
	}
}

func BenchmarkDeleteOnAppendWithCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := inputArr[:]
		slice_utils.DeleteOnAppendWithCopy(arrSlice, 0)
		slice_utils.DeleteOnAppendWithCopy(arrSlice, 5)
		slice_utils.DeleteOnAppendWithCopy(arrSlice, 9)
	}
}

func BenchmarkDeleteByIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inputArr := [10]int{1, 2, 3, -4, 4, -5, 5, 10, -100, -2000}
		arrSlice := inputArr[:]
		slice_utils.DeleteByIndex(arrSlice, 0)
		slice_utils.DeleteByIndex(arrSlice, 5)
		slice_utils.DeleteByIndex(arrSlice, 9)
	}
}

/*
goos: linux
goarch: amd64
pkg: premetiv_sort/slice_utils
cpu: AMD Ryzen 5 2600 Six-Core Processor
BenchmarkInsertByIndex-12                3064696               395.1 ns/op
BenchmarkInsertByAppend-12               2153281               551.7 ns/op
BenchmarkDeleteOnAppend-12              67482402                17.68 ns/op
BenchmarkDeleteOnAppendWithCopy-12      67731129                17.75 ns/op
BenchmarkDeleteByIndex-12               67854204                17.74 ns/op
PASS
ok      premetiv_sort/slice_utils       9.734s
*/
