package firefly

type Point struct {
	X int32
	Y int32
}

type Size struct {
	W uint32
	H uint32
}

type Color uint8

type RGB struct {
	R uint8
	G uint8
	B uint8
}

type Style struct {
	FillColor   Color
	StrokeColor Color
	StrokeWidth uint32
}

func DrawPoint(p Point, c Color) {
	drawPoint(p.X, p.Y, int32(c))
}

func DrawTriangle(a, b, c Point, s Style) {
	drawTriangle(
		a.X, a.Y, b.X, b.Y, c.X, c.Y,
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}
