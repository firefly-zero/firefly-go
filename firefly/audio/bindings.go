package audio

import "unsafe"

// sources (aka generators)

//go:wasmimport audio add_sine
func addSine(parentID uint32, freq float32, phase float32) uint32

//go:wasmimport audio add_square
func addSquare(parentID uint32, freq float32, phase float32) uint32

//go:wasmimport audio add_sawtooth
func addSawtooth(parentID uint32, freq float32, phase float32) uint32

//go:wasmimport audio add_triangle
func addTriangle(parentID uint32, freq float32, phase float32) uint32

//go:wasmimport audio add_noise
func addNoise(parentID uint32, seed int32) uint32

//go:wasmimport audio add_empty
func addEmpty(parentID uint32) uint32

//go:wasmimport audio add_zero
func addZero(parentID uint32) uint32

//go:wasmimport audio add_file
func addFile(parentID uint32, ptr unsafe.Pointer, len uint32) uint32

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
func addClip(parentID uint32, low float32, high float32) uint32

// modulators

//go:wasmimport audio mod_linear
func modLinear(nodeID uint32, param uint32, start float32, end float32, startAt uint32, endAt uint32)

//go:wasmimport audio mod_hold
func modHold(nodeID uint32, param uint32, v1 float32, v2 float32, time uint32)

//go:wasmimport audio mod_sine
func modSine(nodeID uint32, param uint32, freq float32, low float32, high float32)

//go:wasmimport audio reset
func reset(nodeID uint32)

//go:wasmimport audio reset_all
func resetAll(nodeID uint32)

//go:wasmimport audio clear
func clear(nodeID uint32)
