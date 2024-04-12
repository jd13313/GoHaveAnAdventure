package painter

type Palette struct {
	primaryColor   string
	secondaryColor string
	ternaryColor   string
}

func CreatePalette(p string, s string, t string) Palette {
	palette := Palette{
		primaryColor:   p,
		secondaryColor: s,
		ternaryColor:   t,
	}

	return palette
}

// I guess we're putting any predefined palettes here?
var PaletteGreens = CreatePalette("\x1b[38;5;2m", "\x1b[38;5;58m", "\x1b[38;5;42m")
