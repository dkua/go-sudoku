package sudoku

import (
	"bytes"
	"fmt"
)

// Return this map of values as a formatted 2-D grid string.
func Display(values map[string]string, err error) string {
	if err != nil {
		return fmt.Sprintf("Sorry cannot solve the Sudoku here's why: %v", err)
	}

	var buffer bytes.Buffer
	for _, row := range rows {
		for _, col := range cols {
			buffer.WriteString(values[string(row)+string(col)])
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
