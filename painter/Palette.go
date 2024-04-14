package painter

import "fmt"

type Palette struct {
	primaryColor    string
	secondaryColor  string
	ternaryColor    string
	quaternaryColor string
}

func CreatePalette(p Color, s Color, t Color, q Color) Palette {
	palette := Palette{
		primaryColor:    p.GetANSI(),
		secondaryColor:  s.GetANSI(),
		ternaryColor:    t.GetANSI(),
		quaternaryColor: q.GetANSI(),
	}

	return palette
}

func (p Palette) Info() {
	fmt.Println("Palette Info:")

	Draw(`
███ [Primary]
▓▓▓ [Secondary] 		
▒▒▒ [Ternary]
░░░ [Quaternary]
	`, p)
}
