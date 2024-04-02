package main

import "github.com/life4/firefly-go/firefly"

var point firefly.Point

const radius = 10

func init() {
	firefly.Render = render
	firefly.Update = update
}

func update() {
	screen := firefly.GetScreenSize()
	input := firefly.ReadLeft()
	point = firefly.Point{
		X: int32(screen.W/2) + int32(input.X/20) - radius,
		Y: int32(screen.H/2) - int32(input.Y/20) - radius,
	}
}

func render() {
	firefly.Clear(firefly.ColorLight)
	style := firefly.Style{FillColor: 2, StrokeColor: 3, StrokeWidth: 1}
	firefly.DrawCircle(point, radius*2, style)
}
