package audio

type SourceNode struct {
	id uint32
}

// An audio source node created by [Node.AddSine].
type Sine struct{ SourceNode }

// An audio source node created by [Node.AddSquare].
type Square struct{ SourceNode }

// An audio source node created by [Node.AddSawtooth].
type Sawtooth struct{ SourceNode }

// An audio source node created by [Node.AddTriangle].
type Triangle struct{ SourceNode }

// An audio source node created by [Node.AddFile].
type File struct{ SourceNode }

// An audio source node created by [Node.AddNoise].
type Noise struct{ SourceNode }

// An audio source node created by [Node.AddEmpty].
type Empty struct{ SourceNode }

// An audio source node created by [Node.AddZero].
type Zero struct{ SourceNode }

// Reset the node state to how it was when it was just added.
func (n SourceNode) Reset() {
	reset(n.id)
}

// Reset the node and all child nodes to the state to how it was when they were just added.
func (n SourceNode) ResetAll() {
	resetAll(n.id)
}

// Remove all child nodes.
//
// After it is called, you should make sure to discard all references to the old
// child nodes.
func (n SourceNode) Clear() {
	clearNode(n.id)
}

// Modulate oscillation frequency.
func (n Sine) Modulate(low, high Hz, m Modulator) {
	m.Modulate(n.id, 0, float32(low), float32(high))
}

// Modulate oscillation frequency.
func (n Square) Modulate(low, high Hz, m Modulator) {
	m.Modulate(n.id, 0, float32(low), float32(high))
}

// Modulate oscillation frequency.
func (n Sawtooth) Modulate(low, high Hz, m Modulator) {
	m.Modulate(n.id, 0, float32(low), float32(high))
}

// Modulate oscillation frequency.
func (n Triangle) Modulate(low, high Hz, m Modulator) {
	m.Modulate(n.id, 0, float32(low), float32(high))
}

// Set oscillation frequency.
func (n Sine) Set(freq Hz) {
	setParam(n.id, 0, float32(freq))
}

// Set oscillation frequency.
func (n Square) Set(freq Hz) {
	setParam(n.id, 0, float32(freq))
}

// Set oscillation frequency.
func (n Sawtooth) Set(freq Hz) {
	setParam(n.id, 0, float32(freq))
}

// Set oscillation frequency.
func (n Triangle) Set(freq Hz) {
	setParam(n.id, 0, float32(freq))
}

// Go to the specified timestamp in the file.
func (n File) Seek(t Samples) {
	setParam(n.id, 0, float32(t))
}
