package processor

import (
	"fmt"
	"regexp"
	"sorted_map_task/ordered/vanil"
	"sorted_map_task/textprocessor"
)

type TextProcessor struct {
	OrderedMap  *vanil.OrderedMap
	IgnoreToken map[string]int
	TokenRe     *regexp.Regexp
}

func (p *TextProcessor) ProcessText(text string) {
	textprocessor.ProcessText(text, p.OrderedMap, p.TokenRe, p.IgnoreToken)
}

func (p *TextProcessor) Top(count int) string {
	res := textprocessor.TopWordsByUsagePairList(p.OrderedMap, count, p.IgnoreToken)
	str := fmt.Sprint(res)
	return str
}
