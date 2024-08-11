package clipboard

import (
	"log"
	"strings"

	"github.com/atotto/clipboard"
)

// GetTextFromClipboard extracts and returns non-empty lines from the system clipboard.
func GetTextFromClipboard() []string {
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read clipboard: %v", err)
		panic(err)
	}
	var words []string = strings.Split(text, "\n")
	var nonEmptyWords []string

	for _, word := range words {
		trimmedWord := strings.TrimSpace(word)
		if trimmedWord != "" {
			nonEmptyWords = append(nonEmptyWords, trimmedWord)
		}
	}

	return nonEmptyWords
}
