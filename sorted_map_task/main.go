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
	// [{Scarlett 315} {white 81} {gentleman 100} {about 99} {going 86} {Ashley 85} {Melanie 113} {Everyone 77} {moments 67} {shock 86}]
}
