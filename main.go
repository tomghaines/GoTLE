package main

import (
	"fmt"
	"log"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func main() {
	name := "ISS (ZARYA)"
	line1 := "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927"
	line2 := "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125392  2927"

	tleData, err := tle.ParseTLE(name, line1, line2)
	if err != nil {
		log.Fatalf("Error parsing TLE: %v", err)
	}

	formattedTLE := tle.FormatTLE(tleData)
	fmt.Println(formattedTLE)
}
