package firefly

import "unsafe"

// -- GRAPHICS -- //

//go:wasmimport graphics clear_screen
func clearScreen(c int32)

//go:wasmimport graphics get_screen_size
func getScreenSize() int32

//go:wasmimport graphics set_color
func setColor(c, r, g, b int32)

//go:wasmimport graphics set_colors
func setColors(
	r1, g1, b1,
	r2, g2, b2,
	r3, g3, b3,
	r4, g4, b4 int32,
)

//go:wasmimport graphics draw_point
func drawPoint(x, y, c int32)

//go:wasmimport graphics draw_line
func drawLine(x1, y1, x2, y2, c, sw int32)

//go:wasmimport graphics draw_rect
func drawRect(x, y, w, h, fc, sc, sw int32)

//go:wasmimport graphics draw_rounded_rect
func drawRoundedRect(x, y, w, h, cw, ch, fc, sc, sw int32)

//go:wasmimport graphics draw_circle
func drawCircle(x, y, d, fc, sc, sw int32)

//go:wasmimport graphics draw_ellipse
func drawEllipse(x, y, w, h, fc, sc, sw int32)

//go:wasmimport graphics draw_triangle
func drawTriangle(x1, y1, x2, y2, x3, y3, fc, sc, sw int32)

//go:wasmimport graphics draw_arc
func drawArc(x, y, d, ast, asw, fc, sc, sw int32)

//go:wasmimport graphics draw_sector
func drawSector(x, y, d, ast, asw, fc, sc, sw int32)

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

//go:wasmimport input read_pad
func readPad(player uint32) int32

// -- FS -- //

//go:wasmimport fs get_rom_file_size
func getRomFileSize(pathPtr unsafe.Pointer, pathLen uint32) uint32

//go:wasmimport fs load_rom_file
func loadRomFile(
	pathPtr unsafe.Pointer, pathLen uint32,
	bufPtr unsafe.Pointer, bufLen uint32,
) uint32

// -- MISC -- //

//go:wasmimport misc log_debug
func logDebug(ptr unsafe.Pointer, len uint32)
