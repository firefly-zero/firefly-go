package main

import "github.com/firefly-zero/firefly-go/firefly"

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
	sprite = firefly.LoadFile("sprite").Image()
}

func update() {
	frame = (frame + 1) % frames
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	sub := sprite.Sub(
		firefly.Point{X: 32 * frame, Y: 0},
		firefly.Size{W: 32, H: 32},
	)
	firefly.DrawSubImage(sub, firefly.Point{X: 60, Y: 60})
}
