package main

import (
	"main/painter"
	"main/ui"
)

func main() {
	dividedText := "This is a test of the emergency broadcast system. This is only a test. <d> If this were a real emergency, you would be instructed where to go and what to do. <p> This is only a test. <l> But like..what if?"
	colorRed := painter.CreateColor(255, 0, 0)
	colorGreen := painter.CreateColor(0, 255, 0)
	colorBlue := painter.CreateColor(0, 0, 255)
	colorPurple := painter.CreateColor(255, 0, 255)

	ui.CreateDialogueBox(ui.StyleBasic, dividedText, colorRed).Print()
	ui.CreateDialogueBox(ui.StyleCurved, dividedText, colorBlue).Print()
	ui.CreateDialogueBox(ui.StyleDouble, dividedText, colorGreen).Print()
	ui.CreateDialogueBox(ui.StyleThin, dividedText, colorPurple).Print()
	ui.CreateDialogueBox(ui.StyleCurved, "Alert !!!! <d> HP: 2 <d> Your health is dangerously low! Consider using a potion or eating something.", colorGreen).Print()

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

	`, painter.GoldenToad)

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
  `, painter.GoldenToad)

	painter.GoldenToad.Info()
}
