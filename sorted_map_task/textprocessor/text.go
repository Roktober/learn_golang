package textprocessor

import (
	"regexp"
	"sort"
	"sorted_map_task/ordered"
	"sorted_map_task/ordered/vanil"
	"strings"
)

func Tokenize(text string, re *regexp.Regexp) []string {
	text = re.ReplaceAllString(text, "")
	tokens := strings.Split(text, " ")

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

func FilterTokens(tokens []string, minSymbolCount int, ignore map[string]int) []string {
	if len(tokens) >= 0 && len(tokens) <= 2 {
		return []string{""}
	}

	ignore[tokens[0]] = 0
	ignore[tokens[len(tokens)-1]] = 0

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

func ProcessText(text string, orderedMap ordered.PairContainer, re *regexp.Regexp, ignore map[string]int) {
	tokens := FilterTokens(Tokenize(text, re), 4, ignore)
	for _, token := range tokens {
		if token != "" {
			present := orderedMap.KeyExist(token)
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

func TopWordsByUsagePairList(orderedMap *vanil.OrderedMap, count int, ignore map[string]int) ordered.PairList {
	sorted := make(ordered.PairList, orderedMap.Len())
	for i, v := range orderedMap.OrderedItems {
		_, ok := ignore[v]
		if !ok {
			sorted[i] = ordered.MapItemStringInt{Key: v, Value: orderedMap.Get(v)}
		}
	}
	sort.Sort(sorted)
	return sorted[orderedMap.Len()-count:]
}
