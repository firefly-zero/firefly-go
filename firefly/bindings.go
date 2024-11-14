package firefly

import "unsafe"

// -- GRAPHICS -- //

//go:wasmimport graphics clear_screen
func clearScreen(c int32)

//go:wasmimport graphics set_color
func setColor(c, r, g, b int32)

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
func drawArc(x, y, d int32, ast, asw float32, fc, sc, sw int32)

//go:wasmimport graphics draw_sector
func drawSector(x, y, d int32, ast, asw float32, fc, sc, sw int32)

//go:wasmimport graphics draw_text
func drawText(
	textPtr unsafe.Pointer, textLen uint32,
	fontPtr unsafe.Pointer, fontLen uint32,
	x, y, color int32,
)

//go:wasmimport graphics draw_image
func drawImage(ptr unsafe.Pointer, len uint32, x, y int32)

//go:wasmimport graphics draw_sub_image
func drawSubImage(
	ptr unsafe.Pointer, len uint32,
	x, y, subX, subY int32, subWidth, subHeight uint32,
)

//go:wasmimport graphics set_canvas
func setCanvas(ptr unsafe.Pointer, len uint32)

//go:wasmimport graphics unset_canvas
func unsetCanvas()

// -- INPUT -- //

//go:wasmimport input read_pad
func readPad(player uint32) int32

//go:wasmimport input read_buttons
func readButtons(player uint32) uint32

// -- FS -- //

//go:wasmimport fs get_file_size
func getFileSize(pathPtr unsafe.Pointer, pathLen uint32) uint32

//go:wasmimport fs load_file
func loadFile(
	pathPtr unsafe.Pointer, pathLen uint32,
	bufPtr unsafe.Pointer, bufLen uint32,
) uint32

//go:wasmimport fs dump_file
func dumpFile(
	pathPtr unsafe.Pointer, pathLen uint32,
	bufPtr unsafe.Pointer, bufLen uint32,
) uint32

//go:wasmimport fs remove_file
func removeFile(pathPtr unsafe.Pointer, pathLen uint32) uint32

// -- NET -- //

//go:wasmimport net get_me
func getMe() uint32

//go:wasmimport net get_peers
func getPeers() uint32

// -- STATS -- //

//go:wasmimport stats add_progress
func addProgress(peerID, badgeID uint32, val int32) uint32

//go:wasmimport stats add_score
func addScore(peerID, boardID uint32, val int32) int32

// -- MISC -- //

//go:wasmimport misc log_debug
func logDebug(ptr unsafe.Pointer, len uint32)

//go:wasmimport misc log_error
func logError(ptr unsafe.Pointer, len uint32)

//go:wasmimport misc set_seed
func setSeed(seed uint32)

//go:wasmimport misc get_random
func getRandom() uint32

//go:wasmimport misc restart
func restart()

//go:wasmimport misc quit
func quit()
