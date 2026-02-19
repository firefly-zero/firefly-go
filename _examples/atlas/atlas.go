package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var (
	atlas = firefly.NewAtlas(16, 16)

	knight = atlas.Sprite(0, 0)
	wall   = atlas.Sprite(2, 0)
	cat    = atlas.Sprite(3, 0)
	key    = atlas.Sprite(4, 0)
	door   = atlas.Sprite(0, 1)
)

var frame = 0

func boot() {
	atlas.Load("atlas")
}

func render() {
	firefly.ClearScreen(firefly.ColorBlack)

	w := 16
	h := 16

	// top wall
	wall.Draw(firefly.P(2*w, 2*h))
	wall.Draw(firefly.P(3*w, 2*h))
	wall.Draw(firefly.P(4*w, 2*h))
	wall.Draw(firefly.P(5*w, 2*h))
	wall.Draw(firefly.P(6*w, 2*h))

	// bottom wall
	wall.Draw(firefly.P(2*w, 6*h))
	wall.Draw(firefly.P(3*w, 6*h))
	door.Draw(firefly.P(4*w, 6*h))
	wall.Draw(firefly.P(5*w, 6*h))
	wall.Draw(firefly.P(6*w, 6*h))

	// left wall
	wall.Draw(firefly.P(2*w, 2*h))
	wall.Draw(firefly.P(2*w, 3*h))
	wall.Draw(firefly.P(2*w, 4*h))
	wall.Draw(firefly.P(2*w, 5*h))
	wall.Draw(firefly.P(2*w, 6*h))

	// right wall
	wall.Draw(firefly.P(6*w, 2*h))
	wall.Draw(firefly.P(6*w, 3*h))
	wall.Draw(firefly.P(6*w, 4*h))
	wall.Draw(firefly.P(6*w, 5*h))
	wall.Draw(firefly.P(6*w, 6*h))

	// characters and items
	knight.Draw(firefly.P(3*w, 4*h))
	cat.Draw(firefly.P(5*w, 4*h))
	key.Draw(firefly.P(4*w, 5*h))
}
