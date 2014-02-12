package sudoku

import (
	"bytes"
	"fmt"
)

// Return this map of values as a 2-D grid string.
func Display(values map[string]string) string {
	var buffer bytes.Buffer
	for _, row := range rows {
		for _, col := range cols {
			buffer.WriteString(fmt.Sprintf("%v", values[string(row)+string(col)]))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// Return this map of values as a formatted 2-D grid string.
func PrettyDisplay(values map[string]string) string {
	var buffer bytes.Buffer
	for r, row := range rows {
		for c, col := range cols {
			if (c == 3) || (c == 6) {
				buffer.WriteString("| ")
			}
			buffer.WriteString(fmt.Sprintf("%v ", values[string(row)+string(col)]))
		}
		buffer.WriteString("\n")
		if (r == 2) || (r == 5) {
			buffer.WriteString("------+-------+------\n")
		}
	}
	return buffer.String()
}
