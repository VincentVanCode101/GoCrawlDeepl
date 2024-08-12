package cliargs

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestFilterNonEmptyArgs(t *testing.T) {
	tests := []struct {
		args     []string
		expected []string
		hasError bool
	}{
		{
			args:     []string{"program", "car door plane"},
			expected: []string{"car door plane"},
			hasError: false,
		},
		{
			args:     []string{"program", "car  door  plane"},
			expected: []string{"car  door  plane"},
			hasError: false,
		},
		{
			args:     []string{"program", "car door plane", "arg4"},
			expected: []string{"car door plane"},
			hasError: false,
		},
		{
			args:     []string{"program", " "},
			expected: []string{},
			hasError: true,
		},
		{
			args:     []string{"program", ""},
			expected: []string{},
			hasError: true,
		},
		{
			args:     []string{"program"},
			expected: []string{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			// Save the original os.Args and restore it after the test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set os.Args to the test case arguments
			os.Args = tt.args

			result, err := FilterNonEmptyArgs()

			if (err != nil) != tt.hasError {
				t.Errorf("expected error: %v, got: %v", tt.hasError, err)
			}

			if !tt.hasError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}
