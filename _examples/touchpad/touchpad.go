package main

import "github.com/firefly-zero/firefly-go/firefly"

var (
	point  *firefly.Point
	center firefly.Point
)

const radius = 10

func init() {
	firefly.Render = render
	firefly.Update = update
}

func update() {
	center = firefly.Point{
		X: firefly.Width / 2,
		Y: firefly.Height / 2,
	}
	input, pressed := firefly.ReadPad(firefly.Player0)
	if pressed {
		point = &firefly.Point{
			X: center.X + input.X/20 - radius,
			Y: center.Y - input.Y/20 - radius,
		}
	} else {
		point = nil
	}
}

func render() {
	firefly.ClearScreen(firefly.ColorLight)
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
	if point != nil {
		firefly.DrawCircle(*point, radius*2, style)
	}
}
