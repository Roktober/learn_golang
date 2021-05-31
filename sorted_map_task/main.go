package main

import (
	"fmt"
	"log"
	"os"
	"sorted_map_task/ordered/vanil"
	"sorted_map_task/reader"
	"sorted_map_task/textprocessor"
)

func main() {
	megabyte := 1000000 // bytes

	orderedMap := vanil.NewOrderedMap(1000)

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

	for _, el := range textprocessor.TopWordsByUsagePairList(orderedMap, 10) {
		fmt.Println(el.Key, el.Value)
	}
	//textprocessor.TopWordsByUsage(orderedMap, 10)
	//[{this 183} {were 190} {that 194} {soft 182} {from 148} {Scarlett 165} {side 120} {hair 181} {good 112} {said 225}]
	//going 51
	//about 59
	//went 62
	//Melanie 63
	//from 68
	//were 88
	//that 126
	//with 128
	//Scarlett 146
	//said 200
}
