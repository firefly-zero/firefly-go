package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var font firefly.Font

func boot() {
	font = firefly.LoadFont("font", 727)
}

func render() {
	firefly.DrawText(
		"The quick brown fox jumps over the lazy dog",
		font, firefly.Point{X: 60, Y: 10}, firefly.ColorAccent,
	)
}
