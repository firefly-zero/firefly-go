package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Render = render
	firefly.RenderLine = renderLine
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawCircle(
		firefly.Point{X: 50, Y: 10},
		120,
		firefly.Style{
			FillColor:   firefly.ColorRed,
			StrokeColor: firefly.ColorBlack,
			StrokeWidth: 1,
		},
	)
}

func renderLine(l int) int {
	firefly.SetColor(
		firefly.ColorRed,
		firefly.RGB{R: uint8(255 - l*2), G: 0, B: 0},
	)
	return l + 5
}
