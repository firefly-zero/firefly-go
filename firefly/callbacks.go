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

	// Callback to be called when a cheat code is sent from firefly CLI.
	//
	// Accepts the command index and value and returns a response to show in CLI.
	Cheat func(int, int) int
)

//go:export boot
func boot() {
	if Boot != nil {
		Boot()
	}
}

//go:export update
func update() {
	if Update != nil {
		Update()
	}
}

//go:export render
func render() {
	if Render != nil {
		Render()
	}
}

//go:export before_exit
func beforeExit() {
	if BeforeExit != nil {
		BeforeExit()
	}
}

//go:export cheat
func cheat(c, v int32) int32 {
	if Cheat != nil {
		return int32(Cheat(int(c), int(v)))
	}
	return 0
}

var (
	_ = boot
	_ = update
	_ = render
	_ = beforeExit
	_ = cheat
)
