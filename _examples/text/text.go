package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var (
	font  firefly.Font
	sixel firefly.Font
)

func boot() {
	font = firefly.LoadFile("font", nil).Font()
	sixel = firefly.LoadFile("sixel", nil).Font()
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawText(
		"The quick brown fox\njumps over the lazy dog",
		font, firefly.Point{X: 10, Y: 60}, firefly.ColorDarkBlue,
	)
	// https://www.vt100.net/dec/vt320/soft_characters
	firefly.DrawText(
		"???owYn||~ywo??\n?IRJaVNn^NVbJRI",
		sixel, firefly.Point{X: 60, Y: 100}, firefly.ColorBlue,
	)
}
