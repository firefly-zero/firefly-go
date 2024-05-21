package firefly

import "math"

const signMask uint32 = 0x8000_0000

func sqrt(x float32) float32 {
	r := (math.Float32bits(x) + 0x3f80_0000) >> 1
	return math.Float32frombits(r)
}

// Approximates `atan2(y,x)` normalized to the `[0, 4)` range with a maximum
// error of `0.1620` degrees.
func atan2Norm(y float32, x float32) float32 {
	const B = 0.596_227

	// Extract sign bits from floating point values
	uxS := signMask & math.Float32bits(x)
	uyS := signMask & math.Float32bits(y)

	// Determine quadrant offset
	q := float32((^uxS&uyS)>>29 | uxS>>30)

	// Calculate arctangent in the first quadrant
	bxyA := abs(B * x * y)
	n := bxyA + y*y
	atan1q := n / (x*x + bxyA + n)

	// Translate it to the proper quadrant
	uatan2q := (uxS ^ uyS) | math.Float32bits(atan1q)
	return q + math.Float32frombits(uatan2q)
}

func abs(self float32) float32 {
	return math.Float32frombits(math.Float32bits(self) & ^signMask)
}
