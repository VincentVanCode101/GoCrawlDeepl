package url

import (
	"fmt"
	"net/url"
	"strings"
)

// BuildDeeplURL constructs a URL for accessing the DeepL translation service.
func BuildDeeplURL(baseURL, fromLang, toLang, phraseToTranslate string) string {
	escapedPhrase := url.QueryEscape(phraseToTranslate)
	finalEscapedPhrase := strings.ReplaceAll(escapedPhrase, "+", "%20")
	return fmt.Sprintf("%s%s/%s/%s", baseURL, fromLang, toLang, finalEscapedPhrase)
}
