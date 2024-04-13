package painter

import "strconv"

type Color struct {
	r int
	g int
	b int
}

func CreateColor(r int, g int, b int) Color {
	return Color{r, g, b}
}

func (c Color) GetANSI() string {
	rStr := strconv.Itoa(c.r)
	gStr := strconv.Itoa(c.g)
	bStr := strconv.Itoa(c.b)

	return "\033[38;2;" + rStr + ";" + gStr + ";" + bStr + "m"
}
