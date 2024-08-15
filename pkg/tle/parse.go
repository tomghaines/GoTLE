package tle

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Parse the TLE data from two lines and an optional name
func Parse(l1 string, l2 string, name string, l1Position int, l2Position int, verifyChecksum bool) (*TLE, error) {

	// First line

	var debugPrefix string

	if l1Position == 0 {
		debugPrefix = ""
	} else {
		debugPrefix = fmt.Sprintf("[Line %d]: ", l1Position)
	}

	// The first line must start with "1 " and have 69 characters
	if len(l1) != 69 || !strings.HasPrefix(l1, "1 ") {
		return fmt.Errorf("Invalid format for line 1: %s", debugPrefix)
	}

	// Verify the checksum
	func someFunction(l1 string, verifyChecksum bool, debugPrefix string) *TLE {
		if verifyChecksum {
			valid, err := verifyTLELineChecksum(l1, 1, debugPrefix)
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
			if !valid {
				return nil
			}
		}
	
		// Continue processing...
		return &TLE{} // Return your TLE object or whatever is appropriate
	}
	

	// Satellite number

	// Second Line

	// The second line must start with "2" and have 69 characters
	if len(l2) != 69 || !strings.HasPrefix(l2, "2 ") {
		return fmt.Errorf("Invalid format for line 2: %s", debugPrefix)
	}

}

/*
Verify the TLE `line` checksum related to the TLE line `line_number`, which can be 1 or 2.

If the checksum is valid, this function returns `true`. Otherwise, it returns `false`.

`debug_prefix` is a string that will be added to the debugging messages.
*/
func verifyTLELineChecksum(line string, lineNumber int, debugPrefix string) (bool, error) {
	// Try to parse the line checksum
	checksum, err := TLETryParse(Int, string(line[68]), 1, debugPrefix, "line: %s checksum", lineNumber)

	if err != nil {
		return false, fmt.Errorf("%sFailed to parse checksum on line %d: %v", debugPrefix, lineNumber, err)
	}

	expectedChecksum := TLELineChecksum(line[:68])

	if checksum != expectedChecksum {
		return false, fmt.Errorf("%sWrong checksum in TLE line %d (expected = %d, found = %d).", debugPrefix, expectedChecksum, checksum)
	}

	return true, nil
}

/*
Try to parse the `input` to type `T`.

If the operation is successful, it returns the parsed value to `input`.  Otherwise, it
prints an error message and returns `nothing`.

`debugPrefix` is a string to be added to the debugging message, `lineNumber` must be the
current TLE line number (1 or 2), and `field` must be the current TLE field that is being
parsed.
*/
func TLETryParse[T any](input string, lineNumber int, debugPrefix string, field string) (T, error) {
    var parsedValue T
    switch any(parsedValue).(type) {
    case int:
        value, err := strconv.Atoi(input)
        if err != nil {
            return parsedValue, fmt.Errorf("%sThe %s in the TLE line %d could not be parsed: %v", debugPrefix, field, lineNumber, err)
        }
        parsedValue = any(value).(T)
    case float64:
        value, err := strconv.ParseFloat(input, 64)
        if err != nil {
            return parsedValue, fmt.Errorf("%sThe %s in the TLE line %d could not be parsed: %v", debugPrefix, field, lineNumber, err)
        }
        parsedValue = any(value).(T)
    default:
        return parsedValue, fmt.Errorf("unsupported type: %T", parsedValue)
    }
    return parsedValue, nil
}

