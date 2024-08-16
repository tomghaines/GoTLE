package tle

import (
	"unicode"
)

// Compute TLE line checksum & confirm TLE validity
func TLELineChecksum(str string) int {
	checksum := 0

	// Ensure we only process the first 68 characters
	for i, c := range str[:68] {
		// Check if `c` is a digit or a minus sign
		if unicode.IsDigit(c) {
			digit := int(c - '0')
			checksum += digit
		} else if c == '-' {
			checksum += 1
		}
	}

	// Return checksum modulo 10
	return checksum % 10
}
