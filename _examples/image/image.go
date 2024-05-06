package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Boot = boot
	firefly.Render = render
}

var image firefly.Image

func boot() {
	image = firefly.LoadROMFile("img").Image()
}

func render() {
	colors := firefly.ImageColors{
		A: firefly.ColorDark,
		B: firefly.ColorAccent,
		C: firefly.ColorSecondary,
		D: firefly.ColorNone,
	}
	firefly.DrawImage(image, firefly.Point{X: 60, Y: 60}, colors)
}
