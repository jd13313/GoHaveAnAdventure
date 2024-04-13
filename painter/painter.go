package painter

import (
	"fmt"
)

func Draw(a string, p Palette) {
	for _, c := range a {
		switch c {
		case '█':
			fmt.Print(p.primaryColor + string(c))
		case '▓':
			fmt.Print(p.secondaryColor + string(c))
		case '▒':
			fmt.Print(p.ternaryColor + string(c))
		case '░':
			fmt.Print(p.quaternaryColor + string(c))
		default:
			fmt.Print("\033[0m" + string(c))
		}
	}
}
