package main

import "github.com/life4/firefly-go/firefly"

var point firefly.Point
var center firefly.Point

const radius = 10

func init() {
	firefly.Render = render
	firefly.Update = update
}

func update() {
	screen := firefly.GetScreenSize()
	center = firefly.Point{
		X: screen.W / 2,
		Y: screen.H / 2,
	}
	input := firefly.ReadPad()
	point = firefly.Point{
		X: center.X + input.X/20 - radius,
		Y: center.Y - input.Y/20 - radius,
	}
}

func render() {
	firefly.Clear(firefly.ColorLight)
	style := firefly.Style{
		FillColor:   firefly.ColorLight,
		StrokeColor: firefly.ColorDark,
		StrokeWidth: 1,
	}
	firefly.DrawCircle(firefly.Point{
		X: center.X - 50 - radius,
		Y: center.Y - 50 - radius,
	}, 100+radius*2, style)

	style = firefly.Style{
		FillColor:   firefly.ColorAccent,
		StrokeColor: firefly.ColorSecondary,
		StrokeWidth: 1,
	}
	firefly.DrawCircle(point, radius*2, style)
}
