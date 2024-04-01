package main

import "github.com/life4/firefly-go/firefly"

func init() {
	firefly.Render = render
}

func render() {
	firefly.DrawTriangle(
		firefly.Point{X: 60, Y: 10},
		firefly.Point{X: 40, Y: 40},
		firefly.Point{X: 80, Y: 40},
		firefly.Style{FillColor: 2, StrokeColor: 3, StrokeWidth: 1},
	)
}
