package test

import (
	"bytes"
	"testing"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func TestPrintTLE(t *testing.T) {
	// Sample TLE struct
	sampleTLE := tle.TLE{
		Name:                    "ISS (ZARYA)",
		SatelliteNumber:         25544,
		InternationalDesignator: "98067A",
		EpochYear:               2021,
		EpochDay:                264.44882277,
		ElementSetNumber:        999,
		Eccentricity:            0.0008726,
		Inclination:             51.6457,
		RAAN:                    10.2328,
		ArgumentOfPerigee:       85.7921,
		MeanAnomaly:             274.2343,
		MeanMotion:              15.48970678,
		RevolutionNumber:        27362,
		Bstar:                   0.00031947,
		DnO2:                    0.00001345,
		DdnO6:                   0.0,
	}

	// Capture the output
	var buf bytes.Buffer
	tle.PrintTLE(&buf, &sampleTLE)

	// Expected output string
	expectedOutput := `TLE:
Name                     : ISS (ZARYA)
Satellite Number          : 25544
International Designator : 98067A
Epoch (Year / Day)       : 2021 / 264.44882277
Element Set Number       : 999
Eccentricity             : 0.00087260
Inclination              : 51.64570000 deg
RAAN                     : 10.23280000 deg
Argument of Perigee      : 85.79210000 deg
Mean Anomaly             : 274.23430000 deg
Mean Motion (n)          : 15.48970678 revs / day
Revolution Number        : 27362
B*                       : 0.00031947 1 / er
ṅ / 2                   : 0.00001345 rev / day²
n̈ / 6                   : 0.00000000 rev / day³
`

	// Compare the outputs
	if buf.String() != expectedOutput {
		t.Errorf("PrintTLE() output = %v, want %v", buf.String(), expectedOutput)
	}
}
