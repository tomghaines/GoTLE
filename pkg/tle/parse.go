package tle

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Store the elements of a TLE (two-line elemens) using the same units.

Name:

  - "name": Name of the satellite.

First Line:

  - `SatelliteNumber`: Satellite number.
  - `Classification`: Classification ('U', 'C', or 'S').
  - `InternationalDesignator`: International designator.
  - `EpochYear`: Epoch year (two digits).
  - `EpochDay`: Epoch day (day + fraction of the day).
  - `DnO2`: 1st time derivative of mean motion / 2 [rev/day²].
  - `DdnO6`: 2nd time derivative of mean motion / 6 [rev/day³].
  - `Bstar`: B* drag term.
  - `elementSetNumber`: Element set number.

Second Line:

  - `Incliantion`: Inclination [deg].
  - `RAAN`: Right ascension of the ascending node [deg].
  - `Eccentricity`: Eccentricity [ ].
  - `ArgumentOfPerigee`: Argument of perigee [deg].
  - `MeanAnomaly`: Mean anomaly [deg].
  - `MeanMotion`: Mean motion [rev/day].
  - `RevolutionNumber`: Revolution number at epoch [rev].
*/
type TLE struct {
	Name                    string
	SatelliteNumber         int
	InternationalDesignator string
	EpochYear               int
	EpochDay                float64
	ElementSetNumber        int
	Eccentricity            float64
	Inclination             float64
	RAAN                    float64
	ArgumentOfPerigee       float64
	MeanAnomaly             float64
	MeanMotion              float64
	RevolutionNumber        int
	Bstar                   float64
	DnO2                    float64
	DdnO6                   float64
}

// ParseTLE parses TLE lines and returns a TLE struct
func ParseTLE(name, line1, line2 string) (*TLE, error) {
	if len(line1) < 69 || len(line2) < 69 {
		return nil, fmt.Errorf("invalid TLE format: lines too short")
	}

	tle := &TLE{
		Name: name,
	}

	// Line 1 parsing
	satelliteNumber, err := strconv.Atoi(line1[2:7])
	if err != nil {
		return nil, fmt.Errorf("invalid satellite number in line 1: %v", err)
	}
	tle.SatelliteNumber = satelliteNumber

	tle.InternationalDesignator = strings.TrimSpace(line1[9:17])

	epochYear, err := strconv.Atoi("20" + line1[18:20])
	if err != nil {
		return nil, fmt.Errorf("invalid epoch year in line 1: %v", err)
	}
	tle.EpochYear = epochYear

	epochDay, err := strconv.ParseFloat(line1[20:32], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid epoch day in line 1: %v", err)
	}
	tle.EpochDay = epochDay

	elementSetNumber, err := strconv.Atoi(strings.TrimSpace(line1[64:68]))
	if err != nil {
		return nil, fmt.Errorf("invalid element set number in line 1: %v", err)
	}
	tle.ElementSetNumber = elementSetNumber

	// Line 2 parsing
	inclination, err := strconv.ParseFloat(line2[8:16], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid inclination in line 2: %v", err)
	}
	tle.Inclination = inclination

	raan, err := strconv.ParseFloat(line2[17:25], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid RAAN in line 2: %v", err)
	}
	tle.RAAN = raan

	eccentricity, err := strconv.ParseFloat("0."+line2[26:33], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid eccentricity in line 2: %v", err)
	}
	tle.Eccentricity = eccentricity

	argumentOfPerigee, err := strconv.ParseFloat(line2[34:42], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid argument of perigee in line 2: %v", err)
	}
	tle.ArgumentOfPerigee = argumentOfPerigee

	meanAnomaly, err := strconv.ParseFloat(line2[43:51], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid mean anomaly in line 2: %v", err)
	}
	tle.MeanAnomaly = meanAnomaly

	meanMotion, err := strconv.ParseFloat(line2[52:63], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid mean motion in line 2: %v", err)
	}
	tle.MeanMotion = meanMotion

	revolutionNumber, err := strconv.Atoi(strings.TrimSpace(line2[63:68]))
	if err != nil {
		return nil, fmt.Errorf("invalid revolution number in line 2: %v", err)
	}
	tle.RevolutionNumber = revolutionNumber

	// Optional fields (defaults to 0.0 if not present)
	tle.Bstar = 0.0
	tle.DnO2 = 0.0
	tle.DdnO6 = 0.0

	return tle, nil
}
