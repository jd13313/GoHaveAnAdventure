package main

import "/dialogueBox/dialogueBox" // Import the missing package

func main() {
	dividedText := "This is a test of the emergency broadcast system. This is only a test. <d> If this were a real emergency, you would be instructed where to go and what to do. <p> This is only a test. <l> But like..what if?"
	dialogueBox.createDialogueBox(100, "BoldLines", dividedText, dialogueBox.ColorRed).Print()
	dialogueBox.createDialogueBox(100, "CurvyLines", dividedText, dialogueBox.ColorCyan).Print()
	dialogueBox.createDialogueBox(100, "BasicLines", dividedText, dialogueBox.ColorBlue).Print()
	dialogueBox.createDialogueBox(100, "ThinLines", dividedText, dialogueBox.ColorPurple).Print()
}
