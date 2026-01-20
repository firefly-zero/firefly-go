package firefly

import (
	"testing"

	"github.com/orsinium-labs/tinymath"
)

func TestAngleDifference(t *testing.T) {
	tests := []struct {
		name    string
		fromDeg float32
		toDeg   float32
		wantDeg float32
	}{
		{fromDeg: 0, toDeg: 0, wantDeg: 0},
		{fromDeg: 0, toDeg: 90, wantDeg: 90},
		{fromDeg: 45, toDeg: 90, wantDeg: 45},
		{fromDeg: 0, toDeg: 179, wantDeg: 179},
		{fromDeg: 179, toDeg: 0, wantDeg: -179},
		{fromDeg: 0, toDeg: -179, wantDeg: -179},
		{fromDeg: -179, toDeg: 0, wantDeg: 179},
		{fromDeg: 720, toDeg: 360, wantDeg: 0},
		{fromDeg: 700, toDeg: 650, wantDeg: -50},
		// bug case in "tinymath.RemEuclid":
		{fromDeg: 19.000001, toDeg: 19, wantDeg: 0},
	}

	for _, test := range tests {
		result := Degrees(test.fromDeg).Difference(Degrees(test.toDeg))
		resultDeg := tinymath.Round(result.Degrees())
		if resultDeg != test.wantDeg {
			t.Errorf("AngleDifference(%f°, %f°)\nwant: %f°\ngot:  %f°", test.fromDeg, test.toDeg, test.wantDeg, resultDeg)
		}
	}
}

func TestRotateTowards(t *testing.T) {
	tests := []struct {
		name     string
		fromDeg  float32
		toDeg    float32
		deltaDeg float32
		wantDeg  float32
	}{
		{fromDeg: 0, toDeg: 0, deltaDeg: 0, wantDeg: 0},
		{fromDeg: 19, toDeg: 19, deltaDeg: 3, wantDeg: 19},
		// bug case in "tinymath.RemEuclid":
		{fromDeg: 19.000001, toDeg: 19, deltaDeg: 3, wantDeg: 19},
	}

	for _, test := range tests {
		result := Degrees(test.fromDeg).RotateTowards(Degrees(test.toDeg), Degrees(test.deltaDeg))
		resultDeg := result.Degrees()
		if resultDeg != test.wantDeg {
			t.Errorf("RotateTowards(%f°, %f°, %f°)\nwant: %f°\ngot:  %f°", test.fromDeg, test.toDeg, test.deltaDeg, test.wantDeg, resultDeg)
		}
	}
}
