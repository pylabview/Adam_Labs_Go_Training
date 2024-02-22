package tokenizer

import (
	"strings"
)

var (
	suffixes = []string{"ed", "ing", "s"}
)

// working -> work, works -> word, worked -> work
func Stem(word string) string {
	for _, s := range suffixes {
		if strings.HasSuffix(word, s) {
			return word[:len(word)-len(s)]
		}
	}

	return word
}

// "Who's on first?" -> []string{"Who", "s", "on", "first"}
func initialSplit(text string) []string {
	// Most words are 4 letters
	size := len(text) / 4
	fs := make([]string, 0, size)

	i := 0
	for i < len(text) {
		// eat start
		for i < len(text) && !isLetter(text[i]) {
			i++
		}

		if i == len(text) {
			break
		}

		j := i + 1
		for j < len(text) && isLetter(text[j]) {
			j++
		}

		fs = append(fs, text[i:j])
		i = j
	}

	return fs
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func Tokenize(text string) []string {
	// var tokens []string
	fs := initialSplit(text)
	tokens := make([]string, 0, len(fs))
	for _, tok := range fs {
		tok = strings.ToLower(tok)
		tok = Stem(tok)
		if tok != "" {
			tokens = append(tokens, tok)
		}
	}

	return tokens
}
