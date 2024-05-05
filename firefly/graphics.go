package firefly

import (
	"unsafe"
)

const (
	// The screen width in pixels.
	Width = 240

	// The screen height in pixels.
	Height = 160
)

type Point struct {
	X int
	Y int
}

func (p Point) Size() Size {
	return Size{W: p.X, H: p.Y}
}

func (p Point) Pad() Pad {
	return Pad(p)
}

func (p Point) Abs() Point {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}
	return p
}

func (p Point) Add(r Point) Point {
	p.X += r.X
	p.Y += r.Y
	return p
}

func (p Point) Sub(r Point) Point {
	p.X -= r.X
	p.Y -= r.Y
	return p
}

func (p Point) ComponentMin(r Point) Point {
	if r.X < p.X {
		p.X = r.X
	}
	if r.Y < p.Y {
		p.Y = r.Y
	}
	return p
}

func (p Point) ComponentMax(r Point) Point {
	if r.X > p.X {
		p.X = r.X
	}
	if r.Y > p.Y {
		p.Y = r.Y
	}
	return p
}

type Size struct {
	W int
	H int
}

func (s Size) Point() Point {
	return Point{X: s.W, Y: s.H}
}

func (s Size) Pad() Pad {
	return Pad{X: s.W, Y: s.H}
}

func (s Size) Abs() Size {
	if s.W < 0 {
		s.W = -s.W
	}
	if s.H < 0 {
		s.H = -s.H
	}
	return s
}

func (s Size) Add(r Size) Size {
	s.W += r.W
	s.H += r.H
	return s
}

func (s Size) Sub(r Size) Size {
	s.W -= r.W
	s.H -= r.H
	return s
}

func (s Size) ComponentMin(r Size) Size {
	if r.W < s.W {
		s.W = r.W
	}
	if r.H < s.H {
		s.H = r.H
	}
	return s
}

func (s Size) ComponentMax(r Size) Size {
	if r.W > s.W {
		s.W = r.W
	}
	if r.H > s.H {
		s.H = r.H
	}
	return s
}

type Color uint8

const (
	ColorNone      Color = 0
	ColorDark      Color = 1
	ColorAccent    Color = 2
	ColorSecondary Color = 3
	ColorLight     Color = 4
)

type RGB struct {
	R uint8
	G uint8
	B uint8
}

type Style struct {
	FillColor   Color
	StrokeColor Color
	StrokeWidth int
}

type ImageColors struct {
	A Color
	B Color
	C Color
	D Color
}

func ClearScreen(c Color) {
	clearScreen(int32(c))
}

func SetColor(c Color, v RGB) {
	setColor(int32(c), int32(v.R), int32(v.G), int32(v.B))
}

func DrawPoint(p Point, c Color) {
	drawPoint(int32(p.X), int32(p.Y), int32(c))
}

func DrawTriangle(a, b, c Point, s Style) {
	drawTriangle(
		int32(a.X), int32(a.Y), int32(b.X), int32(b.Y), int32(c.X), int32(c.Y),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

func DrawCircle(p Point, d uint32, s Style) {
	drawCircle(
		int32(p.X), int32(p.Y), int32(d),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

func DrawText(t string, f Font, p Point, c Color) {
	textPtr := unsafe.Pointer(unsafe.StringData(t))
	rawPtr := unsafe.Pointer(unsafe.SliceData(f.raw))
	drawText(
		textPtr, uint32(len(t)),
		rawPtr, uint32(len(f.raw)),
		int32(p.X), int32(p.Y), int32(c),
	)
}

func DrawImage(i Image, p Point, c ImageColors) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(c.A), int32(c.B), int32(c.C), int32(c.D),
	)
}

func DrawSubImage(i SubImage, p Point, c ImageColors) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawSubImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(i.point.X), int32(i.point.Y),
		uint32(i.size.W), uint32(i.size.H),
		int32(c.A), int32(c.B), int32(c.C), int32(c.D),
	)
}
