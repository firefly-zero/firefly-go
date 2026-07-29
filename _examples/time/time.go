package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var font firefly.Font

func boot() {
	font = firefly.LoadFile("font", nil).Font()
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawText(
		firefly.GetTime().String(),
		font, firefly.Point{X: 10, Y: 60}, firefly.ColorDarkBlue,
	)
}
