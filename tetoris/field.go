package tetoris

import (
	"fmt"
	"strconv"
	"strings"
)

// Field is tetoris field
type Field struct {
	buf []uint
}

const separator = "-"

// New returns new tetoris field
func New(input string) *Field {
	inputs := strings.Split(input, separator)
	buf := make([]uint, len(inputs))
	for i, row := range inputs {
		n, err := strconv.ParseInt("0x"+row, 0, 64)
		if err != nil {
			panic(err)
		}
		buf[i] = uint(n)
	}
	return &Field{buf: buf}
}

// Map returns 2 dimention mapped field's buffer
func (f *Field) Map() [][]uint {
	rows := make([][]uint, 8)
	for i := 0; i < 8; i++ {
		rows[i] = make([]uint, len(f.buf))
	}

	for y := 0; y < 8; y++ {
		for x, line := range f.buf {
			rows[y][x] = (line >> uint(y) & 1)
		}
	}

	return rows
}

// TickedMap returns solved line map
func (f *Field) TickedMap() [][]uint {
	m := f.Map()
	var solve [][]uint
	l := len(m)
	for i := 0; i < l; i++ {
		n := m[i]
		if !all(n) {
			solve = append(solve, n)
		}
	}
	for i := len(solve); i < l; i++ {
		n := m[i]
		solve = append(solve, make([]uint, len(n)))
	}

	return solve
}

func all(s []uint) bool {
	i := 0
	for {
		if s[i] < 1 {
			return false
		}

		if i < len(s)-1 {
			i = i + 1
		} else {
			return true
		}
	}
}

func (f *Field) solve() string {
	m := f.TickedMap()
	buf := make([]uint, len(f.buf))

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			buf[x] = buf[x] | (m[y][x] << uint(y))
		}
	}

	result := make([]string, len(buf))

	for i, s := range buf {
		result[i] = fmt.Sprintf("%02x", s)
	}

	return strings.Join(result, separator)
}

// Description makes visualize field's data
func (f *Field) Description() {
	m := f.Map()

	for i := len(m) - 1; i >= 0; i-- {
		l := m[i]
		fmt.Println(l)
	}

	fmt.Println("")

	n := f.TickedMap()
	for i := len(n) - 1; i >= 0; i-- {
		l := n[i]
		fmt.Println(l)
	}
}
