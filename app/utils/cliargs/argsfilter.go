package cliargs

import (
	"errors"
	"log"
	"os"
	"strings"
)

// FilterNonEmptyArgs processes command-line arguments to remove empty values and returns a cleaned list or an error if none are provided.
func FilterNonEmptyArgs() ([]string, error) {

	var cliArgs []string = os.Args[1:]
	if len(cliArgs) == 0 {
		return nil, errors.New("no cli args provided")
	}
	log.Printf("The cli args are: %v\n", cliArgs)

	cliArgsJoined := strings.Join(cliArgs, " ")

	var words []string = strings.Split(cliArgsJoined, "\n")
	var nonEmptyWords []string

	for _, word := range words {
		trimmedWord := strings.TrimSpace(word)
		if trimmedWord != "" {
			nonEmptyWords = append(nonEmptyWords, trimmedWord)
		}
	}

	return nonEmptyWords, nil
}
