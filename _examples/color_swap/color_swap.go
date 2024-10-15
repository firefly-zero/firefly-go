package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

var (
	image firefly.Image
	frame int
	color firefly.Color
)

func boot() {
	image = firefly.LoadROMFile("img").Image()
}

func update() {
	frame = (frame + 1) % 30
	if frame == 0 {
		image.SetColor(2, color)
		rotateColor()
	}
}

func rotateColor() {
	switch color { //nolint:exhaustive
	case firefly.ColorCyan:
		color = firefly.ColorRed
	case firefly.ColorRed:
		color = firefly.ColorOrange
	case firefly.ColorOrange:
		color = firefly.ColorGreen
	default:
		color = firefly.ColorCyan
	}
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawImage(image, firefly.Point{X: 60, Y: 60})
}
