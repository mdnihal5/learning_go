// Package readfile provides a function to read a file.
package readfile

import (
	"os"
	"strings"
)

func ReadWords(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	array := strings.Fields(string(data))
	return array, nil
}
