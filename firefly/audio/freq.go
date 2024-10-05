package audio

// Frequency in Hz.
type Hz float32

const (
	// C0, MIDI note #12.
	C0  = Hz(16.351)
	CS0 = Hz(17.324)
	D0  = Hz(18.354)
	DS0 = Hz(19.445)
	E0  = Hz(20.601)
	F0  = Hz(21.827)
	FS0 = Hz(23.124)
	G0  = Hz(24.499)
	GS0 = Hz(25.956)
	// A0, the lowest note of a piano.
	A0  = Hz(27.5)
	AS0 = Hz(29.135)
	// B0, the lowest note of a 5 string bass.
	B0 = Hz(30.868)
	// C1, the lowest note of double bass with C extension.
	C1  = Hz(32.703)
	CS1 = Hz(34.648)
	D1  = Hz(36.708)
	DS1 = Hz(38.891)
	// E1, the lowest note of a bass.
	E1  = Hz(41.203)
	F1  = Hz(43.654)
	FS1 = Hz(46.249)
	G1  = Hz(48.999)
	GS1 = Hz(51.913)
	A1  = Hz(55.)
	AS1 = Hz(58.27)
	B1  = Hz(61.735)
	C2  = Hz(65.406)
	CS2 = Hz(69.296)
	D2  = Hz(73.416)
	DS2 = Hz(77.782)
	// E2, the lowest note of a guitar.
	E2  = Hz(82.407)
	F2  = Hz(87.307)
	FS2 = Hz(92.499)
	G2  = Hz(97.999)
	GS2 = Hz(103.826)
	A2  = Hz(110.)
	AS2 = Hz(116.541)
	B2  = Hz(123.471)
	C3  = Hz(130.813)
	CS3 = Hz(138.591)
	D3  = Hz(146.832)
	DS3 = Hz(155.563)
	E3  = Hz(164.814)
	F3  = Hz(174.614)
	FS3 = Hz(184.997)
	// G3, the lowest note of a violin.
	G3  = Hz(195.998)
	GS3 = Hz(207.652)
	A3  = Hz(220.)
	AS3 = Hz(233.082)
	B3  = Hz(246.942)
	// C4, the "middle C".
	C4  = Hz(261.626)
	CS4 = Hz(277.183)
	D4  = Hz(293.665)
	DS4 = Hz(311.127)
	E4  = Hz(329.628)
	F4  = Hz(349.228)
	FS4 = Hz(369.994)
	G4  = Hz(391.995)
	GS4 = Hz(415.305)
	// A4, the tuning reference note.
	A4  = Hz(440.)
	AS4 = Hz(466.164)
	B4  = Hz(493.883)
	C5  = Hz(523.251)
	CS5 = Hz(554.365)
	D5  = Hz(587.33)
	DS5 = Hz(622.254)
	E5  = Hz(659.255)
	F5  = Hz(698.456)
	FS5 = Hz(739.989)
	G5  = Hz(783.991)
	GS5 = Hz(830.609)
	A5  = Hz(880.)
	AS5 = Hz(932.328)
	B5  = Hz(987.767)
	C6  = Hz(1046.502)
	CS6 = Hz(1108.731)
	D6  = Hz(1174.659)
	DS6 = Hz(1244.508)
	E6  = Hz(1318.51)
	F6  = Hz(1396.913)
	FS6 = Hz(1479.978)
	G6  = Hz(1567.982)
	GS6 = Hz(1661.219)
	A6  = Hz(1760.)
	AS6 = Hz(1864.655)
	B6  = Hz(1975.533)
	C7  = Hz(2093.005)
	CS7 = Hz(2217.461)
	D7  = Hz(2349.318)
	DS7 = Hz(2489.016)
	E7  = Hz(2637.021)
	F7  = Hz(2793.826)
	FS7 = Hz(2959.955)
	G7  = Hz(3135.964)
	GS7 = Hz(3322.438)
	A7  = Hz(3520.)
	AS7 = Hz(3729.31)
	B7  = Hz(3951.066)
	// C8, the highest note of a piano.
	C8  = Hz(4186.009)
	CS8 = Hz(4434.922)
	D8  = Hz(4698.636)
	DS8 = Hz(4978.032)
	E8  = Hz(5274.042)
	F8  = Hz(5587.652)
	FS8 = Hz(5919.91)
	G8  = Hz(6271.928)
	GS8 = Hz(6644.876)
	A8  = Hz(7040.)
	AS8 = Hz(7458.62)
	B8  = Hz(7902.132)
	C9  = Hz(8372.018)
	CS9 = Hz(8869.844)
	D9  = Hz(9397.272)
	DS9 = Hz(9956.064)
	E9  = Hz(10548.084)
	F9  = Hz(11175.304)
	FS9 = Hz(11839.82)
	G9  = Hz(12543.856)
	// G#9, MIDI note #128, the top of the MIDI tuning range.
	GS9 = Hz(13289.752)
	A9  = Hz(14080.)
	AS9 = Hz(14917.24)
	// B9. For most of adults, it is already beyond the hearing range.
	B9 = Hz(15804.264)
)

func MIDI(note uint8) Hz { //nolint:cyclop
	// https://inspiredacoustics.com/en/MIDI_note_numbers_and_center_frequencies
	// https://en.wikipedia.org/wiki/Musical_note#MIDI
	var f float32
	switch note % 12 {
	case 0:
		f = 8.1758
	case 1:
		f = 8.66
	case 2:
		f = 9.18
	case 3:
		f = 9.72
	case 4:
		f = 10.30
	case 5:
		f = 10.91
	case 6:
		f = 11.56
	case 7:
		f = 12.25
	case 8:
		f = 12.98
	case 9:
		f = 13.75
	case 10:
		f = 14.57
	default:
		f = 15.43
	}
	oct := note / 12
	f *= float32(uint32(1) << oct)
	return Hz(f)
}
