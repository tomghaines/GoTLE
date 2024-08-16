package test

import (
	"testing"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func TestTLEChecksum(t *testing.T) {
	l1 := "1 25544U 98067A   21275.48833333  .00016717  00000-0  10270-3 0  9007"
	l2 := "2 25544  51.6431 285.7046 0004107  34.1570 326.0938 15.48815330282953"

	valid, err := tle.VerifyTLELineChecksum(l1, 1, "")
	if err != nil || !valid {
		t.Fatalf("Line 1 failed checksum validation: %v", err)
	}

	valid2, err2 := tle.VerifyTLELineChecksum(l2, 2, "")
	if err2 != nil || !valid2 {
		t.Fatalf("Line 2 failed checksum validation: %v", err2)
	}
}
