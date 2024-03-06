package field

type Field struct {
	n     int
	m     int
	field [][]byte
	isVis [][]bool
}

func New(n int, m int, minesNum int) (Field, error) {
	f := make([][]byte, n+2)
	v := make([][]bool, n+2)
	for i := 0; i < n+2; i++ {
		f[i] = make([]byte, m+2)
		v[i] = make([]bool, m+2)
	}

	field := Field{n, m, f, v}

	err := field.fill(minesNum)
	if err != nil {
		return field, err
	}
	return field, nil
}

func (field *Field) isMine(i int, j int) bool {
	return field.field[i][j] == '*' || field.field[i][j] == 'M'
}

func (field *Field) modify(i int, j int) {
	if field.field[i][j] == '-' || field.field[i][j] == 'm' {
		n := field.minesAround(i, j)
		c := byte(n)
		field.field[i][j] = c
		field.isVis[i][j] = true
		if n == 0 {
			field.modify(i-1, j)
			field.modify(i-1, j-1)
			field.modify(i-1, j+1)
			field.modify(i+1, j)
			field.modify(i+1, j+1)
			field.modify(i+1, j-1)
			field.modify(i, j-1)
			field.modify(i, j+1)
		}
	}
}

func (field *Field) minesAround(i int, j int) uint8 {
	var res uint8 = 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if field.isMine(x, y) {
				res++
			}
		}
	}

	if field.field[i][j] == '*' {
		res--
	}

	return res
}

func (field *Field) winCondition() bool {
	for i := 1; i <= field.n; i++ {
		for j := 1; j <= field.m; j++ {
			if !field.isVis[i][j] && (field.field[i][j] == '-' || field.field[i][j] == 'm') {
				return false
			}
		}
	}
	return true
}
