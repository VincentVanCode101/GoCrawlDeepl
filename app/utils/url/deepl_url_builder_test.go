package url

import (
	"testing"
)

func TestBuildDeeplURL(t *testing.T) {
	tests := []struct {
		baseURL           string
		fromLang          string
		toLang            string
		phraseToTranslate string
		expected          string
	}{
		{
			baseURL:           "https://www.deepl.com/en/translator#",
			fromLang:          "en",
			toLang:            "de",
			phraseToTranslate: "Hello, world!",
			expected:          "https://www.deepl.com/en/translator#en/de/Hello%2C%20world%21",
		},
		{
			baseURL:           "https://www.deepl.com/en/translator#",
			fromLang:          "en",
			toLang:            "de",
			phraseToTranslate: "car door",
			expected:          "https://www.deepl.com/en/translator#en/de/car%20door",
		},
		{
			baseURL:           "https://www.deepl.com/en/translator#",
			fromLang:          "en",
			toLang:            "de",
			phraseToTranslate: "pasta+scrambled eggs=yummy",
			expected:          "https://www.deepl.com/en/translator#en/de/pasta%2Bscrambled%20eggs%3Dyummy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.phraseToTranslate, func(t *testing.T) {
			actual := BuildDeeplURL(tt.baseURL, tt.fromLang, tt.toLang, tt.phraseToTranslate)
			if actual != tt.expected {
				t.Errorf("BuildDeeplURL(%s, %s, %s, %s) = %s; want %s", tt.baseURL, tt.fromLang, tt.toLang, tt.phraseToTranslate, actual, tt.expected)
			}
		})
	}
}
