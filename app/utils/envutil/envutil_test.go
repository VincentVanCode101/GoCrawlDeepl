package envutil

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGetLanguages(t *testing.T) {
	// Save the original environment variables to restore later
	originalFromLang := os.Getenv("FROM_LANGUAGE")
	originalToLang := os.Getenv("TO_LANGUAGE")

	// Restore the original environment variables after the test
	defer func() {
		os.Setenv("FROM_LANGUAGE", originalFromLang)
		os.Setenv("TO_LANGUAGE", originalToLang)
	}()

	tests := []struct {
		name              string
		envFromLang       string
		envToLang         string
		expectedFromLang  string
		expectedToLang    string
		expectedLogOutput []string
	}{
		{
			name:              "both environment variables set",
			envFromLang:       "fr",
			envToLang:         "es",
			expectedFromLang:  "fr",
			expectedToLang:    "es",
			expectedLogOutput: []string{},
		},
		{
			name:              "only FROM_LANGUAGE set",
			envFromLang:       "it",
			envToLang:         "",
			expectedFromLang:  "it",
			expectedToLang:    "de",
			expectedLogOutput: []string{"No environment variable for TO_LANGUAGE found. Defaulting to 'de'."},
		},
		{
			name:              "only TO_LANGUAGE set",
			envFromLang:       "",
			envToLang:         "pt",
			expectedFromLang:  "en",
			expectedToLang:    "pt",
			expectedLogOutput: []string{"No environment variable for FROM_LANGUAGE found. Defaulting to 'en'."},
		},
		{
			name:             "neither environment variable set",
			envFromLang:      "",
			envToLang:        "",
			expectedFromLang: "en",
			expectedToLang:   "de",
			expectedLogOutput: []string{
				"No environment variable for FROM_LANGUAGE found. Defaulting to 'en'.",
				"No environment variable for TO_LANGUAGE found. Defaulting to 'de'.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("FROM_LANGUAGE", tt.envFromLang)
			os.Setenv("TO_LANGUAGE", tt.envToLang)

			// Capture log output
			var logOutput bytes.Buffer
			log.SetOutput(&logOutput)
			defer log.SetOutput(os.Stderr)

			fromLang, toLang, err := GetLanguages()
			if err != nil {
				t.Fatalf("GetLanguages() returned an error: %v", err)
			}

			if fromLang != tt.expectedFromLang {
				t.Errorf("Expected FROM_LANGUAGE to be %q, but got %q", tt.expectedFromLang, fromLang)
			}
			if toLang != tt.expectedToLang {
				t.Errorf("Expected TO_LANGUAGE to be %q, but got %q", tt.expectedToLang, toLang)
			}

			logged := logOutput.String()
			for _, expected := range tt.expectedLogOutput {
				if !strings.Contains(logged, expected) {
					t.Errorf("Expected log output to contain %q, but got %q", expected, logged)
				}
			}
		})
	}
}
