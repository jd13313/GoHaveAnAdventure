package main

import (
	"fmt"
	"main/painter"
	"main/ui"

	"main/utility"
)

func main() {
	// rand.Seed(time.Now().UnixNano())

	palette := painter.PaletteWhites

	hokeyName := utility.RandomResponse([]string{
		"stranger",
		"pal",
		"buddy",
		"friend",
		"partner",
	})

	ui.CreateDialogueBox(
		ui.StyleCurved,
		`
      You come upon an old man. He turns to you and speaks:
      <d> Well hello there! Haven't seen you around these parts before.
      <p> What's your name, `+hokeyName+`?
    `,
		palette.GetPrimary(),
	).Print()

	var name string

	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error scanning name: ", err)
	}

	nameReaction := utility.RandomResponse([]string{
		"an unusual name for these parts...",
		"as fine a name as any!",
		"a dandy name, if I do say so myself.",
		"quite the name you've got there!",
		"strange...that's my name too! Just kidding.",
	})

	ui.CreateDialogueBox(
		ui.StyleCurved,
		`
      The man nods and replies: <d>
      `+name+`? That's `+nameReaction+`
      So what brings you 'round here?
       <d> 1) Just passing through. <l> 2) Looking for adventure. <l> 3) I'm lost. <l> 4) I'm on a quest.
    `,
		palette.GetPrimary(),
	).Print()
}
