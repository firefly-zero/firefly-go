// Image source: https://opengameart.org/content/zelda-like-tilesets-and-sprites
package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var image firefly.Image

func boot() {
	image = firefly.LoadImage("img", 94)
}

func render() {
	firefly.DrawImage(
		image, firefly.Point{X: 60, Y: 60},
		firefly.ColorDark, firefly.ColorAccent, firefly.ColorSecondary, firefly.ColorNone,
	)
}
