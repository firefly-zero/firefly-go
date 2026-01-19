// Package random contains a reimplementation of the [math/rand] package
// that uses [firefly]'s [firefly.GetRandom]
package random

import (
	"math"
	"math/rand"

	"github.com/firefly-zero/firefly-go/firefly"
)

var globalRand = Rand{}

// Rand provides a reimplementation of [math/rand.Rand] that uses Firefly Zero's
// random number generator via [firefly.GetRandom].
type Rand struct{}

// ensure it implements the interface
var _ rand.Source = Rand{}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
// Implements [rand.Source].
func (r Rand) Int63() int64 {
	hi := firefly.GetRandom() & math.MaxInt32
	lo := firefly.GetRandom()
	return (int64(hi) << 32) | int64(lo)
}

// Seed sets the pseudo-random number generator seed.
// The actual seed is truncated down into a uint32 due to Firefly Zero
// using unsigned 32-bit integer for its seed.
//
// Seed implements [rand.Source].
func (r Rand) Seed(seed int64) {
	firefly.SetSeed(uint32(seed))
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r Rand) Uint32() uint32 { return firefly.GetRandom() }

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (r Rand) Uint64() uint64 {
	hi := firefly.GetRandom()
	lo := firefly.GetRandom()
	return (uint64(hi) << 32) | uint64(lo)
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r Rand) Int31() int32 { return int32(firefly.GetRandom() & math.MaxInt32) }

// Int returns a non-negative pseudo-random int.
func (r Rand) Int() int { return int(firefly.GetRandom()) }

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}
	return v % n
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r Rand) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func (r Rand) Float64() float64 {
	return float64(r.Int63n(1<<53)) / (1 << 53)
}

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0).
func (r Rand) Float32() float32 {
	return float32(r.Int31n(1<<24)) / (1 << 24)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default [Rand].
func Int63() int64 { return globalRand.Int63() }

// Uint32 returns a pseudo-random 32-bit value as a uint32
// from the default [Rand].
func Uint32() uint32 { return globalRand.Uint32() }

// Uint64 returns a pseudo-random 64-bit value as a uint64
// from the default [Rand].
func Uint64() uint64 { return globalRand.Uint64() }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default [Rand].
func Int31() int32 { return globalRand.Int31() }

// Int returns a non-negative pseudo-random int
// from the default [Rand].
func Int() int { return globalRand.Int() }

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Rand].
// It panics if n <= 0.
func Int63n(n int64) int64 { return globalRand.Int63n(n) }

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Rand].
// It panics if n <= 0.
func Int31n(n int32) int32 { return globalRand.Int31n(n) }

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Rand].
// It panics if n <= 0.
func Intn(n int) int { return globalRand.Intn(n) }

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default [Rand].
func Float64() float64 { return globalRand.Float64() }

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default [Rand].
func Float32() float32 { return globalRand.Float32() }

// Shuffle pseudo-randomizes the order of elements using the default [Rand].
// n is the number of elements.
// It panics if n < 0.
// swap swaps the elements with indexes i and j.
//
// Implements the Fisher-Yates shuffle algorithm.
func Shuffle(n int, swap func(i, j int)) {
	if n <= 0 {
		panic("invalid argument to Shuffle")
	}
	for i := n - 1; i > 0; i = i - 1 {
		j := Intn(i + 1)
		swap(i, j)
	}
}

// ShuffleSlice pseudo-randomizes a generic slice using the default [Rand].
// n is the number of elements.
//
// Implements the Fisher-Yates shuffle algorithm.
func ShuffleSlice[E any](slice []E) {
	for i := len(slice) - 1; i > 0; i = i - 1 {
		j := Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
