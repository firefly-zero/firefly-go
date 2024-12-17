package firefly

var (
	// Callback to be called once when the app starts.
	//
	// Called after memory is initialized and host functions are registered
	// but before any other callback.
	Boot func()

	// Callback to be called on every update.
	//
	// Don't use it to draw on the screen, use [Render] instead.
	Update func()

	// Callback to be called before rendering the frame.
	//
	// Don't use it to update the state, use [Update] instead.
	Render func()

	// Callback to be called before exiting the app.
	BeforeExit func()

	// Callback to be called before rendering a horizontal line on the screen.
	//
	// Accepts the index of the line about to be rendered
	// and returns the index of the line for which it should be called next time.
	// Use it to update color palette to support more than 4 colors per frame.
	RenderLine func(int) int

	// Callback to be called when a cheat code is sent from firefly CLI.
	//
	// Accepts the command index and value and returns a response to show in CLI.
	Cheat func(int, int) int
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

//go:export before_exit
func beforeExit() { //nolint
	if BeforeExit != nil {
		BeforeExit()
	}
}

//go:export render_line
func renderLine(l int32) int32 { //nolint
	if RenderLine != nil {
		return int32(RenderLine(int(l)))
	}
	return 0
}

//go:export cheat
func cheat(c, v int32) int32 { //nolint
	if RenderLine != nil {
		return int32(Cheat(int(c), int(v)))
	}
	return 0
}
