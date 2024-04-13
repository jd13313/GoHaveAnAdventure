package main

import (
	"main/dialogueBox"
	"main/painter"
)

func main() {
	dividedText := "This is a test of the emergency broadcast system. This is only a test. <d> If this were a real emergency, you would be instructed where to go and what to do. <p> This is only a test. <l> But like..what if?"
	dialogueBox.CreateDialogueBox(dialogueBox.StyleBasic, dividedText, dialogueBox.ColorRed).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleCurved, dividedText, dialogueBox.ColorCyan).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleDouble, dividedText, dialogueBox.ColorBlue).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleThin, dividedText, dialogueBox.ColorPurple).Print()
	dialogueBox.CreateDialogueBox(dialogueBox.StyleCurved, "Alert !!!! <d> HP: 2 <d> Your health is dangerously low! Consider using a potion or eating something.", dialogueBox.ColorGreen).Print()

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
   ░░░░░░░░░░░░░▓░        ░░░░░░░░░░░░░░░░
    ░░░░░░░░░░░░▓░░ ███ ░░░░░░░░░░░░░░░░░
    ░░░░░░░░░░░░░█░░▀█▀░░░░░░░░░░░░░░░░░░
     ░░░░░░░░░░░░░█▀▓█▓█░░░░░░░░░░░░░░░░
       ░░░░░░░░░░░░░▓▓▓░█░░░░░░░░░░░░░
           ░░░░░░░░░█▓█░░█░░░░░░
                    █ █
                   ▄█ █▄
  `, painter.PaletteGreens)
}
