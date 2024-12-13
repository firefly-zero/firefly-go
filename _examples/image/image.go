package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var image firefly.Image

func boot() {
	image = firefly.LoadFile("img", nil).Image()
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawImage(image, firefly.Point{X: 60, Y: 60})
}
