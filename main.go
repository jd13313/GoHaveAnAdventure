package main

func main() {
	lText1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. <np> Phasellus metus mauris, porttitor <lb> et consequat quis, posuere ut mi. Duis nisi enim, vestibulum nec arcu vel, aliquam fringilla felis. Donec aliquam augue ex, eget semper metus tempor ut. Nam aliquet fermentum sapien non eleifend. Mauris elementum et metus ut pulvinar. Phasellus imperdiet pretium dui a vulputate. Mauris mi magna, tempor nec nibh quis, ultrices ullamcorper dui."
	lText2 := "Four score and seven years ago our fathers brought forth on this continent, <np> a new nation, conceived in Liberty, and dedicated to the proposition that all men are created equal. Now we are engaged in a great civil war, testing whether that nation, or any nation so conceived and so dedicated, can long endure. We are met on a great battle-field of that war. We have come to dedicate a portion of that field, as a final resting place for those who here gave their lives that that nation might live. It is altogether fitting and proper that we should do this."
	newDialogueBox(150, "DoubleLines", lText1, "Cyan").Print()
	newDialogueBox(100, "CurvyLines", lText2, "Green").Print()
	newDialogueBox(20, "ThinLines", "Hello World", "Red").Print()
}
