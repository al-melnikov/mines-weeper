package field

import (
	"fmt"
	"unicode"

	"github.com/muesli/termenv"
)

func (field *Field) Print() {

	restoreConsole, err := termenv.EnableVirtualTerminalProcessing(termenv.DefaultOutput())
	if err != nil {
		panic(err)
	}
	defer restoreConsole()

	p := termenv.ColorProfile()

	fmt.Printf("\n")
	for i := 1; i <= field.m; i++ {
		if i < 10 {
			fmt.Printf("%v  ", termenv.String(fmt.Sprint(i)).Foreground(p.Color("#00FF00")))
		} else {
			fmt.Printf("%v ", termenv.String(fmt.Sprint(i)).Foreground(p.Color("#00FF00")))
		}
	}

	fmt.Printf("\n\n")
	for i := 1; i <= field.n; i++ {
		for j := 1; j <= field.m; j++ {
			if field.field[i][j] != '*' && field.field[i][j] != 'M' {
				field.field[i][j] = field.minesAround(i, j)
			}
			if field.field[i][j] == 'M' {
				field.field[i][j] = '*'
			}
			if field.field[i][j] <= 8 {
				var colStr string = "#"
				colStr += fmt.Sprintf("%X%X%X%X%X%X", 2*field.field[i][j]-1, 2*field.field[i][j]-1,
					16-2*field.field[i][j], 16-2*field.field[i][j],
					16-2*field.field[i][j], 16-2*field.field[i][j])
				//colStr += "0000"
				//fmt.Printf("%v  ", field.field[i][j])
				fmt.Printf("%v  ", termenv.String(fmt.Sprint(field.field[i][j])).Foreground(p.Color(colStr)))
			} else {
				fmt.Printf("%c  ", unicode.ToUpper(rune(field.field[i][j])))
			}
		}
		fmt.Printf("  %v", termenv.String(fmt.Sprint(i)).Foreground(p.Color("#00FF00")))
		fmt.Printf("\n")
	}
}

func (field *Field) PrintForPlayer() {

	restoreConsole, err := termenv.EnableVirtualTerminalProcessing(termenv.DefaultOutput())
	if err != nil {
		panic(err)
	}
	defer restoreConsole()

	p := termenv.ColorProfile()

	fmt.Printf("\n")
	for i := 1; i <= field.m; i++ {
		if i < 10 {
			fmt.Printf("%v  ", termenv.String(fmt.Sprint(i)).Foreground(p.Color("#00FF00")))
		} else {
			fmt.Printf("%v ", termenv.String(fmt.Sprint(i)).Foreground(p.Color("#00FF00")))
		}
	}

	fmt.Printf("\n\n")
	for i := 1; i <= field.n; i++ {
		for j := 1; j <= field.m; j++ {
			if field.isVis[i][j] {
				if field.field[i][j] <= 8 {
					fmt.Printf("%v  ", field.field[i][j])
				} else {
					fmt.Printf("%c  ", unicode.ToUpper(rune(field.field[i][j])))
				}

			} else {
				fmt.Printf("%c  ", '-')
			}
		}
		fmt.Printf("  %v", termenv.String(fmt.Sprint(i)).Foreground(p.Color("#00FF00")))
		fmt.Printf("\n")
	}
}
