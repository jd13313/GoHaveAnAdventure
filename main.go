package main

import (
	"main/dialogueBox"
)

func main() {
	dividedText := "This is a test of the emergency broadcast system. This is only a test. <d> If this were a real emergency, you would be instructed where to go and what to do. <p> This is only a test. <l> But like..what if?"
	dialogueBox.CreateDialogueBox(dialogueBox.BasicLines, dividedText, dialogueBox.ColorRed).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.CurvyLines, dividedText, dialogueBox.ColorCyan).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.DoubleLines, dividedText, dialogueBox.ColorBlue).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.ThinLines, dividedText, dialogueBox.ColorPurple).Print()
}
