package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sorted_map_task/ordered/vanil"
	"sorted_map_task/reader"
	"sorted_map_task/textprocessor"
)

func main() {
	megabyte := 1000000 // bytes

	re, err := regexp.Compile(`[^\w ]`)
	if err != nil {
		panic(err)
	}

	orderedMap := vanil.NewOrderedMap(1000)
	ingoreToken := make(map[string]int)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}(file)

	scanner := reader.CreateBufferedScanner(file, megabyte, reader.ScanLinesByDotNewLine)

	for scanner.Scan() {
		text := scanner.Text()
		textprocessor.ProcessText(text, orderedMap, re, ingoreToken)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, el := range textprocessor.TopWordsByUsagePairList(orderedMap, 10, ingoreToken) {
		fmt.Println(el.Key, el.Value)
	}
	fmt.Println(ingoreToken)
	//textprocessor.TopWordsByUsage(orderedMap, 10)
	//[{this 183} {were 190} {that 194} {soft 182} {from 148} {Scarlett 165} {side 120} {hair 181} {good 112} {said 225}]
	//dont 34
	//like 35
	//looked 39
	//must 40
	//they 42
	//will 42
	//when 42
	//went 62
	//from 68
	//with 128
}
