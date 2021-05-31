package textprocessor_test

import (
	"log"
	"os"
	"sorted_map_task/ordered"
	"sorted_map_task/ordered/vanil"
	"sorted_map_task/reader"
	"sorted_map_task/textprocessor"
	"testing"
)

func GenericBenchmarkProcessText_small_text(b *testing.B, create func() ordered.PairContainer) {
	for n := 0; n < b.N; n++ {
		orderedMap := create()
		textprocessor.ProcessText("dawdawd dawdaw 1-id dwdawd dawdawdwa", orderedMap)
	}
}

func GenericBenchmarkProcessText_large_text(b *testing.B, create func() ordered.PairContainer) {
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
		orderedMap := create()
		scanner := reader.CreateBufferedScanner(file, 1000000, reader.ScanLinesByDotNewLine)
		b.StartTimer()
		for scanner.Scan() {
			text := scanner.Text()
			textprocessor.ProcessText(text, orderedMap)
		}
	}
}

func BenchmarkTopWordsByUsage_100(b *testing.B) {
	orderedMap := ordered.NewOrderedMap(10)
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

	orderedMap := ordered.NewOrderedMap(100)
	scanner := reader.CreateBufferedScanner(file, 1000000, reader.ScanLinesByDotNewLine)
	for scanner.Scan() {
		text := scanner.Text()
		textprocessor.ProcessText(text, orderedMap)
	}

	for n := 0; n < b.N; n++ {
		textprocessor.TopWordsByUsage(orderedMap, 200)
	}
}

func BenchmarkProcessTextMapImpl_small_text(b *testing.B) {
	GenericBenchmarkProcessText_small_text(b, func() ordered.PairContainer {
		return ordered.NewOrderedMap(100)
	})
}

func BenchmarkProcessTextMapImpl_large_text(b *testing.B) {
	GenericBenchmarkProcessText_large_text(b, func() ordered.PairContainer {
		return ordered.NewOrderedMap(100)
	})
}

func BenchmarkProcessTextVanil_small_text(b *testing.B) {
	GenericBenchmarkProcessText_small_text(b, func() ordered.PairContainer {
		return vanil.NewOrderedMap(100)
	})
}

func BenchmarkProcessTextVanil_large_text(b *testing.B) {
	GenericBenchmarkProcessText_large_text(b, func() ordered.PairContainer {
		return vanil.NewOrderedMap(100)
	})
}

/*
goos: windows
goarch: amd64
pkg: sorted_map_task/textprocessor
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkTopWordsByUsage_100-12                  8082734               146.3 ns/op           136 B/op          3 allocs/op
BenchmarkTopWordsByUsage_200-12                    20395             58080 ns/op            9227 B/op          6 allocs/op
BenchmarkProcessTextMapImpl_small_text-12         337890              3429 ns/op            2377 B/op         30 allocs/op
BenchmarkProcessTextMapImpl_large_text-12         113811             10210 ns/op              46 B/op          0 allocs/op
BenchmarkProcessTextVanil_small_text-12           260190              4556 ns/op            7414 B/op         29 allocs/op
BenchmarkProcessTextVanil_large_text-12           103288             10026 ns/op              49 B/op          0 allocs/op
ok      sorted_map_task/textprocessor   20.478s
*/
