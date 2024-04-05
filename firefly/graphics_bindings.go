package firefly

import "unsafe"

//go:wasmimport graphics clear
func clearScreen(c int32)

//go:wasmimport graphics get_screen_size
func getScreenSize() int32

//go:wasmimport graphics draw_point
func drawPoint(x, y, c int32)

//go:wasmimport graphics draw_triangle
func drawTriangle(
	p1_x, p1_y, p2_x, p2_y, p3_x, p3_y, fill_color, stroke_color, stroke_width int32,
)

//go:wasmimport graphics draw_circle
func drawCircle(x, y, d, fill_color, stroke_color, stroke_width int32)

//go:wasmimport graphics draw_text
func drawText(
	textPtr unsafe.Pointer, textLen uint32,
	fontPtr unsafe.Pointer, fontLen uint32,
	x, y, color int32,
)
