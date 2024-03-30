package main

import "github.com/life4/firefly-go/firefly"

var pos = firefly.Point{X: 60, Y: 40}
var dir int32 = 2

const diameter = 20

func init() {
	firefly.Update = update
	firefly.Render = render
}

func update() {
	size := firefly.GetScreenSize()
	if pos.X <= 0 || pos.X+diameter >= int32(size.W) {
		dir = -dir
	}
	pos.X += dir
}

func render() {
	firefly.Clear(3)
	style := firefly.Style{FillColor: 1, StrokeColor: 2, StrokeWidth: 1}
	firefly.DrawCircle(pos, diameter, style)
}
