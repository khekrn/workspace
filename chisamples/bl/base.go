package bl

import (
	"strings"
)

// TokenizeText converts each sentences to tokens
func TokenizeText(text string) []string {
	if len(text) == 0 {
		return nil
	}
	tokens := strings.Split(text, " ")
	return tokens
}

// WordCount - generates word count for given text
func WordCount(text string) map[string]int {
	if len(text) == 0 {
		return nil
	}

	wc := map[string]int{}
	tokens := TokenizeText(text)
	for _, item := range tokens {
		if _, exist := wc[item]; exist {
			wc[item] = wc[item] + 1
		} else {
			wc[item] = 1
		}
	}

	return wc
}
