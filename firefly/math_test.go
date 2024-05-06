package firefly

import "testing"

func eq[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a != b {
		t.Fatalf("failed check: %v == %v", a, b)
	}
}

func TestSqrt(t *testing.T) {
	// Numbers are off but that's fine,
	// we can afford imprecise results for how we use it.
	eq(t, sqrt(4.), 2.)
	eq(t, sqrt(9.), 3.125)
	eq(t, sqrt(12.25), 3.53125)
	eq(t, sqrt(144.), 12.5)
}
