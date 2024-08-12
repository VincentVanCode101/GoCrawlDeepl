package cliargs

import (
	"errors"
	"log"
	"os"
	"strings"
)

// FilterNonEmptyArgs processes command-line arguments to remove empty values and returns a cleaned list or an error if none are provided.
func FilterNonEmptyArgs() ([]string, error) {

	var allCliArgs []string = os.Args[1:]
	if len(allCliArgs) == 0 {
		return nil, errors.New("no cli args provided")
	}
	log.Printf("Lenght of cli args: %v\n", len(allCliArgs))
	log.Printf("The cli args are: %v\n", allCliArgs)

	neededCliArgs := allCliArgs[0]
	var words []string = strings.Split(neededCliArgs, "\n")
	var nonEmptyWords []string

	for _, word := range words {
		trimmedWord := strings.TrimSpace(word)
		if trimmedWord != "" {
			nonEmptyWords = append(nonEmptyWords, trimmedWord)
		}
	}

	if len(nonEmptyWords) == 0 {
		return nil, errors.New("no non-empty cli args provided")
	}

	return nonEmptyWords, nil
}
