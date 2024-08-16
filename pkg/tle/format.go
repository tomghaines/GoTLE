package tle

import "fmt"

// FormatTLE formats a TLE struct into a formatted string
func FormatTLE(tle *TLE) string {
	return fmt.Sprintf(`TLE:
Name                     : %s
Satellite Number         : %d
International Designator : %s
Epoch (Year / Day)       : %d / %.8f
Element Set Number       : %d
Eccentricity             : %.8f
Inclination              : %.8f deg
RAAN                     : %.8f deg
Argument of Perigee      : %.8f deg
Mean Anomaly             : %.8f deg
Mean Motion (n)          : %.8f revs / day
Revolution Number        : %d
B*                       : %.8f 1 / er
ṅ / 2                    : %.8f rev / day²
n̈ / 6                    : %.8f rev / day³`,
		tle.Name,
		tle.SatelliteNumber,
		tle.InternationalDesignator,
		tle.EpochYear,
		tle.EpochDay,
		tle.ElementSetNumber,
		tle.Eccentricity,
		tle.Inclination,
		tle.RAAN,
		tle.ArgumentOfPerigee,
		tle.MeanAnomaly,
		tle.MeanMotion,
		tle.RevolutionNumber,
		tle.Bstar,
		tle.DnO2,
		tle.DdnO6,
	)
}
