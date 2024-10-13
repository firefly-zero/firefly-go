package main

import "github.com/firefly-zero/firefly-go/firefly"

var (
	peers     firefly.Peers
	canvas    firefly.Canvas
	positions [4]*firefly.Point
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

func boot() {
	peers = firefly.GetPeers()
	canvas = firefly.NewCanvas(firefly.Size{W: 120, H: 120})
	firefly.SetCanvas(canvas)
	firefly.ClearScreen(firefly.ColorLightGray)
	firefly.UnsetCanvas()
}

func update() {
	for i, peer := range peers.Slice() {
		pad, touched := firefly.ReadPad(peer)
		if !touched {
			positions[i] = nil
			continue
		}
		point := firefly.Point{
			X: pad.X/20 + 50,
			Y: pad.Y/20 + 50,
		}
		positions[i] = &point
		buttons := firefly.ReadButtons(peer)
		if buttons.A {
			firefly.SetCanvas(canvas)
			firefly.DrawPoint(point, firefly.ColorBlue)
			firefly.UnsetCanvas()
		}
	}
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.SetCanvas(canvas)
	shift := firefly.Point{X: 10, Y: 10}
	firefly.DrawCanvas(shift)
	firefly.UnsetCanvas()
	for _, point := range positions {
		if point != nil {
			firefly.DrawPoint(point.Add(shift), firefly.ColorDarkBlue)
		}
	}
}
