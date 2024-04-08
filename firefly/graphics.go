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

func Clear(c Color) {
	clearScreen(int32(c))
}

func GetScreenSize() Size {
	raw := getScreenSize()
	return Size{
		W: u32((raw >> 16) & 0xffff),
		H: u32(raw & 0xffff),
	}
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

func DrawImage(i Image, p Point, c1, c2, c3, c4 Color) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(c1), int32(c2), int32(c3), int32(c4),
	)
}

func DrawSubImage(i Image, p Point, subP Point, subS Size, c1, c2, c3, c4 Color) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawSubImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(subP.X), int32(subP.Y),
		uint32(subS.W), uint32(subS.H),
		int32(c1), int32(c2), int32(c3), int32(c4),
	)
}
