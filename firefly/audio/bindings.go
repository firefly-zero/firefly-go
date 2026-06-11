package audio

import "unsafe"

// sources (aka generators)

//go:wasmimport audio add_sine
func addSine(parentID uint32, freq, phase float32) uint32

//go:wasmimport audio add_square
func addSquare(parentID uint32, freq, phase float32) uint32

//go:wasmimport audio add_sawtooth
func addSawtooth(parentID uint32, freq, phase float32) uint32

//go:wasmimport audio add_triangle
func addTriangle(parentID uint32, freq, phase float32) uint32

//go:wasmimport audio add_noise
func addNoise(parentID uint32, seed int32) uint32

//go:wasmimport audio add_empty
func addEmpty(parentID uint32) uint32

//go:wasmimport audio add_zero
func addZero(parentID uint32) uint32

//go:wasmimport audio add_file
func addFile(parentID uint32, ptr unsafe.Pointer, size uint32) uint32

// nodes

//go:wasmimport audio add_mix
func addMix(parentID uint32) uint32

//go:wasmimport audio add_all_for_one
func addAllForOne(parentID uint32) uint32

//go:wasmimport audio add_gain
func addGain(parentID uint32, lvl float32) uint32

//go:wasmimport audio add_loop
func addLoop(parentID uint32) uint32

//go:wasmimport audio add_concat
func addConcat(parentID uint32) uint32

//go:wasmimport audio add_pan
func addPan(parentID uint32, lvl float32) uint32

//go:wasmimport audio add_mute
func addMute(parentID uint32) uint32

//go:wasmimport audio add_pause
func addPause(parentID uint32) uint32

//go:wasmimport audio add_track_position
func addTrackPosition(parentID uint32) uint32

//go:wasmimport audio add_low_pass
func addLowPass(parentID uint32, freq float32, q float32) uint32

//go:wasmimport audio add_high_pass
func addHighPass(parentID uint32, freq float32, q float32) uint32

//go:wasmimport audio add_take_left
func addTakeLeft(parentID uint32) uint32

//go:wasmimport audio add_take_right
func addTakeRight(parentID uint32) uint32

//go:wasmimport audio add_swap
func addSwap(parentID uint32) uint32

//go:wasmimport audio add_clip
func addClip(parentID uint32, low, high float32) uint32

// modulators

//go:wasmimport audio mod_linear
func modLinear(nodeID, param uint32, low, high float32, startAt, endAt uint32)

//go:wasmimport audio mod_hold
func modHold(nodeID, param uint32, low, high float32, time uint32)

//go:wasmimport audio mod_adsr
func modAdsr(
	nodeID, param uint32, low, high float32,
	attack, decay, sustain uint32, sustainLevel float32, release uint32,
)

//go:wasmimport audio mod_sine
func modSine(nodeID, param uint32, freq, low, high float32)

//go:wasmimport audio mod_square
func modSquare(nodeID, param uint32, low, high float32, period uint32)

//go:wasmimport audio mod_square
func modSawtooth(nodeID, param uint32, low, high float32, period uint32)

//go:wasmimport audio set_param
func setParam(nodeID, param uint32, val float32)

//go:wasmimport audio reset
func reset(nodeID uint32)

//go:wasmimport audio reset_all
func resetAll(nodeID uint32)

//go:wasmimport audio clear
func clearNode(nodeID uint32)
