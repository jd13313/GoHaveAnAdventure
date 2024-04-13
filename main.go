package main

import (
	"main/dialogueBox"
	"main/painter"
)

func main() {
	dividedText := "This is a test of the emergency broadcast system. This is only a test. <d> If this were a real emergency, you would be instructed where to go and what to do. <p> This is only a test. <l> But like..what if?"
	colorRed := painter.CreateColor(255, 0, 0)
	colorGreen := painter.CreateColor(0, 255, 0)
	colorBlue := painter.CreateColor(0, 0, 255)
	colorPurple := painter.CreateColor(255, 0, 255)

	dialogueBox.CreateDialogueBox(dialogueBox.StyleBasic, dividedText, colorRed).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleCurved, dividedText, colorBlue).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleDouble, dividedText, colorGreen).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleThin, dividedText, colorPurple).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleCurved, "Alert !!!! <d> HP: 2 <d> Your health is dangerously low! Consider using a potion or eating something.", colorGreen).Print()

	painter.Draw(`                     
                         ███████████                        
                        █▓▓▓▓▓▓▓▓▓▓▓█                       
                       █▓▓▒▓▓▓▓▓▓▓▒▓▓█                      
                      █▓▓▒▒▒▓▓▓▓▓▒▒▒▓▓█                     
                      █▓▓▓▒▓▓▓▓▓▓▓▒▓▓▓█                     
                      █▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓█                     
                      █▓▒▓▓▓▓▓▓▓▓▓▓▓▒▓█                     
                      █▓▓▒▒▒▒▒▒▒▒▒▒▒▓▓█                     
                       █▓▓▒▒▒▒▒▒▒▒▒▓▓█                      
                        █▓▓▓▓▓▓▓▓▓▓▓█                       
                         ███████████                        

	`, painter.PaletteGreens)

	painter.Draw(`

    ███████  ████████████
    ██▒▒▒▒▒███▒▒▒▒▒▒▒▒▒▒▒███
  ███▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒███
████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██
██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█
██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████
█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██
█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒███
█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒███
██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█
██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█
 █▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█
 █▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██
 ██▒▒▒▒▒▒▒▒▒▒▒▒▓▒▒▒▒▓▒▒▒▒▒▒▒▒▒▒▒▒▒█
  ███▒▒▒▒▒▒▒▒▒▓▓▓▒▓▓▓▓▒▒▒▒▒▒▒▒▒▒███
    ███▒▒▒▒▒▒█▓▓▓▓▓▓▓█▒▒▒▒▒██████
       ██▒▒▒▒█▓▓▓▓▓▓▓███████
        ████▒█▓▓▓▓▓▓▓█
           ███▓▓▓█▓▓▓█
             █▓▓██▓▓▓█
             █▓███▓▓▓█
            ██▓██▓▓▓▓█
            █▓▓▓▓▓▓▓▓██
            █▓▓▓▓▓▓▓▓▓█
            █▓▓▓▓▓▓▓▓▓██
            █▓▓▓▓▓▓▓▓▓▓██
            █▓▓▓▓▓▓▓▓▓▓▓██
           ██▓▓▓▓▓▓▓▓▓▓▓▓██
          ██▓▓▓▓▓▓▓▓▓▓▓▓▓▓█
            ██████████████
    `, painter.PaletteGreens)

	painter.Draw(`
        Primary: █
        Secondary: ▓
        Ternary: ▒
    `, painter.PaletteGreens)

	painter.Draw(`
  ░░░                                     ░░░░░░░
  ░░░░░░░                             ░░░░░░░░░░
  ░░░░░░░░░     ▓                 ░░░░░░░░░░░░░░
  ░░░░░░░░░░░   ▓              ░░░░░░░░░░░░░░░
  ░░░░░░░░░░░░░░▓           ░░░░░░░░░░░░░░░░
   ░░░░░░░░░░░░░▓░  ▓ ▓   ░░░░░░░░░░░░░░░░
    ░░░░░░░░░░░░▓░░ ███ ░░░░░░░░░░░░░░░░░
    ░░░░░░░░░░░░░█░░▀█▀░░░░░░░░░░░░░░░░░░
     ░░░░░░░░░░░░░█▀▓█▓█░░░░░░░░░░░░░░░░
       ░░░░░░░░░░░░░▓▓▓░█░░░░░░░░░░░░░
           ░░░░░░░░░█▓█░░█░░░░░░
                    █ █
                   ▄█ █▄
  `, painter.PaletteGreens)
}
