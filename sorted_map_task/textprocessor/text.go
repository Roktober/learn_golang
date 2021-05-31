package textprocessor

import (
	"regexp"
	"sort"
	"sorted_map_task/ordered"
	"strings"
)

func Tokenize(text string) []string {
	re, err := regexp.Compile(`[^\w' ]`)
	if err != nil {
		panic(err)
	}
	text = re.ReplaceAllString(text, "")
	tokens := strings.Split(text, " ")

	result := make([]string, len(tokens))
	i := 0
	for _, token := range tokens {
		token = strings.Trim(token, "'")
		if token != "" {
			result[i] = token
			i++
		}
	}
	return result[:i]
}

func FilterTokens(tokens []string, minSymbolCount int) []string {
	if len(tokens) >= 0 && len(tokens) <= 2 {
		return []string{""}
	}
	tokens = tokens[1 : len(tokens)-1] // Пропуск первого и последнего токена в предложении
	for i, token := range tokens {
		if len(token) < minSymbolCount {
			tokens[i] = ""
		}
	}
	return FilterEmptyToken(tokens)
}

func FilterEmptyToken(tokens []string) []string {
	result := make([]string, len(tokens))

	i := 0
	for _, token := range tokens {
		if token != "" {
			result[i] = token
			i++
		}
	}

	return result[:i]
}

func ProcessText(text string, orderedMap *ordered.MapStringInt) {
	tokens := FilterTokens(Tokenize(text), 4)
	for _, token := range tokens {
		if token != "" {
			present, _ := orderedMap.KeyPresent(token)
			if present {
				orderedMap.Put(token, orderedMap.Get(token)+1)
			} else {
				orderedMap.Put(token, 1)
			}
		}
	}
}

func TopWordsByUsage(orderedMap *ordered.MapStringInt, count int) []ordered.MapItemStringInt {
	values := make([]int, orderedMap.Size)
	topValues := make([]ordered.MapItemStringInt, count)
	i := 0
	for el := orderedMap.OrderedValues.Head(); el != nil; el = el.Next() {
		values[i] = el.Value().Value
		i++
	}
	sort.Ints(values)

	values = values[orderedMap.Size-count:]
	holded := 0
	for el := orderedMap.OrderedValues.Head(); el != nil; el = el.Next() {
		val := el.Value().Value
		for i, max := range values {
			if val == max {
				values[i] = 0
				topValues[holded] = *el.Value()
				holded += 1
				break
			}
		}
		if holded == count {
			break
		}
	}
	return topValues
}
