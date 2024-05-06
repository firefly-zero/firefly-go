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
