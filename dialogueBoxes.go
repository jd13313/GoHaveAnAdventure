package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type LineStyle struct {
	CornerTopLeft     string `yaml:"CornerTopLeft"`
	CornerTopRight    string `yaml:"CornerTopRight"`
	Side              string `yaml:"Side"`
	DividerLeft       string `yaml:"DividerLeft"`
	DividerRight      string `yaml:"DividerRight"`
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

func getConfig() Config {
	data, err := os.ReadFile("configs/config.yaml")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return config
}

func newDialogueBox(w int, s string, t string, c string) DialogueBox {
	config := getConfig()

	// Create dialogue box
	dialogueBox := DialogueBox{
		Width: w,
		Style: config.BoxStyles[s],
		Text:  t,
		Color: c,
	}

	return dialogueBox
}

func (d DialogueBox) Print() {
	lines := splitStringIntoLines(d.Text, d.Width, d.Style)

	// Set box color
	fmt.Print(d.Color)

	// Print top border
	drawHorizontalBorder(d.Width, "top", d.Style)

	for _, l := range lines {
		l = strings.TrimSpace(l)
		remainingWidth := d.Width - len([]rune(l)) - 4
		whitespaceStart := " "
		var whitespaceEnd string

		if remainingWidth < 0 {
			remainingWidth = 0
		}

		whitespaceEnd = strings.Repeat(" ", remainingWidth+2)

		if l == "<d>" {
			fmt.Println(d.Style.DividerLeft + strings.Repeat(d.Style.Line, d.Width-1) + d.Style.DividerRight)
		} else {
			fmt.Println(d.Style.Side + whitespaceStart + l + whitespaceEnd + d.Style.Side)
		}
	}

	drawHorizontalBorder(d.Width, "bottom", d.Style)

	// Reset color
	fmt.Print("\033[0m")
}

func drawHorizontalBorder(w int, d string, s LineStyle) {
	// Print left character
	if d == "top" {
		fmt.Print(s.CornerTopLeft)
	} else if d == "bottom" {
		fmt.Print(s.CornerBottomLeft)
	}

	// Print horiz line
	for i := 0; i <= w-2; i++ {
		fmt.Print(s.Line)
	}

	// Print right character
	if d == "top" {
		fmt.Print(s.CornerTopRight, "\n")
	} else if d == "bottom" {
		fmt.Print(s.CornerBottomRight, "\n")
	}
}

func splitStringIntoLines(s string, w int, ls LineStyle) []string {
	var lines []string
	words := strings.Fields(s)
	ln := 0

	for _, word := range words {
		if ln >= len(lines) {
			lines = append(lines, "")
		}

		// Handle special, reserved words such as <np> and <lb> for New Paragraph and Line Break
		switch word {
		case "<p>": // New Paragraph
			ln = ln + 2
			lines = append(lines, " ")
			continue
		case "<l>": // New Line
			ln++
			lines = append(lines, " ")
			continue
		case "<d>": // Divider
			ln = ln + 2
			lines = append(lines, "<d>")
			continue
		}

		// Handle normal words
		lineLength := len([]rune(lines[ln]))
		wordLength := len([]rune(word))

		if lineLength+wordLength <= w-4 {
			lines[ln] += word + " "
		} else {
			ln++
			lines = append(lines, word+" ")
		}
	}

	return lines
}
