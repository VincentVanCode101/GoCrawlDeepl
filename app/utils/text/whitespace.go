package text

import "unicode"

// ContainsWhitespace checks if the input string contains any whitespace characters.
func ContainsWhitespace(phrase string) bool {
	for _, char := range phrase {
		if unicode.IsSpace(char) {
			return true
		}
	}
	return false
}
