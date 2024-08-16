package text

import "fmt"

// FormatTranslation formats the translation output with optional word type information.
func FormatTranslation(phrase string, translations map[string]string) string {
	translationText := fmt.Sprintf("Input:\n%s\n\nMain translation:\n%s", phrase, translations["mainTranslations"])

	if typeOfWord, ok := translations["typeOfToBeTranslatedWord"]; ok {
		translationText = fmt.Sprintf("Input:\n%s (%s)\n\nMain translation:\n%s", phrase, typeOfWord, translations["mainTranslations"])
	}
	return translationText
}
