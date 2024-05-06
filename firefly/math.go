package firefly

import "math"

const uvnan = 0x7FE00000

func sqrt(x float32) float32 {
	// https://github.com/tarcieri/micromath/blob/main/src/float/sqrt.rs
	if x >= 0. {
		r := (math.Float32bits(x) + 0x3f80_0000) >> 1
		return math.Float32frombits(r)
	}
	return math.Float32frombits(uvnan)
}

// Approximates `atan(x)` approximation in radians with a maximum error of `0.002`.
func atan(x float32) float32 {
	// https://github.com/tarcieri/micromath/blob/main/src/float/atan.rs
	return math.Pi / 2. * atanNorm(x)
}

// Approximates `atan(x)` normalized to the `[−1,1]` range with a maximum
// error of `0.1620` degrees.
func atanNorm(x float32) float32 {
	// https://github.com/tarcieri/micromath/blob/main/src/float/atan.rs
	// Extract the sign bit
	uxS := 0x8000_0000 & math.Float32bits(x)

	// Calculate the arctangent in the first quadrant
	bxA := 0.596227 * x
	if bxA < 0 {
		bxA = -bxA
	}
	n := bxA + x*x
	atan1q := n / (1.0 + bxA + n)

	// Restore the sign bit and convert to float
	return math.Float32frombits(uxS | math.Float32bits(atan1q))
}
