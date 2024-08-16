package tle

import (
	"unicode"
)

// Function to compute TLE line checksum & confirm TLE validity

/*
Compute the checksum of the line `str` modulo 10.

The algorithm is simple: add all the numbers in the line, ignoring letters, spaces, periods,
and plus signs, but assigning +1 to the minus signs. The checksum is the remainder of the
division by 10.
*/

func TLELineChecksum(str string) int {
	checksum := 0

	// Ensure we only process the first 68 characters
	for _, c := range str[:68] {
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
