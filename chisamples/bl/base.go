package bl

import (
	"strings"

	"github.com/rs/zerolog/log"
)

// TokenizeText converts each sentences to tokens
func TokenizeText(text string) []string {
	log.Info().Msg("in TokenizeText")

	if len(text) == 0 {
		log.Error().Msg("in TokenizeText - empty text found")
		return nil
	}
	tokens := strings.Split(text, " ")

	log.Print("return from TokenizeText ", tokens)
	return tokens
}

// WordCount - generates word count for given text
func WordCount(text string) map[string]int {
	log.Info().Msg("in WordCount")

	if len(text) == 0 {
		log.Error().Msg("in WordCount - empty text found")
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
