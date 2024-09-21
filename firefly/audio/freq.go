package audio

type Freq float32

const (
	// C0, MIDI note #12.
	C0  = Freq(16.351)
	CS0 = Freq(17.324)
	D0  = Freq(18.354)
	DS0 = Freq(19.445)
	E0  = Freq(20.601)
	F0  = Freq(21.827)
	FS0 = Freq(23.124)
	G0  = Freq(24.499)
	GS0 = Freq(25.956)
	// A0, the lowest note of a piano.
	A0  = Freq(27.5)
	AS0 = Freq(29.135)
	// B0, the lowest note of a 5 string bass.
	B0 = Freq(30.868)
	// C1, the lowest note of double bass with C extension.
	C1  = Freq(32.703)
	CS1 = Freq(34.648)
	D1  = Freq(36.708)
	DS1 = Freq(38.891)
	// E1, the lowest note of a bass.
	E1  = Freq(41.203)
	F1  = Freq(43.654)
	FS1 = Freq(46.249)
	G1  = Freq(48.999)
	GS1 = Freq(51.913)
	A1  = Freq(55.)
	AS1 = Freq(58.27)
	B1  = Freq(61.735)
	C2  = Freq(65.406)
	CS2 = Freq(69.296)
	D2  = Freq(73.416)
	DS2 = Freq(77.782)
	// E2, the lowest note of a guitar.
	E2  = Freq(82.407)
	F2  = Freq(87.307)
	FS2 = Freq(92.499)
	G2  = Freq(97.999)
	GS2 = Freq(103.826)
	A2  = Freq(110.)
	AS2 = Freq(116.541)
	B2  = Freq(123.471)
	C3  = Freq(130.813)
	CS3 = Freq(138.591)
	D3  = Freq(146.832)
	DS3 = Freq(155.563)
	E3  = Freq(164.814)
	F3  = Freq(174.614)
	FS3 = Freq(184.997)
	// G3, the lowest note of a violin.
	G3  = Freq(195.998)
	GS3 = Freq(207.652)
	A3  = Freq(220.)
	AS3 = Freq(233.082)
	B3  = Freq(246.942)
	// C4, the "middle C".
	C4  = Freq(261.626)
	CS4 = Freq(277.183)
	D4  = Freq(293.665)
	DS4 = Freq(311.127)
	E4  = Freq(329.628)
	F4  = Freq(349.228)
	FS4 = Freq(369.994)
	G4  = Freq(391.995)
	GS4 = Freq(415.305)
	// A4, the tuning reference note.
	A4  = Freq(440.)
	AS4 = Freq(466.164)
	B4  = Freq(493.883)
	C5  = Freq(523.251)
	CS5 = Freq(554.365)
	D5  = Freq(587.33)
	DS5 = Freq(622.254)
	E5  = Freq(659.255)
	F5  = Freq(698.456)
	FS5 = Freq(739.989)
	G5  = Freq(783.991)
	GS5 = Freq(830.609)
	A5  = Freq(880.)
	AS5 = Freq(932.328)
	B5  = Freq(987.767)
	C6  = Freq(1046.502)
	CS6 = Freq(1108.731)
	D6  = Freq(1174.659)
	DS6 = Freq(1244.508)
	E6  = Freq(1318.51)
	F6  = Freq(1396.913)
	FS6 = Freq(1479.978)
	G6  = Freq(1567.982)
	GS6 = Freq(1661.219)
	A6  = Freq(1760.)
	AS6 = Freq(1864.655)
	B6  = Freq(1975.533)
	C7  = Freq(2093.005)
	CS7 = Freq(2217.461)
	D7  = Freq(2349.318)
	DS7 = Freq(2489.016)
	E7  = Freq(2637.021)
	F7  = Freq(2793.826)
	FS7 = Freq(2959.955)
	G7  = Freq(3135.964)
	GS7 = Freq(3322.438)
	A7  = Freq(3520.)
	AS7 = Freq(3729.31)
	B7  = Freq(3951.066)
	// C8, the highest note of a piano.
	C8  = Freq(4186.009)
	CS8 = Freq(4434.922)
	D8  = Freq(4698.636)
	DS8 = Freq(4978.032)
	E8  = Freq(5274.042)
	F8  = Freq(5587.652)
	FS8 = Freq(5919.91)
	G8  = Freq(6271.928)
	GS8 = Freq(6644.876)
	A8  = Freq(7040.)
	AS8 = Freq(7458.62)
	B8  = Freq(7902.132)
	C9  = Freq(8372.018)
	CS9 = Freq(8869.844)
	D9  = Freq(9397.272)
	DS9 = Freq(9956.064)
	E9  = Freq(10548.084)
	F9  = Freq(11175.304)
	FS9 = Freq(11839.82)
	G9  = Freq(12543.856)
	// G#9, MIDI note #128, the top of the MIDI tuning range.
	GS9 = Freq(13289.752)
	A9  = Freq(14080.)
	AS9 = Freq(14917.24)
	// B9. For most of adults, it is already beyond the hearing range.
	B9 = Freq(15804.264)
)

func Hz(f float32) Freq {
	return Freq(f)
}

func MIDI(note uint8) Freq { //nolint:cyclop
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
	return Freq(f)
}
