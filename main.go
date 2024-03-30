package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type LineStyle struct {
	CornerTopLeft     string `yaml:"CornerTopLeft"`
	CornerTopRight    string `yaml:"CornerTopRight"`
	Side              string `yaml:"Side"`
	Line              string `yaml:"Line"`
	CornerBottomLeft  string `yaml:"CornerBottomLeft"`
	CornerBottomRight string `yaml:"CornerBottomRight"`
}

type BoxStyles struct {
	BoldLines   LineStyle `yaml:"BoldLines"`
	DoubleLines LineStyle `yaml:"DoubleLines"`
	BasicLines  LineStyle `yaml:"BasicLines"`
}

type Config struct {
	BoxStyles BoxStyles `yaml:"BoxStyles"`
}

func main() {
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println(config.BoxStyles.BasicLines.CornerTopLeft)

	createTextBox("This took ~ 60 lines of code written in Go!", 120, config.BoxStyles.DoubleLines)
}

func createTextBox(t string, w int, s LineStyle) {
	chunks := splitStringIntoChunks(t, w-2)

	drawHorizontalBorder(w, "top", s)

	for _, chunk := range chunks {
		rw := w - len(chunk)

		// Print left border and chunk
		fmt.Print(s.Side, string(chunk))

		// Print empty spaces to fill the remaining width
		if rw > 0 {
			for j := 0; j <= rw-2; j++ {
				fmt.Print(" ")
			}
		}

		// Print right border
		fmt.Print(s.Side, "\n")
	}

	drawHorizontalBorder(w, "bottom", s)
}

func splitStringIntoChunks(s string, cs int) []string {
	var chunks []string
	runes := []rune(s)

	for i := 0; i < len(runes); i += cs {
		end := i + cs

		if end > len(runes) {
			end = len(runes)
		}

		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}

func drawHorizontalBorder(w int, d string, s LineStyle) {
	// Print left corner
	if d == "top" {
		fmt.Print(s.CornerTopLeft)
	} else {
		fmt.Print(s.CornerBottomLeft)
	}

	// Print horiz line
	for i := 0; i <= w-2; i++ {
		fmt.Print(s.Line)
	}

	// Print right corner
	if d == "top" {
		fmt.Print(s.CornerTopRight, "\n")
	} else {
		fmt.Print(s.CornerBottomRight, "\n")
	}
}
