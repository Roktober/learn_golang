package textprocessor_test

import (
	"log"
	"os"
	"sorted_map_task/ordered"
	"sorted_map_task/reader"
	"sorted_map_task/textprocessor"
	"testing"
)

func BenchmarkProcessText_small_text(b *testing.B) {
	for n := 0; n < b.N; n++ {
		orderedMap := ordered.New(10)
		textprocessor.ProcessText("dawdawd dawdaw 1-id dwdawd dawdawdwa", orderedMap)
	}
}

func BenchmarkProcessText_large_text(b *testing.B) {
	file, err := os.Open("../file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		orderedMap := ordered.New(100)
		scanner := reader.CreateBufferedScanner(file, 1000000, reader.ScanLinesByDotNewLine)
		b.StartTimer()
		for scanner.Scan() {
			text := scanner.Text()
			textprocessor.ProcessText(text, orderedMap)
		}
	}
}

func BenchmarkTopWordsByUsage_100(b *testing.B) {
	orderedMap := ordered.New(10)
	textprocessor.ProcessText("aaa, bbbb, cccccccc, kkkkk, ddddddd, dddddd, ddadad", orderedMap)
	for n := 0; n < b.N; n++ {
		textprocessor.TopWordsByUsage(orderedMap, 3)
	}
}

func BenchmarkTopWordsByUsage_200(b *testing.B) {
	file, err := os.Open("../file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	orderedMap := ordered.New(100)
	scanner := reader.CreateBufferedScanner(file, 1000000, reader.ScanLinesByDotNewLine)
	for scanner.Scan() {
		text := scanner.Text()
		textprocessor.ProcessText(text, orderedMap)
	}

	for n := 0; n < b.N; n++ {
		textprocessor.TopWordsByUsage(orderedMap, 200)
	}
}

/*
goos: windows
goarch: amd64
pkg: sorted_map_task/textprocessor
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkProcessText_small_text-12        346326              3334 ns/op            1554 B/op         31 allocs/op
BenchmarkProcessText_large_text-12        104787             10212 ns/op              50 B/op          0 allocs/op
PASS
ok      sorted_map_task/textprocessor   32.183s
*/
