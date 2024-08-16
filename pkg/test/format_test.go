package test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func TestFormatTLE(t *testing.T) {
	sampleTLE := &tle.TLE{
		Name:                    "ISS (ZARYA)",
		SatelliteNumber:         25544,
		InternationalDesignator: "98067A",
		EpochYear:               2021,
		EpochDay:                275.48833333,
		ElementSetNumber:        9007,
		Eccentricity:            0.00016717,
		Inclination:             51.6431,
		RAAN:                    285.7046,
		ArgumentOfPerigee:       34.1570,
		MeanAnomaly:             326.0938,
		MeanMotion:              15.48815330,
		RevolutionNumber:        27362,
		Bstar:                   0.00000000,
		DnO2:                    0.00000000,
		DdnO6:                   0.00000000,
	}

	expectedOutput := `TLE:
Name                     : ISS (ZARYA)
Satellite Number         : 25544
International Designator : 98067A
Epoch (Year / Day)       : 2021 / 275.48833333
Element Set Number       : 9007
Eccentricity             : 0.00016717
Inclination              : 51.64310000 deg
RAAN                     : 285.70460000 deg
Argument of Perigee      : 34.15700000 deg
Mean Anomaly             : 326.09380000 deg
Mean Motion (n)          : 15.48815330 revs / day
Revolution Number        : 27362
B*                       : 0.00000000 1 / er
ṅ / 2                    : 0.00000000 rev / day²
n̈ / 6                    : 0.00000000 rev / day³`

	var buf bytes.Buffer
	buf.WriteString(tle.FormatTLE(sampleTLE))

	actualOutput := strings.TrimSpace(buf.String())
	expectedOutputTrimmed := strings.TrimSpace(expectedOutput)

	if actualOutput != expectedOutputTrimmed {
		t.Errorf("FormatTLE() output = %v, want %v", actualOutput, expectedOutputTrimmed)
	}
}
