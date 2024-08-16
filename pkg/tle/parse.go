package tle

import (
	"fmt"
	"strconv"
)

// ParseTLE parses TLE lines and returns a TLE struct
func ParseTLE(name, line1, line2 string, l1Pos, l2Pos int) (*TLE, error) {
	if len(line1) < 69 || len(line2) < 69 {
		return nil, fmt.Errorf("invalid TLE format")
	}

	tle := &TLE{
		Name: name,
	}

	// Line 1 parsing
	satelliteNumber, err := strconv.Atoi(line1[2:7])
	if err != nil {
		return nil, fmt.Errorf("invalid satellite number in line 1")
	}
	tle.SatelliteNumber = satelliteNumber

	tle.InternationalDesignator = line1[9:17]

	epochYear, err := strconv.Atoi("20" + line1[18:20])
	if err != nil {
		return nil, fmt.Errorf("invalid epoch year in line 1")
	}
	tle.EpochYear = epochYear

	epochDay, err := strconv.ParseFloat(line1[20:32], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid epoch day in line 1")
	}
	tle.EpochDay = epochDay

	elementSetNumber, err := strconv.Atoi(line1[64:68])
	if err != nil {
		return nil, fmt.Errorf("invalid element set number in line 1")
	}
	tle.ElementSetNumber = elementSetNumber

	// Line 2 parsing
	inclination, err := strconv.ParseFloat(line2[8:16], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid inclination in line 2")
	}
	tle.Inclination = inclination

	raan, err := strconv.ParseFloat(line2[17:25], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid RAAN in line 2")
	}
	tle.RAAN = raan

	eccentricity, err := strconv.ParseFloat("0."+line2[26:33], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid eccentricity in line 2")
	}
	tle.Eccentricity = eccentricity

	argumentOfPerigee, err := strconv.ParseFloat(line2[34:42], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid argument of perigee in line 2")
	}
	tle.ArgumentOfPerigee = argumentOfPerigee

	meanAnomaly, err := strconv.ParseFloat(line2[43:51], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid mean anomaly in line 2")
	}
	tle.MeanAnomaly = meanAnomaly

	meanMotion, err := strconv.ParseFloat(line2[52:63], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid mean motion in line 2")
	}
	tle.MeanMotion = meanMotion

	revolutionNumber, err := strconv.Atoi(line2[63:68])
	if err != nil {
		return nil, fmt.Errorf("invalid revolution number in line 2")
	}
	tle.RevolutionNumber = revolutionNumber

	// Optional fields (defaults to 0.0 if not present)
	tle.Bstar = 0.0
	tle.DnO2 = 0.0
	tle.DdnO6 = 0.0

	return tle, nil
}
