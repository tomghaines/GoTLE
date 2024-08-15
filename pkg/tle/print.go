package tle

import (
	"fmt"
	"io"
)

// Prints the TLE information to the provided io.Writer
func Print(w io.Writer, types *TLE) {
	fmt.Fprintln(w, "TLE:")

	// Print TLE information
	fmt.Fprintf(w, "Name                     : %s\n", types.Name)
	fmt.Fprintf(w, "Satellite Number          : %d\n", types.SatelliteNumber)
	fmt.Fprintf(w, "International Designator : %s\n", types.InternationalDesignator)
	fmt.Fprintf(w, "Epoch (Year / Day)       : %d / %.8f\n", types.EpochYear, types.EpochDay)
	fmt.Fprintf(w, "Element Set Number       : %d\n", types.ElementSetNumber)
	fmt.Fprintf(w, "Eccentricity             : %.8f\n", types.Eccentricity)
	fmt.Fprintf(w, "Inclination              : %.8f deg\n", types.Inclination)
	fmt.Fprintf(w, "RAAN                     : %.8f deg\n", types.RAAN)
	fmt.Fprintf(w, "Argument of Perigee      : %.8f deg\n", types.ArgumentOfPerigee)
	fmt.Fprintf(w, "Mean Anomaly             : %.8f deg\n", types.MeanAnomaly)
	fmt.Fprintf(w, "Mean Motion (n)          : %.8f revs / day\n", types.MeanMotion)
	fmt.Fprintf(w, "Revolution Number        : %d\n", types.RevolutionNumber)
	fmt.Fprintf(w, "B*                       : %.8f 1 / er\n", types.Bstar)
	fmt.Fprintf(w, "ṅ / 2                   : %.8f rev / day²\n", types.DnO2)
	fmt.Fprintf(w, "n̈ / 6                   : %.8f rev / day³\n", types.DdnO6)
}
