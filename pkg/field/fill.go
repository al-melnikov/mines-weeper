package field

import (
	"errors"
	"math/rand"
)

func (field *Field) fill(minesNum int) error {
	if field.n == 0 || field.m == 0 {
		return errors.New("empty field")
	}

	if minesNum >= field.n*field.m {
		return errors.New("too many mines")
	}

	for i := 1; i <= field.n; i++ {
		for j := 1; j <= field.m; j++ {
			field.field[i][j] = '-'
		}
	}

	for t := 0; t < minesNum; t++ {
		i := rand.Intn(field.n) + 1 // [0, n) -> [1, n]
		j := rand.Intn(field.m) + 1
		if field.field[i][j] == '*' {
			t--
			continue
		}
		field.field[i][j] = '*'
	}

	return nil
}
