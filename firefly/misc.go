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

// Exit the app after the current update is finished.
func Quit() {
	quit()
}

// Restart the app.
func Restart() {
	restart()
}
