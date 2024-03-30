package main

func main() {
	DialogueBox1 := newDialogueBox(120, "DoubleLines", "Hello, World!", "Red")
	DialogueBox2 := newDialogueBox(120, "CurvyLines", "Hello, World!", "Green")
	DialogueBox3 := newDialogueBox(120, "ThinLines", "Hello, World!", "Cyan")
	DialogueBox4 := newDialogueBox(120, "BoldLines", "Hello, World!", "Purple")

	DialogueBox1.Print()
	DialogueBox2.Print()
	DialogueBox3.Print()
	DialogueBox4.Print()
}
