package main

func main() {
	dividedText := "This is a test of the emergency broadcast system. This is only a test. <d> If this were a real emergency, you would be instructed where to go and what to do. <p> This is only a test. <l> But like..what if?"
	newDialogueBox(100, "BoldLines", dividedText, ColorRed).Print()
	newDialogueBox(100, "CurvyLines", dividedText, ColorCyan).Print()
	newDialogueBox(100, "BasicLines", dividedText, ColorBlue).Print()
	newDialogueBox(100, "ThinLines", dividedText, ColorPurple).Print()
}
