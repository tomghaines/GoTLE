package test

import (
	"testing"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func TestReadTLE(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedName   string
		expectedError  bool
		verifyChecksum bool
	}{
		{
			name:           "Valid TLE with Name",
			input:          "ISS (ZARYA)\n1 25544U 98067A   21275.48833333  .00016717  00000-0  10270-3 0  9007\n2 25544  51.6431 285.7046 0004107  34.1570 326.0938 15.48815330282953",
			expectedName:   "ISS (ZARYA)",
			expectedError:  false,
			verifyChecksum: true,
		},
		{
			name:           "Valid TLE without Name",
			input:          "ISS (ZARYA)\n1 25544U 98067A   21275.48833333  .00016717  00000-0  10270-3 0  9007\n2 25544  51.6431 285.7046 0004107  34.1570 326.0938 15.48815330282953",
			expectedName:   "ISS (ZARYA)",
			expectedError:  false,
			verifyChecksum: true,
		},
		{
			name:           "Invalid TLE Line Count",
			input:          "1 25544U 98067A   21275.48833333  .00016717  00000-0  10270-3 0  9007\n2 25544  51.6431 285.7046 0004107  34.1570 326.0938 15.48815330282953\nExtra Line",
			expectedName:   "",
			expectedError:  true,
			verifyChecksum: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tleResult, err := tle.ReadTLE(tt.input, tt.verifyChecksum)

			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}

			if !tt.expectedError && err == nil {
				if tleResult.Name != tt.expectedName {
					t.Errorf("Expected Name: %v, got: %v", tt.expectedName, tleResult.Name)
				}
			}
		})
	}
}
