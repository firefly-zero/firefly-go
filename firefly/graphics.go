package firefly

import (
	"unsafe"
)

// Use just int for both signed and unsigned 32-bit integers.
//
// It simplifies API: X is compatible with Width and len is assignable to X
// without an explicit type conversion.
//
// Since wasm is 32-bit system without uints, converting int to uint32 or int32
// is no-op in runtime.
type i32 = int
type u32 = int

type Point struct {
	X i32
	Y i32
}

type Size struct {
	W u32
	H u32
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
	StrokeWidth u32
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
