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
		"The quick brown fox\njumps over the lazy dog",
		font, firefly.Point{X: 10, Y: 60}, firefly.ColorDarkBlue,
	)
}
