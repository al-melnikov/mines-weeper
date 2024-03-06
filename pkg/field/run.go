package field

import (
	"fmt"

	"github.com/al-melnikov/mines-weeper/pkg/util"

	"github.com/muesli/termenv"
)

func Run(n int, m int, minesNum int) error {

	f, err := New(n, m, minesNum)
	if err != nil {
		return err
	}

	p := termenv.ColorProfile()

	for {
		util.ClearTerminal()
		f.PrintForPlayer()
		gameOver := f.Input()

		if gameOver {
			util.ClearTerminal()
			f.Print()
			fmt.Printf("%v\n", termenv.String("GAME OVER!!").Bold().Foreground(p.Color("#FF0000")))
			break
		}
		if f.winCondition() {
			util.ClearTerminal()
			f.Print()
			fmt.Printf("%v\n", termenv.String("YOU WON!!").Bold().Foreground(p.Color("#00FF00")))
			break
		}
	}

	return nil
}
