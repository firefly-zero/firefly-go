package firefly

//go:wasmimport graphics get_screen_size
func getScreenSize() int32

//go:wasmimport graphics draw_point
func drawPoint(x, y, c int32)

//go:wasmimport graphics draw_triangle
func drawTriangle(
	p1_x, p1_y, p2_x, p2_y, p3_x, p3_y, fill_color, stroke_color, stroke_width int32,
)
