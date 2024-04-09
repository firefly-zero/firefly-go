package firefly

import "unsafe"

// -- GRAPHICS -- //

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

//go:wasmimport graphics draw_image
func drawImage(
	ptr unsafe.Pointer, len uint32,
	x, y, c1, c2, c3, c4 int32,
)

//go:wasmimport graphics draw_sub_image
func drawSubImage(
	ptr unsafe.Pointer, len uint32,
	x, y, sub_x, sub_y int32, sub_width, sub_height uint32,
	c1, c2, c3, c4 int32,
)

// -- INPUT -- //

//go:wasmimport input read_left
func readLeft() int32

// -- FS -- //

//go:wasmimport fs load_rom_file
func loadRomFile(
	pathPtr unsafe.Pointer, pathLen uint32,
	bufPtr unsafe.Pointer, bufLen uint32,
) uint32

// -- MISC -- //

//go:wasmimport misc log_debug
func logDebug(ptr unsafe.Pointer, len uint32)
