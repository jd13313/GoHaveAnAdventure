package painter

import (
	"fmt"
)

// type Palette struct {
// 	primaryColor   string
// 	secondaryColor string
// }

// const PaintCharacters = {
// 	"PrimaryChar":"█",
// 	"SecondaryChar": "▓",
// }

func Draw(a string) {
	for _, c := range a {
		switch c {
		case '█':
			fmt.Print("\x1b[38;5;0m" + string(c))
		case '▓':
			fmt.Print("\x1b[38;5;226m" + string(c))
		case '▒':
			fmt.Print("\x1b[38;5;244m" + string(c))
		default:
			fmt.Print("\033[0m" + string(c))
		}
	}
}
