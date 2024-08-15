package tle

import (
	"fmt"
	"strings"
)

// Parse the TLE data from two lines and an optional name
func Parse(l1 string, l2 string, name string) (*TLE, error) {
	// First line

	// The first line must start with "1 " and have 69 characters
	if len(l1) != 69 || !strings.HasPrefix(l1, "1 ") {
		return nil, fmt.Errorf("Invalid format for line 1: %s", l1)
	}

	// Second Line

	// The second line must start with "2" and have 69 characters
	if len(l2) != 69 || !strings.HasPrefix(l1, "2 ") {
		return nil, fmt.Errorf("Invalid format for line 2: %s", l2)
	}

}
