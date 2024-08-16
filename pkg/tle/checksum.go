package tle

import (
	"unicode"
)

// TLELineChecksum calculates the checksum of a TLE line, as per TLE format specifications.
func TLELineChecksum(str string) int {
	checksum := 0

	// Iterate through all characters except the checksum digit (last character)
	for _, c := range str[:len(str)-1] {
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

// ValidateTLELine checks if the TLE line's checksum is valid.
func ValidateTLELine(line string) (bool, int, int) {
	if len(line) < 69 {
		return false, -1, -1 // Line too short for valid TLE line
	}

	expectedChecksum := int(line[68] - '0')
	calculatedChecksum := TLELineChecksum(line)

	return expectedChecksum == calculatedChecksum, expectedChecksum, calculatedChecksum
}
