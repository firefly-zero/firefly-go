package audio

// Modulator can be attached to a node to change a node parameter over time.
//
// Modulators include both LFOs (Low-Frequency Oscillator) and envelopes.
// The difference is that LFOs keep oscillating between values
// while envelopes go from one value to another and then stop.
//
// Internally, modulators only produce values from 0 to 1.
// Then, to get the final value of the parameter,
// the value from the modulator is projected on the range
// between `lowe` and `high` arguments passed together
// with the modulator when attaching a modulator to a node.
// For example, [Sine.Modulate] accepts the range of modulated values
// for the sine wave frequency (which can be used for vibrato effect).
//
// Even if a node has multiple parameters that can be modulated,
// currently  single node may have at most one modulator attached.
type Modulator interface {
	Modulate(nodeID uint32, param uint32, low, high float32)
}

// Linear (ramp up or down) envelope.
//
// It looks like this: `⎽╱⎺` (or `⎺╲⎽` if `low` is bigger than `high`).
//
// The value before `StartAt` is 0, the value after `EndAt` is 1,
// and the value between `StartAt` and `EndAt` changes linearly from 0 to 1.
//
// Most often used with [Gain] for fade in and fade out effect.
type LinearModulator struct {
	StartAt Samples
	EndAt   Samples
}

var _ Modulator = LinearModulator{}

// Modulate implements [Modulator].
func (m LinearModulator) Modulate(nodeID uint32, param uint32, low, high float32) {
	modLinear(
		nodeID,
		param,
		low,
		high,
		uint32(m.StartAt),
		uint32(m.EndAt),
	)
}

// Hold envelope.
//
// It looks like this: `⎽│⎺` (or `⎺│⎽` if `low` is bigger than `high`).
//
// The value before `Time` is 0 and the value after `Time` is 1.
// Equivalent to [LinearModulator] with `StartAt` being equal to `EndAt`.
type HoldModulator struct {
	Time Samples
}

var _ Modulator = HoldModulator{}

// Modulate implements [Modulator].
func (m HoldModulator) Modulate(nodeID uint32, param uint32, low, high float32) {
	modHold(nodeID, param, low, high, uint32(m.Time))
}

// ADSR envelope.
//
// It looks like this: `🭋🭍🬹🬿`
//
//  1. Until `Attack`, the value goes from 0 to 1;
//  2. Until `Decay`, it goes from 1 to `SustainLevel`;
//  3. Until `Sustain`, it holds `SustainLevel`;
//  4. Until `Release`, it goes from `SustainLevel` to 0;
//  5. After `Release`, it holds 0.
//
// Most commonly used with [Gain].
type ADSRModulator struct {
	// When the value reaches 1.
	Attack Samples
	// When the value reaches `SustainLevel`.
	Decay Samples
	// Until when the value holds `SustainLevel`.
	Sustain Samples
	// The value generated from `Decay` until `Sustain`.
	SustainLevel float32
	// When the value drops to 0.
	Release Samples
}

var _ Modulator = ADSRModulator{}

// Modulate implements [Modulator].
func (m ADSRModulator) Modulate(nodeID uint32, param uint32, low, high float32) {
	modAdsr(
		nodeID, param, low, high,
		uint32(m.Attack),
		uint32(m.Decay),
		uint32(m.Sustain), m.SustainLevel,
		uint32(m.Release),
	)
}

// Sine wave low-frequency oscillator.
//
// It looks like this: `∿`.
//
// Most commonly used with [Sine] (or another wave generator)
// to produce vibrato effect.
type SineModulator struct {
	Freq Hz
}

var _ Modulator = SineModulator{}

// Modulate implements [Modulator].
func (m SineModulator) Modulate(nodeID uint32, param uint32, low, high float32) {
	modSine(nodeID, param, float32(m.Freq), low, high)
}

// Square wave low-frequency oscillator.
//
// It looks like this: `🭿🭾🭿🭾🭿🭾🭿🭾`
type SquareModulator struct {
	Period Samples
}

var _ Modulator = SquareModulator{}

// Modulate implements [Modulator].
func (m SquareModulator) Modulate(nodeID uint32, param uint32, low, high float32) {
	modSquare(nodeID, param, low, high, uint32(m.Period))
}

// Sawtooth wave low-frequency oscillator.
//
// It looks like this: `╱│╱│╱│╱│`
type SawtoothModulator struct {
	Period Samples
}

var _ Modulator = SawtoothModulator{}

// Modulate implements [Modulator].
func (m SawtoothModulator) Modulate(nodeID uint32, param uint32, low, high float32) {
	modSawtooth(nodeID, param, low, high, uint32(m.Period))
}
