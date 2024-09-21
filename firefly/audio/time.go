package audio

import "time"

const SampleRate = 44_100

// Time in number of samples.
type Samples uint32

// Time in seconds.
func Seconds(s uint32) Samples {
	return Samples(s * SampleRate)
}

// Time in milliseconds.
func MS(s uint32) Samples {
	return Samples(s * SampleRate / 1000)
}

// Time from [time.Duration].
func Duration(t time.Duration) Samples {
	s := t.Nanoseconds() * SampleRate / 1e9
	return Samples(uint32(s))
}
