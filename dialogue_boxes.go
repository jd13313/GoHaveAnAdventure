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

type BoxStyles map[string]LineStyle

type Config struct {
	BoxStyles BoxStyles `yaml:"BoxStyles"`
}

type DialogueBox struct {
	Width int
	Style LineStyle
	Text  string
	Color string
}

func newDialogueBox(w int, s string, t string, c string) DialogueBox {
	// Handle config file
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// Define color options
	colorOptions := map[string]string{
		"Red":    "\033[91m",
		"Green":  "\033[92m",
		"Yellow": "\033[93m",
		"Blue":   "\033[94m",
		"Purple": "\033[95m",
		"Cyan":   "\033[96m",
		"White":  "\033[97m",
		"Reset":  "\033[0m",
	}

	// Create dialogue box
	dialogueBox := DialogueBox{
		Width: w,
		Style: config.BoxStyles[s],
		Text:  t,
		Color: colorOptions[c],
	}

	return dialogueBox
}

func (d DialogueBox) Print() {
	chunks := splitStringIntoChunks(d.Text, d.Width-4)

	// Set box color
	fmt.Print(d.Color)

	// Print top border
	drawHorizontalBorder(d.Width, "top", d.Style)

	for _, chunk := range chunks {
		rw := d.Width - len(chunk)

		// Print left border and chunk
		fmt.Print(d.Style.Side, " ", string(chunk))

		// Print empty spaces to fill the remaining width
		if rw > 0 {
			for j := 0; j <= rw-4; j++ {
				fmt.Print(" ")
			}
		}

		// Print right border
		fmt.Print(" ", d.Style.Side, "\n")
	}

	drawHorizontalBorder(d.Width, "bottom", d.Style)

	// Reset color
	fmt.Print("\033[0m")
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
