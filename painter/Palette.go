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
	Draw(`
	███ [Primary]
	▓▓▓ [Secondary] 		
	▒▒▒ [Ternary]
	░░░ [Quaternary]
	`, p)
	fmt.Println("")
}

func (p Palette) GetPrimary() string {
	return p.primaryColor
}

func (p Palette) GetSecondary() string {
	return p.secondaryColor
}

func (p Palette) GetTernary() string {
	return p.ternaryColor
}

func (p Palette) GetQuaternary() string {
	return p.quaternaryColor
}
