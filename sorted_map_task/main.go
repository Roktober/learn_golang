package main

import (
	"log"
	"os"
	"sorted_map_task/ordered"
	"sorted_map_task/reader"
	"sorted_map_task/textprocessor"
)

func main() {
	megabyte := 1000000 // bytes

	orderedMap := ordered.New(1000)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := reader.CreateBufferedScanner(file, megabyte, reader.ScanLinesByDotNewLine)

	for scanner.Scan() {
		text := scanner.Text()
		textprocessor.ProcessText(text, orderedMap)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	textprocessor.TopWordsByUsage(orderedMap, 10)
	//[{this 183} {were 190} {that 194} {soft 182} {from 148} {Scarlett 165} {side 120} {hair 181} {good 112} {said 225}]
}
