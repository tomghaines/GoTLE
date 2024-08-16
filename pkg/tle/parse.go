package tle

import (
	"fmt"
	"strconv"
)

// TLE represents a Two-Line Element set
type TLE struct {
	Name                    string
	SatelliteNumber         int
	InternationalDesignator string
	EpochYear               int
	EpochDay                float64
	ElementSetNumber        int
	Inclination             float64
	RAAN                    float64
	Eccentricity            float64
	ArgumentOfPerigee       float64
	MeanAnomaly             float64
	MeanMotion              float64
	RevolutionNumber        int
	Bstar                   float64 // B* (drag term) in TLE
	DnO2                    float64 // ṅ / 2 (mean motion derivative) in TLE
	DdnO6                   float64 // n̈ / 6 (mean motion second derivative) in TLE
}

// ParseTLE parses the TLE data and returns a TLE struct
func ParseTLE(name, line1, line2 string) (*TLE, error) {
	// Initialize the TLE struct
	tle := &TLE{
		Name: name,
	}

	// Parse line 1
	if len(line1) < 69 {
		return nil, fmt.Errorf("invalid line 1 length")
	}

	// Extract fields from line 1
	tle.SatelliteNumber, _ = strconv.Atoi(line1[2:7])
	tle.InternationalDesignator = line1[7:14]
	epochYear, _ := strconv.Atoi(line1[14:16])
	epochDay, _ := strconv.ParseFloat(line1[16:32], 64)
	tle.EpochYear = 2000 + epochYear // Assuming the TLE is from the 2000s
	tle.EpochDay = epochDay

	// Extract element set number from line 1
	elementSetNumber, _ := strconv.Atoi(line1[64:68])
	tle.ElementSetNumber = elementSetNumber

	// Parse line 2
	if len(line2) < 69 {
		return nil, fmt.Errorf("invalid line 2 length")
	}

	// Extract fields from line 2
	tle.Inclination, _ = strconv.ParseFloat(line2[8:16], 64)
	tle.RAAN, _ = strconv.ParseFloat(line2[17:25], 64)
	tle.Eccentricity, _ = strconv.ParseFloat("0."+line2[26:33], 64)
	tle.ArgumentOfPerigee, _ = strconv.ParseFloat(line2[34:42], 64)
	tle.MeanAnomaly, _ = strconv.ParseFloat(line2[43:51], 64)
	tle.MeanMotion, _ = strconv.ParseFloat(line2[52:63], 64)
	tle.RevolutionNumber, _ = strconv.Atoi(line2[63:68])
	// B*, ṅ / 2, and n̈ / 6 are not available in line 2
	// They are often zero in standard TLE formats

	return tle, nil
}
