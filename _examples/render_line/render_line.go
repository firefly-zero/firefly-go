package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Render = render
	firefly.RenderLine = renderLine
}

func render() {
	firefly.Clear(firefly.ColorLight)
	firefly.DrawCircle(
		firefly.Point{X: 10, Y: 10},
		120,
		firefly.Style{
			FillColor:   firefly.ColorAccent,
			StrokeColor: firefly.ColorDark,
			StrokeWidth: 1,
		},
	)
}

func renderLine(l int) int {
	if l >= 100 {
		firefly.SetColor(firefly.ColorAccent, firefly.RGB{R: 0xFF, G: 0, B: 0})
		return 0
	} else if l >= 50 {
		firefly.SetColor(firefly.ColorAccent, firefly.RGB{R: 0, G: 0, B: 0xFF})
		return 100
	} else {
		firefly.SetColor(firefly.ColorAccent, firefly.RGB{R: 0, G: 0xFF, B: 0})
		return 50
	}
}
