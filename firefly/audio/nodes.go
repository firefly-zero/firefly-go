package audio

import "unsafe"

type Node struct{ SourceNode }

var Out = Node{SourceNode{0}}

// An audio node created by [Node.AddMix].
type Mix struct{ Node }

// An audio node created by [Node.AddAllForOne].
type AllForOne struct{ Node }

// An audio node created by [Node.AddGain].
type Gain struct{ Node }

// An audio node created by [Node.AddLoop].
type Loop struct{ Node }

// An audio node created by [Node.AddConcat].
type Concat struct{ Node }

// An audio node created by [Node.AddPan].
type Pan struct{ Node }

// An audio node created by [Node.AddMute].
type Mute struct{ Node }

// An audio node created by [Node.AddPause].
type Pause struct{ Node }

// An audio node created by [Node.AddTrackPosition].
type TrackPosition struct{ Node }

// An audio node created by [Node.AddLowPass].
type LowPass struct{ Node }

// An audio node created by [Node.AddHighPass].
type HighPass struct{ Node }

// An audio node created by [Node.AddTakeLeft].
type TakeLeft struct{ Node }

// An audio node created by [Node.AddTakeRight].
type TakeRight struct{ Node }

// An audio node created by [Node.AddSwap].
type Swap struct{ Node }

// An audio node created by [Node.AddClip].
type Clip struct{ Node }

// Add sine wave oscillator source (`∿`).
func (n Node) AddSine(freq Freq, phase float32) Sine {
	id := addSine(n.id, float32(freq), phase)
	return Sine{SourceNode{id}}
}

// Add square wave oscillator source (`⎍`).
func (n Node) AddSquare(freq Freq, phase float32) Square {
	id := addSquare(n.id, float32(freq), phase)
	return Square{SourceNode{id}}
}

// Add sawtooth wave oscillator source (`╱│`).
func (n Node) AddSawtooth(freq Freq, phase float32) Sawtooth {
	id := addSawtooth(n.id, float32(freq), phase)
	return Sawtooth{SourceNode{id}}
}

// Add triangle wave oscillator source (`╱╲`).
func (n Node) AddTriangle(freq Freq, phase float32) Triangle {
	id := addTriangle(n.id, float32(freq), phase)
	return Triangle{SourceNode{id}}
}

// Add white noise source (amplitude on each tick is random).
func (n Node) AddNoise(seed int) Noise {
	id := addNoise(n.id, int32(seed))
	return Noise{SourceNode{id}}
}

// Add always stopped source.
func (n Node) AddEmpty() Empty {
	id := addEmpty(n.id)
	return Empty{SourceNode{id}}
}

// Add silent source producing zeros.
func (n Node) AddZero() Zero {
	id := addZero(n.id)
	return Zero{SourceNode{id}}
}

// Add source playing audio from a file.
func (n Node) AddFile(path string) File {
	ptr := unsafe.Pointer(unsafe.StringData(path))
	id := addFile(n.id, ptr, uint32(len(path)))
	return File{SourceNode{id}}
}

// Add node simply mixing all inputs.
func (n Node) AddMix() Mix {
	id := addMix(n.id)
	return Mix{Node{SourceNode{id}}}
}

// Add mixer node that stops if any of the sources stops.
func (n Node) AddAllForOne() AllForOne {
	id := addAllForOne(n.id)
	return AllForOne{Node{SourceNode{id}}}
}

// Add gain control node.
func (n Node) AddGain(lvl float32) Gain {
	id := addGain(n.id, lvl)
	return Gain{Node{SourceNode{id}}}
}

// Add a loop node that resets the input if it stops.
func (n Node) AddLoop() Loop {
	id := addLoop(n.id)
	return Loop{Node{SourceNode{id}}}
}

// Add a node that plays the inputs one after the other, in the order as they added.
func (n Node) AddConcat() Concat {
	id := addConcat(n.id)
	return Concat{Node{SourceNode{id}}}
}

// Add node panning the audio to the left (0.), right (1.), or something in between.
func (n Node) AddPan(lvl float32) Pan {
	id := addPan(n.id, lvl)
	return Pan{Node{SourceNode{id}}}
}

// Add node that can be muted using modulation.
func (n Node) AddMute() Mute {
	id := addMute(n.id)
	return Mute{Node{SourceNode{id}}}
}

// Add node that can be paused using modulation.
func (n Node) AddPause() Pause {
	id := addPause(n.id)
	return Pause{Node{SourceNode{id}}}
}

// Add node tracking the elapsed playback time.
func (n Node) AddTrackPosition() TrackPosition {
	id := addTrackPosition(n.id)
	return TrackPosition{Node{SourceNode{id}}}
}

// Add lowpass filter node.
func (n Node) AddLowPass(freq float32, q float32) LowPass {
	id := addLowPass(n.id, freq, q)
	return LowPass{Node{SourceNode{id}}}
}

// Add highpass filter node.
func (n Node) AddHighPass(freq float32, q float32) HighPass {
	id := addHighPass(n.id, freq, q)
	return HighPass{Node{SourceNode{id}}}
}

// Add node converting stereo to mono by taking the left channel.
func (n Node) AddTakeLeft() TakeLeft {
	id := addTakeLeft(n.id)
	return TakeLeft{Node{SourceNode{id}}}
}

// Add node converting stereo to mono by taking the right channel.
func (n Node) AddTakeRight() TakeRight {
	id := addTakeRight(n.id)
	return TakeRight{Node{SourceNode{id}}}
}

// Add node swapping left and right channels of the stereo input.
func (n Node) AddSwap() Swap {
	id := addSwap(n.id)
	return Swap{Node{SourceNode{id}}}
}

// Add node clamping the input amplitude. Can be used for hard distortion.
func (n Node) AddClip(low float32, high float32) Clip {
	id := addClip(n.id, low, high)
	return Clip{Node{SourceNode{id}}}
}

// Reset the node and all child nodes to the state to how it was when they were just added.
func (n Node) ResetAll() {
	resetAll(n.id)
}

// Remove all child nodes.
//
// After it is called, you should make sure to discard all references to the old
// child nodes.
func (n Node) Clear() {
	clearNode(n.id)
}

// Modulate the gain level.
func (n Gain) Modulate(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the pan value (from 0. to 1.: 0. is only left, 1. is only right).
func (n Pan) Modulate(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the muted state.
//
// Below 0.5 is muted, above is unmuted.
func (n Mute) Modulate(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the paused state.
//
// Below 0.5 is paused, above is playing.
func (n Pause) Modulate(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the cut-off frequency.
func (n LowPass) ModulateFreq(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the cut-off frequency.
func (n HighPass) ModulateFreq(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the low cut amplitude and adjust the high amplitude to keep the gap.
//
// In other words, the difference between low and high cut points will stay the same.
func (n Clip) ModulateBoth(low, high float32, m Modulator) {
	m.Modulate(n.id, 0, low, high)
}

// Modulate the low cut amplitude.
func (n Clip) ModulateLow(low, high float32, m Modulator) {
	m.Modulate(n.id, 1, low, high)
}

// Modulate the high cut amplitude.
func (n Clip) ModulateHigh(low, high float32, m Modulator) {
	m.Modulate(n.id, 2, low, high)
}

// Set the gain level.
func (n Gain) Set(val float32) {
	setParam(n.id, 0, val)
}

// Set the pan value (from 0. to 1.: 0. is only left, 1. is only right).
func (n Pan) Set(val float32) {
	setParam(n.id, 0, val)
}

func (n Mute) Mute() {
	setParam(n.id, 0, 0.)
}

func (n Mute) Unmute() {
	setParam(n.id, 0, 1.)
}

func (n Pause) Pause() {
	setParam(n.id, 0, 0.)
}

func (n Pause) Play() {
	setParam(n.id, 0, 1.)
}

// Set the cut-off frequency.
func (n LowPass) SetFreq(freq Freq) {
	setParam(n.id, 0, float32(freq))
}

// Set the cut-off frequency.
func (n HighPass) SetFreq(freq Freq) {
	setParam(n.id, 0, float32(freq))
}

// Set the low cut amplitude and adjust the high amplitude to keep the gap.
//
// In other words, the difference between low and high cut points will stay the same.
func (n Clip) SetBoth(val float32) {
	setParam(n.id, 0, val)
}

// Set the low cut amplitude.
func (n Clip) SetLow(val float32) {
	setParam(n.id, 1, val)
}

// Set the high cut amplitude.
func (n Clip) SetHigh(val float32) {
	setParam(n.id, 2, val)
}
