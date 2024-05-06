package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var font firefly.Font

func boot() {
	font = firefly.LoadROMFile("font").Font()
}

func render() {
	firefly.DrawText(
		"The quick brown fox\njumps over the lazy dog",
		font, firefly.Point{X: 10, Y: 60}, firefly.ColorAccent,
	)
}
