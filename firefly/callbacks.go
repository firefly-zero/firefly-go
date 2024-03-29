package firefly

var (
	Boot   func()
	Update func()
	Render func()
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
