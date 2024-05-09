package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Render = render
	firefly.RenderLine = renderLine
}

func render() {
	firefly.ClearScreen(firefly.ColorLight)
	firefly.DrawCircle(
		firefly.Point{X: 50, Y: 10},
		120,
		firefly.Style{
			FillColor:   firefly.ColorAccent,
			StrokeColor: firefly.ColorDark,
			StrokeWidth: 1,
		},
	)
}

func renderLine(l int) int {
	firefly.SetColor(
		firefly.ColorAccent,
		firefly.RGB{R: uint8(255 - l*2), G: 0, B: 0},
	)
	return l + 5
}
