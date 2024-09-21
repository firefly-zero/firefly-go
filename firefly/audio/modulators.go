package audio

type Modulator interface {
	Modulate(nodeID uint32, param uint32)
}

// Linear (ramp up or down) envelope.
//
// It looks like this: `⎽╱⎺` or `⎺╲⎽`.
//
// The value before `start_at` is `start`, the value after `end_at` is `end`,
// and the value between `start_at` and `end_at` changes linearly from `start` to `end`.
type LinearModulator struct {
	Start   float32
	End     float32
	StartAt Samples
	EndAt   Samples
}

var _ Modulator = LinearModulator{}

func (m LinearModulator) Modulate(nodeID uint32, param uint32) {
	modLinear(
		nodeID,
		param,
		m.Start,
		m.End,
		uint32(m.StartAt),
		uint32(m.EndAt),
	)
}

// Hold envelope.
//
// It looks like this: `⎽│⎺` or `⎺│⎽`.
//
// The value before `time` is `before` and the value after `time` is `after`.
// Equivalent to [`LinearModulator`] with `start_at` being equal to `end_at`.
type HoldModulator struct {
	Before float32
	After  float32
	Time   Samples
}

var _ Modulator = HoldModulator{}

func (m HoldModulator) Modulate(nodeID uint32, param uint32) {
	modHold(nodeID, param, m.Before, m.After, uint32(m.Time))
}

// Sine wave low-frequency oscillator.
//
// It looks like this: `∿`.
//
// `low` is the lowest produced value, `high` is the highest.
type SineModulator struct {
	Freq Hz
	Low  float32
	High float32
}

var _ Modulator = SineModulator{}

func (m SineModulator) Modulate(nodeID uint32, param uint32) {
	modSine(nodeID, param, float32(m.Freq), m.Low, m.High)
}
