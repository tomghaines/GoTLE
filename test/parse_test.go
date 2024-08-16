package test

import (
	"testing"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func TestParseTLE(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		line1          string
		line2          string
		expectedName   string
		expectedError  bool
		verifyChecksum bool
	}{
		{
			name:           "Valid TLE",
			line1:          "1 25544U 98067A   21275.48833333  .00016717  00000-0  10270-3 0  9007",
			line2:          "2 25544  51.6431 285.7046 0004107  34.1570 326.0938 15.48815330282953",
			expectedName:   "ISS (ZARYA)",
			expectedError:  false,
			verifyChecksum: true,
		},
		{
			name:           "Invalid Checksum",
			line1:          "1 25544U 98067A   21275.48833333  .00016717  00000-0  10270-3 0  9007", // Invalid checksum
			line2:          "2 25544  51.6431 285.7046 0004107  34.1570 326.0938 15.48815330282953",
			expectedName:   "ISS (ZARYA)",
			expectedError:  true,
			verifyChecksum: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tleResult, err := tle.ParseTLE(tt.line1, tt.line2, tt.expectedName, 1, 2, tt.verifyChecksum)

			// Check if error is as expected
			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}

			// If no error expected, validate the returned TLE
			if !tt.expectedError && err == nil {
				if tleResult.Name != tt.expectedName {
					t.Errorf("Expected Name: %v, got: %v", tt.expectedName, tleResult.Name)
				}
			}
		})
	}
}
