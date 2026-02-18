package firefly_test

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

var testImage = []byte{
	0x22,       // magic number
	0x04, 0x00, // image width
	0x01, // transparency
	// body
	0x01, 0x23, // row 1
	0x45, 0x67, // row 2
	0x89, 0xab, // row 3
	0xcd, 0xef, // row 4
}

func TestImage_GetPixel(t *testing.T) {
	t.Parallel()
	P := firefly.P
	tests := []struct {
		name  string
		raw   []byte
		pixel firefly.Point
		want  firefly.Color
	}{
		{name: "negative point", raw: testImage, pixel: P(-1, -1), want: firefly.ColorNone},
		{name: "point out of bounds", raw: testImage, pixel: P(100, 100), want: firefly.ColorNone},
		{name: "x0y0", raw: testImage, pixel: P(0, 0), want: firefly.ColorBlack},
		{name: "x1y1", raw: testImage, pixel: P(1, 1), want: firefly.ColorLightGreen},
		{name: "x2y3", raw: testImage, pixel: P(2, 3), want: firefly.ColorGray},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			image := firefly.File{test.raw}.Image()
			got := image.GetPixel(test.pixel)
			if got != test.want {
				t.Errorf("pixel: {%d, %d}, want %s, but got %s", test.pixel.X, test.pixel.Y, test.want, got)
			}
		})
	}
}

func TestImagePixels(t *testing.T) {
	t.Parallel()
	image := firefly.File{testImage}.Image()
	got := image.Pixels()
	want := 16
	if got != want {
		t.Errorf("want %d, but got %d", want, got)
	}
}
