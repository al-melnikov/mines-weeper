package field

import (
	"fmt"
)

func (field *Field) Input() bool {

	var i int
	var j int
	var c string

	fmt.Scanf("%v %v %s\n", &i, &j, &c)

	if len(c) > 0 && (c[0] == 'm' || c[0] == 'M') {
		if field.field[i][j] == '*' {
			field.field[i][j] = 'M'
		}
		if field.field[i][j] == '-' {
			field.field[i][j] = 'm'
		}
		field.isVis[i][j] = true
		return false
	}

	if field.isMine(i, j) {
		return true
	}
	field.modify(i, j)
	return false
}
