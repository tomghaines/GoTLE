package tle

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Parse the TLE data from two lines and an optional name
func ParseTLE(l1 string, l2 string, name string, l1Position int, l2Position int, verifyChecksum bool) (*TLE, error) {

	debugPrefix := ""

	if l1Position != 0 {
		debugPrefix = ""
	} else {
		debugPrefix = fmt.Sprintf("[Line %d]: ", l1Position)
	}

	// First line

	// The first line must start with "1 " and have 69 characters
	if len(l1) != 69 || l1[:2] != "1 " {
		return nil, fmt.Errorf("invalid format for line 1: %s", debugPrefix)
	}

	// Verify the checksum

	if verifyChecksum {
		valid, err := VerifyTLELineChecksum(l1, 1, debugPrefix)
		if err != nil || !valid {
			return nil, err
		}
	}

	// Satellite Number

	satelliteNumber, err := TLETryParse[int](l1[2:7], 1, debugPrefix, "satellite number")
	if err != nil {
		return nil, err
	}

	// Classification

	classifcation := rune(l1[7])

	// International Designator

	internationalDesignator := strings.TrimSpace(l1[9:17])

	// Epoch

	epochYear, err := TLETryParse[int](l1[18:20], 1, debugPrefix, "epoch year")
	if err != nil {
		return nil, err
	}

	epochDay, err := TLETryParse[float64](l1[20:32], 1, debugPrefix, "epoch day")
	if err != nil {
		return nil, err
	}

	// Mean Motion Derivatives

	dnO2, err := TLETryParse[float64](l1[33:43], 1, debugPrefix, "first derivative of mean motion (dn_o6)")
	if err != nil {
		return nil, err
	}

	aux := fmt.Sprintf("%s.%s", l1[44:45], l1[45:50])
	ddnO6Dec, err := TLETryParse[float64](aux, 1, debugPrefix, "second derivative of mean motion (ddn_o6)")
	if err != nil {
		return nil, err
	}

	ddnO6Exp, err := TLETryParse[float64](l1[50:52], 1, debugPrefix, "second derivative of mean motion (ddn_o6)")
	if err != nil {
		return nil, err
	}
	ddnO6 := ddnO6Dec * math.Pow(10, ddnO6Exp)

	// BSTAR

	aux = fmt.Sprintf("%s.%s", l1[53:54], l1[54:59])
	bstarDec, err := TLETryParse[float64](aux, 1, debugPrefix, "BSTAR")
	if err != nil {
		return nil, err
	}

	bstarExp, err := TLETryParse[float64](l1[59:61], 1, debugPrefix, "BSTAR")
	if err != nil {
		return nil, err
	}

	bstar := bstarDec * math.Pow(10, bstarExp)

	// Element Number

	elementSetNumber, err := TLETryParse[int](l1[64:68], 1, debugPrefix, "element set number")
	if err != nil {
		return nil, err
	}

	// Second Line

	// The second line must start with "2" and have 69 characters
	if len(l2) != 69 || !strings.HasPrefix(l2, "2 ") {
		return nil, fmt.Errorf("invalid format for line 2: %s", debugPrefix)
	}

	// Verify the Checksum

	if verifyChecksum {
		valid, err := VerifyTLELineChecksum(l2, 2, debugPrefix)
		if err != nil || !valid {
			return nil, err
		}
	}

	// Compare Satellite Number with the One in the First Line

	satelliteNumberL2, err := TLETryParse[int](l2[2:7], 2, debugPrefix, "satellite number")
	if err != nil {
		return nil, err
	}

	if satelliteNumberL2 != satelliteNumber {
		return nil, fmt.Errorf("%ssatellite number in line 2 does not match line 1: expected %d, got %d", debugPrefix, satelliteNumber, satelliteNumberL2)
	}

	// Inclination

	inclination, err := TLETryParse[float64](l2[8:16], 2, debugPrefix, "inclination")
	if err != nil {
		return nil, err
	}

	// RAAN

	raan, err := TLETryParse[float64](l2[17:25], 2, debugPrefix, "right ascension of the ascending node (RAAN)")
	if err != nil {
		return nil, err
	}

	// Eccentricity

	eccentricity, err := TLETryParse[float64](l2[26:33], 2, debugPrefix, "eccentricity")
	if err != nil {
		return nil, err
	}

	// Argument of Perigee

	argumentOfPerigee, err := TLETryParse[float64](l2[34:42], 2, debugPrefix, "argument of perigee")
	if err != nil {
		return nil, err
	}

	// Mean Anomoly

	meanAnomoly, err := TLETryParse[float64](l2[43:51], 2, debugPrefix, "mean anomoly")
	if err != nil {
		return nil, err
	}

	// Mean Motion

	meanMotion, err := TLETryParse[float64](l2[52:63], 2, debugPrefix, "mean motion")
	if err != nil {
		return nil, err
	}

	// Revolution Number at Epoch
	revolutionNumber, err := TLETryParse[int](l2[63:68], 2, debugPrefix, "revolution number")
	if err != nil {
		return nil, err
	}

	// Create TLE

	return &TLE{
		Name:                    name,
		SatelliteNumber:         satelliteNumber,
		Classification:          classifcation,
		InternationalDesignator: internationalDesignator,
		EpochYear:               epochYear,
		EpochDay:                epochDay,
		DnO2:                    dnO2,
		DdnO6:                   ddnO6,
		Bstar:                   bstar,
		ElementSetNumber:        elementSetNumber,
		Inclination:             inclination,
		RAAN:                    raan,
		Eccentricity:            eccentricity,
		ArgumentOfPerigee:       argumentOfPerigee,
		MeanAnomaly:             meanAnomoly,
		MeanMotion:              meanMotion,
		RevolutionNumber:        revolutionNumber,
	}, nil
}

/*
Verify the TLE `line` checksum related to the TLE line `line_number`, which can be 1 or 2.

If the checksum is valid, this function returns `true`. Otherwise, it returns `false`.

`debug_prefix` is a string that will be added to the debugging messages.
*/
func VerifyTLELineChecksum(line string, lineNumber int, debugPrefix string) (bool, error) {
	// Try to parse the line checksum
	checksumStr := string(line[68])
	checksum, err := strconv.Atoi(checksumStr)
	if err != nil {
		return false, fmt.Errorf("%sThe line %d checksum could not be parsed: %v", debugPrefix, lineNumber, err)
	}

	expectedChecksum := TLELineChecksum(line[:68])
	if checksum != expectedChecksum {
		return false, fmt.Errorf("%sWrong checksum in TLE line %d (expected = %d, found = %d)", debugPrefix, lineNumber, expectedChecksum, checksum)
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
