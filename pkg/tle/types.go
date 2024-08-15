package tle

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
	// Name
	Name string

	// First Line
	SatelliteNumber         int
	Classification          string
	InternationalDesignator string
	EpochYear               int
	EpochDay                float64
	DnO2                    float64
	DdnO6                   float64
	Bstar                   float64
	ElementSetNumber        int

	// Second Line
	Inclination       float64
	RAAN              float64
	Eccentricity      float64
	ArgumentOfPerigee float64
	MeanAnomaly       float64
	MeanMotion        float64
	RevolutionNumber  int
}
