package main

import "github.com/firefly-zero/firefly-go/firefly"

func init() {
	firefly.Render = render
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawQR(
		"https://fireflyzero.com",
		firefly.Point{X: 100, Y: 60},
		firefly.ColorDarkBlue,
		firefly.ColorLightGray,
	)
}
