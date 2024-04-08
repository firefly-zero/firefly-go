// Image source: https://opengameart.org/content/zelda-like-tilesets-and-sprites
package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

var sprite firefly.Image

const frames = 12

var frame = 0

func boot() {
	sprite = firefly.LoadImage("sprite", 3076)
}

func update() {
	frame = (frame + 1) % 12
}

func render() {
	firefly.Clear(firefly.ColorLight)
	firefly.DrawSubImage(
		sprite,
		firefly.Point{X: 60, Y: 60},
		firefly.Point{X: 32 * frame, Y: 0},
		firefly.Size{W: 32, H: 32},
		firefly.ColorDark, firefly.ColorSecondary, firefly.ColorNone, firefly.ColorNone,
	)
}
