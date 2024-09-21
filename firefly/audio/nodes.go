package audio

type Node struct {
	id uint32
}

var Out = Node{0}

// A marker for a specific node type. See [Node.AddSine].
type Sine struct{ Node }

// A marker for a specific node type. See [Node.AddMix].
type Mix struct{ Node }

// A marker for a specific node type. See [Node.AddAllForOne].
type AllForOne struct{ Node }

// A marker for a specific node type. See [Node.AddGain].
type Gain struct{ Node }

// A marker for a specific node type. See [Node.AddLoop].
type Loop struct{ Node }

// A marker for a specific node type. See [Node.AddConcat].
type Concat struct{ Node }

// A marker for a specific node type. See [Node.AddPan].
type Pan struct{ Node }

// A marker for a specific node type. See [Node.AddMute].
type Mute struct{ Node }

// A marker for a specific node type. See [Node.AddPause].
type Pause struct{ Node }

// A marker for a specific node type. See [Node.AddTrackPosition].
type TrackPosition struct{ Node }

// A marker for a specific node type. See [Node.AddLowPass].
type LowPass struct{ Node }

// A marker for a specific node type. See [Node.AddHighPass].
type HighPass struct{ Node }

// A marker for a specific node type. See [Node.AddTakeLeft].
type TakeLeft struct{ Node }

// A marker for a specific node type. See [Node.AddTakeRight].
type TakeRight struct{ Node }

// A marker for a specific node type. See [Node.AddSwap].
type Swap struct{ Node }

// A marker for a specific node type. See [Node.AddClip].
type Clip struct{ Node }

// A marker for a specific node type. See [Node.AddSquare].
type Square struct{ Node }

// A marker for a specific node type. See [Node.AddSawtooth].
type Sawtooth struct{ Node }

// A marker for a specific node type. See [Node.AddTriangle].
type Triangle struct{ Node }

// A marker for a specific node type. See [Node.AddNoise].
type Noise struct{ Node }

// A marker for a specific node type. See [Node.AddEmpty].
type Empty struct{ Node }

// A marker for a specific node type. See [Node.AddZero].
type Zero struct{ Node }

// Add sine wave oscillator source (`∿`).
func (n Node) AddSine(f Freq, phase float32) Sine {
	id := addSine(n.id, float32(f), phase)
	return Sine{Node{id}}
}

// Add square wave oscillator source (`⎍`).
func (n Node) AddSquare(f Freq, phase float32) Square {
	id := addSquare(n.id, float32(f), phase)
	return Square{Node{id}}
}

// Add sawtooth wave oscillator source (`╱│`).
func (n Node) AddSawtooth(f Freq, phase float32) Sawtooth {
	id := addSawtooth(n.id, float32(f), phase)
	return Sawtooth{Node{id}}
}

// Add triangle wave oscillator source (`╱╲`).
func (n Node) AddTriangle(f Freq, phase float32) Triangle {
	id := addTriangle(n.id, float32(f), phase)
	return Triangle{Node{id}}
}

// Add white noise source (amplitude on each tick is random).
func (n Node) AddNoise(seed int) Noise {
	id := addNoise(n.id, int32(seed))
	return Noise{Node{id}}
}

// Add always stopped source.
func (n Node) AddEmpty() Empty {
	id := addEmpty(n.id)
	return Empty{Node{id}}
}

// Add silent source producing zeros.
func (n Node) AddZero() Zero {
	id := addZero(n.id)
	return Zero{Node{id}}
}

// Add node simply mixing all inputs.
func (n Node) AddMix() Mix {
	id := addMix(n.id)
	return Mix{Node{id}}
}

// Add mixer node that stops if any of the sources stops.
func (n Node) AddAllForOne() AllForOne {
	id := addAllForOne(n.id)
	return AllForOne{Node{id}}
}

// Add gain control node.
func (n Node) AddGain(lvl float32) Gain {
	id := addGain(n.id, lvl)
	return Gain{Node{id}}
}

// Add a loop node that resets the input if it stops.
func (n Node) AddLoop() Loop {
	id := addLoop(n.id)
	return Loop{Node{id}}
}

// Add a node that plays the inputs one after the other, in the order as they added.
func (n Node) AddConcat() Concat {
	id := addConcat(n.id)
	return Concat{Node{id}}
}

// Add node panning the audio to the left (0.), right (1.), or something in between.
func (n Node) AddPan(lvl float32) Pan {
	id := addPan(n.id, lvl)
	return Pan{Node{id}}
}

// Add node that can be muted using modulation.
func (n Node) AddMute() Mute {
	id := addMute(n.id)
	return Mute{Node{id}}
}

// Add node that can be paused using modulation.
func (n Node) AddPause() Pause {
	id := addPause(n.id)
	return Pause{Node{id}}
}

// Add node tracking the elapsed playback time.
func (n Node) AddTrackPosition() TrackPosition {
	id := addTrackPosition(n.id)
	return TrackPosition{Node{id}}
}

// Add lowpass filter node.
func (n Node) AddLowPass(freq float32, q float32) LowPass {
	id := addLowPass(n.id, freq, q)
	return LowPass{Node{id}}
}

// Add highpass filter node.
func (n Node) AddHighPass(freq float32, q float32) HighPass {
	id := addHighPass(n.id, freq, q)
	return HighPass{Node{id}}
}

// Add node converting stereo to mono by taking the left channel.
func (n Node) AddTakeLeft() TakeLeft {
	id := addTakeLeft(n.id)
	return TakeLeft{Node{id}}
}

// Add node converting stereo to mono by taking the right channel.
func (n Node) AddTakeRight() TakeRight {
	id := addTakeRight(n.id)
	return TakeRight{Node{id}}
}

// Add node swapping left and right channels of the stereo input.
func (n Node) AddSwap() Swap {
	id := addSwap(n.id)
	return Swap{Node{id}}
}

// Add node clamping the input amplitude. Can be used for hard distortion.
func (n Node) AddClip(low float32, high float32) Clip {
	id := addClip(n.id, low, high)
	return Clip{Node{id}}
}

// Reset the node state to how it was when it was just added.
func (n Node) Reset() {
	reset(n.id)
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
	clear(n.id)
}
