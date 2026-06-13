package audio

import "time"

const SampleRate = 44_100

type Time uint32

// Time in number of samples.
func Samples(s uint32) Time {
	return Time(s)
}

func Seconds(s uint32) Time {
	return Time(s * SampleRate)
}

// Time in milliseconds.
func MS(s uint32) Time {
	return Time(s * SampleRate / 1000)
}

// Time from [time.Duration].
func Duration(t time.Duration) Time {
	s := t.Nanoseconds() * SampleRate / 1e9
	return Time(uint32(s))
}

func (t Time) Seconds() uint32 {
	return uint32(t) / SampleRate
}

func (t Time) MS() uint32 {
	return uint32(t) * 1000 / SampleRate
}

func (t Time) Frequency() Freq {
	return SampleRate / Freq(t)
}
