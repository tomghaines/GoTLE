package tle

import (
	"errors"
	"strings"
)

func Filter(slice []string, condition func(string) bool) []string {
	var result []string
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func ReadTLE(str string, verifyChecksum bool) (*TLE, error) {
	// Split the string into lines
	lines := strings.Split(str, "\n")

	// Filter out empty lines and lines starting with '#'
	var filteredLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			filteredLines = append(filteredLines, line)
		}
	}

	numLines := len(filteredLines)

	// Check if the number of lines is 2 or 3
	if numLines != 2 && numLines != 3 {
		return nil, errors.New("the string must contain only one TLE (2 or 3 lines)")
	}

	// Parse the TLE lines
	var tle *TLE
	var err error
	if numLines == 2 {
		tle, err = ParseTLE(filteredLines[0], filteredLines[1], "UNDEFINED", 1, 2, verifyChecksum)
	} else {
		tle, err = ParseTLE(filteredLines[1], filteredLines[2], filteredLines[0], 1, 2, verifyChecksum)
	}

	if err != nil {
		return nil, err
	}

	if tle == nil {
		return nil, errors.New("the TLE is not valid")
	}

	return tle, nil
}
