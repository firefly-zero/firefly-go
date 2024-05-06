package firefly

var (
	Boot       func()
	Update     func()
	Render     func()
	RenderLine func(int) int
)

//go:export boot
func boot() { //nolint
	if Boot != nil {
		Boot()
	}
}

//go:export update
func update() { //nolint
	if Update != nil {
		Update()
	}
}

//go:export render
func render() { //nolint
	if Render != nil {
		Render()
	}
}

//go:export render_line
func renderLine(l int32) int32 { //nolint
	if RenderLine != nil {
		return int32(RenderLine(int(l)))
	}
	return 0
}
