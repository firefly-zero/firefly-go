package main

import "github.com/life4/firefly-go/firefly"

var pos = firefly.Point{X: 60, Y: 40}
var dir int = 2

const diameter = 20

func init() {
	firefly.Update = update
	firefly.Render = render
}

func update() {
	size := firefly.GetScreenSize()
	if pos.X <= 0 || pos.X+diameter >= size.W {
		dir = -dir
	}
	pos.X += dir
}

func render() {
	firefly.Clear(firefly.ColorLight)
	style := firefly.Style{
		FillColor:   firefly.ColorAccent,
		StrokeColor: firefly.ColorSecondary,
		StrokeWidth: 1,
	}
	firefly.DrawCircle(pos, diameter, style)
}
