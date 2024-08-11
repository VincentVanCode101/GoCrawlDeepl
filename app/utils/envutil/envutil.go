package envutil

import (
	"log"
	"os"
)

// GetLanguages retrieves the FROM_LANGUAGE and TO_LANGUAGE from the environment variables.
func GetLanguages() (string, string, error) {
	fromLang, isPresent := os.LookupEnv("FROM_LANGUAGE")
	if !isPresent || fromLang == "" {
		log.Println("No environment variable for FROM_LANGUAGE found. Defaulting to 'en'.")
		fromLang = "en"
	}

	toLang, isPresent := os.LookupEnv("TO_LANGUAGE")
	if !isPresent || toLang == "" {
		log.Println("No environment variable for TO_LANGUAGE found. Defaulting to 'de'.")
		toLang = "de"
	}

	return fromLang, toLang, nil
}
