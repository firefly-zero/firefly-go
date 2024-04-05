package firefly

import "unsafe"

type Point struct {
	X int32
	Y int32
}

type Size struct {
	W uint32
	H uint32
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
	StrokeWidth uint32
}

func Clear(c Color) {
	clearScreen(int32(c))
}

func GetScreenSize() Size {
	raw := getScreenSize()
	return Size{
		W: uint32((raw >> 16) & 0xffff),
		H: uint32(raw & 0xffff),
	}
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

func DrawCircle(p Point, d uint32, s Style) {
	drawCircle(
		p.X, p.Y, int32(d),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

func DrawText(t string, f Font, p Point, c Color) {
	textBytes := []byte(t)
	textPtr := unsafe.Pointer(unsafe.SliceData(textBytes))
	drawText(
		textPtr, uint32(len(t)),
		unsafe.Pointer(&f.raw), uint32(len(f.raw)),
		p.X, p.Y, int32(c),
	)
}
