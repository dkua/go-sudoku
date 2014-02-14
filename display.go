package sudoku

import (
	"bytes"
	"fmt"
)

// Return this map of values as a formatted 2-D grid string.
func Display(values map[string]string, prettify bool, err error) string {
	if err != nil {
		return "Sorry couldn't solve your Sudoku puzzle"
	}

	var buffer bytes.Buffer
	for r, row := range rows {
		for c, col := range cols {
			if prettify && ((c == 3) || (c == 6)) {
				buffer.WriteString("| ")
			}
			buffer.WriteString(fmt.Sprintf("%v ", values[string(row)+string(col)]))
		}
		buffer.WriteString("\n")
		if prettify && ((r == 2) || (r == 5)) {
			buffer.WriteString("------+-------+------\n")
		}
	}
	return buffer.String()
}
