package firefly

import "unsafe"

type Language uint16

const (
	English   Language = 0x656e // en 🇬🇧 💂
	Dutch     Language = 0x6e6c // nl 🇳🇱 🧀
	French    Language = 0x6672 // fr 🇫🇷 🥐
	German    Language = 0x6465 // de 🇩🇪 🥨
	Italian   Language = 0x6974 // it 🇮🇹 🍕
	Polish    Language = 0x706f // pl 🇵🇱 🥟
	Russian   Language = 0x7275 // ru 🇷🇺 🪆
	Spanish   Language = 0x7370 // sp 🇪🇸 🐂
	Swedish   Language = 0x7365 // se 🇸🇪 ❄️
	Turkish   Language = 0x7470 // tk 🇹🇷 🕌
	Ukrainian Language = 0x746b // ua 🇺🇦 ✊
	TokiPona  Language = 0x7561 // tp 🇨🇦 🙂
)

func (lang Language) Code() string {
	b := [2]uint8{uint8(lang >> 8), uint8(lang)}
	return unsafe.String(&b[0], 2)
}

func (lang Language) NameEnglish() string {
	switch lang {
	case English:
		return "English"
	case Dutch:
		return "Dutch"
	case French:
		return "French"
	case German:
		return "German"
	case Italian:
		return "Italian"
	case Polish:
		return "Polish"
	case Russian:
		return "Russian"
	case Spanish:
		return "Spanish"
	case Swedish:
		return "Swedish"
	case TokiPona:
		return "TokiPona"
	case Turkish:
		return "Turkish"
	case Ukrainian:
		return "Ukrainian"
	}
	return lang.Code()
}

func (lang Language) NameNative() string {
	switch lang {
	case English:
		return "English"
	case Dutch:
		return "Nederlands"
	case French:
		return "Français"
	case German:
		return "Deutsch"
	case Italian:
		return "Italiano"
	case Polish:
		return "Polski"
	case Russian:
		return "Русский"
	case Spanish:
		return "Español"
	case Swedish:
		return "Svenska"
	case TokiPona:
		return "toki pona"
	case Turkish:
		return "Türkçe"
	case Ukrainian:
		return "Українська"
	}
	return lang.Code()
}

func (lang Language) Encoding() string {
	switch lang {
	case English, Dutch, TokiPona:
		return "ascii"
	case Italian, Spanish, Swedish:
		return "iso_8859_1"
	case German, French:
		return "iso_8859_2"
	case Polish:
		return "iso_8859_13"
	case Russian, Ukrainian:
		return "iso_8859_5"
	case Turkish:
		return "iso_8859_9"
	}
	return "ascii"
}

type Theme struct {
	ID uint8
	// The main color of text and boxes.
	Primary Color
	// The color of disable options, muted text, etc.
	Secondary Color
	// The color of important elements, active options, etc.
	Accent Color
	// The background color, the most contrast color to primary.
	BG Color
}

type Settings struct {
	// The preferred color scheme of the player.
	Theme Theme

	// The configured interface language.
	Language Language

	// If true, the screen is rotated 180 degrees.
	//
	// In other words, the player holds the device upside-down.
	// The touchpad is now on the right and the buttons are on the left.
	RotateScreen bool

	// The player has photosensitivity. The app should avoid any rapid flashes.
	ReduceFlashing bool

	// The player wants increased contrast for colors.
	//
	// If set, the black and white colors in the default
	// palette are adjusted automatically. All other colors
	// in the default palette or all colors in a custom palette
	// should be adjusted by the app.
	Contrast bool

	// If true, the player wants to see easter eggs, holiday effects, and weird jokes.
	EasterEggs bool
}

// Log a debug message.
func LogDebug(t string) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	logDebug(ptr, uint32(len(t)))
}

// Log an error message.
func LogError(t string) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	logError(ptr, uint32(len(t)))
}

// Log a debug message from byte slice.
//
// Input slice must be valid UTF-8 formatted (without BOM).
//
// Useful when you are building strings in a byte array or slice, such as via a [bytes.Buffer].
// This function allows you to use such content without allocating a new string.
//
// It is allowed to modify the byte slice after the function call.
func LogDebugBytes(t []byte) {
	ptr := unsafe.Pointer(unsafe.SliceData(t))
	logDebug(ptr, uint32(len(t)))
}

// Log an error message from byte slice.
//
// Input slice must be valid UTF-8 formatted (without BOM).
//
// Useful when you are building strings in a byte array or slice, such as via a [bytes.Buffer].
// This function allows you to use such content without allocating a new string.
//
// It is allowed to modify the byte slice after the function call.
func LogErrorBytes(t []byte) {
	ptr := unsafe.Pointer(unsafe.SliceData(t))
	logError(ptr, uint32(len(t)))
}

// Set the seed used to generate random values.
func SetSeed(seed uint32) {
	setSeed(seed)
}

// Get a random value.
func GetRandom() uint32 {
	return getRandom()
}

// Get human-readable name of the given peer.
func GetName(p Peer) string {
	buf := [16]byte{}
	ptr := unsafe.Pointer(&buf)
	length := getName(uint32(p), ptr)
	return unsafe.String(&buf[0], length)
}

func GetSettings(p Peer) Settings {
	raw := getSettings(uint32(p))
	code := uint16(raw>>8) | uint16(raw)
	language := Language(code)
	flags := raw >> 16
	themeRaw := raw >> 32
	theme := Theme{
		ID:        uint8(themeRaw),
		Primary:   parseColor(themeRaw >> 20),
		Secondary: parseColor(themeRaw >> 16),
		Accent:    parseColor(themeRaw >> 12),
		BG:        parseColor(themeRaw >> 8),
	}
	return Settings{
		Theme:          theme,
		Language:       language,
		RotateScreen:   (flags & 0b0001) != 0,
		ReduceFlashing: (flags & 0b0010) != 0,
		Contrast:       (flags & 0b0100) != 0,
		EasterEggs:     (flags & 0b1000) != 0,
	}
}

//go:inline
func parseColor(c uint64) Color {
	return Color(c&0xf + 1)
}

// Exit the app after the current update is finished.
func Quit() {
	quit()
}

// Restart the app.
func Restart() {
	restart()
}
