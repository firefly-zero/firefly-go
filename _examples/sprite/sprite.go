// Image source: https://opengameart.org/content/zelda-like-tilesets-and-sprites
package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

var sprite firefly.Image

// How many animation frames the sprite has.
const frames = 12

var frame = 0

func boot() {
	sprite = firefly.LoadImage("sprite")
}

func update() {
	frame = (frame + 1) % frames
}

func render() {
	firefly.Clear(firefly.ColorLight)
	sub := sprite.Sub(
		firefly.Point{X: 32 * frame, Y: 0},
		firefly.Size{W: 32, H: 32},
	)
	firefly.DrawSubImage(
		sub, firefly.Point{X: 60, Y: 60},
		firefly.ColorDark, firefly.ColorSecondary, firefly.ColorNone, firefly.ColorNone,
	)
}
