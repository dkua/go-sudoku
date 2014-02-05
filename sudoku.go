package sudoku

import (
	"fmt"
	"strings"
)

var digits string = "123456789"
var rows string = "ABCDEFGHI"
var cols string = digits
var squares []string = cross(rows, cols)
var unitlist [][]string = createUnitList(rows, cols)
var units map[string][][]string = createUnits(unitlist, squares)
var peers map[string][]string = createPeers(units)

// Cross product of elements in A and in B.
func cross(A, B string) []string {
	cross_product := make([]string, 0)
	for _, a := range A {
		for _, b := range B {
			cross_product = append(cross_product, string(a)+string(b))
		}
	}
	return cross_product
}

func createUnitList(rows, cols string) [][]string {
	unitlist := make([][]string, 0)
	rs := []string{"ABC", "DEF", "GHI"}
	cs := []string{"123", "456", "789"}

	for _, col := range cols {
		unitlist = append(unitlist, cross(rows, string(col)))
	}

	for _, row := range rows {
		unitlist = append(unitlist, cross(string(row), cols))
	}

	for _, r := range rs {
		for _, c := range cs {
			unitlist = append(unitlist, cross(string(r), string(c)))
		}
	}
	return unitlist
}

func createUnits(unitlist [][]string, squares []string) map[string][][]string {
	units := make(map[string][][]string, 0)

	for _, s := range squares {
		unit := make([][]string, 0)
		for _, u := range unitlist {
			for _, u_string := range u {
				if strings.Contains(u_string, s) {
					unit = append(unit, u)
					break
				}
			}
		}
		units[s] = unit
	}
	return units
}

func createPeers(units map[string][][]string) map[string][]string {
	peers := make(map[string][]string, 0)

	for unit, unit_list := range units {
		peer := make(map[string]bool, 0)
		for _, unit_sublist := range unit_list {
			for _, u := range unit_sublist {
				if _, present := peer[u]; !present {
					if u != unit {
						peer[u] = true
					}
				}
			}
		}

		for key, _ := range peer {
			peers[unit] = append(peers[unit], key)
		}
	}
	return peers
}

// Convert grid to a map of possible values, {square: digits}, or return nil
// if a contradiction is detected.
func parseGrid(grid string) map[string]string {
	values := make(map[string]string, len(squares))
	for _, square := range squares {
		values[square] = digits
	}
	for s, d := range gridValues(grid) {
		if strings.Contains(digits, d) {
			values = assign(values, s, d)
			if values == nil {
				return nil
			}
		}
	}
	return values
}

// Convert grid into a map of {string: char} with "0" or "." for empties.
func gridValues(grid string) map[string]string {
	grid_values := make(map[string]string, len(grid))
	chars := make([]string, 0)
	for _, c := range grid {
		char := string(c)
		if (strings.Contains(digits, char)) || (strings.Contains("0.", char)) {
			chars = append(chars, char)
		}
	}

	for index, square := range squares {
		grid_values[square] = chars[index]
	}
	return grid_values
}

// Eliminate all the other values (except d) from values[s] and propagate.
// Return values, except return nil if a contradiction is detected.
func assign(values map[string]string, s string, d string) map[string]string {
	other_values := strings.Replace(values[s], d, "", -1)
	for _, val := range other_values {
		if eliminate(values, s, string(val)) == nil {
			return nil
		}
	}
	return values
}

// Eliminate d from values[s]; propagate when values or places <= 2.
// Return values, except return nil if a contradiction is detected.
func eliminate(values map[string]string, s string, d string) map[string]string {
	if !strings.Contains(values[s], d) {
		return values // Already eliminated
	}

	values[s] = strings.Replace(values[s], d, "", -1)

	// If a square s is eliminated to one value, then eliminate that value from the peers.
	value_length := len(values[s])
	if value_length == 0 {
		return nil // Already eliminated.
	} else if value_length == 1 {
		for _, peer := range peers[s] {
			if eliminate(values, peer, values[s]) == nil {
				return nil // Contradiction - remove last value.
			}
		}
	}

	// If a unit is reduced to to only one place for a value, then put it there.
	for _, unit := range units[s] {
		dplaces := make([]string, 0)
		for _, square := range unit {
			if strings.Contains(values[square], d) {
				dplaces = append(dplaces, square)
			}
		}
		dplaces_length := len(dplaces)
		if dplaces_length == 0 {
			return nil // Contradiction - no place for this value.
		} else if dplaces_length == 1 {
			// d can only be in one place in the unit, so assign it there.
			if assign(values, dplaces[0], d) == nil {
				return nil
			}
		}
	}

	return values
}

// Using depth-first search and propagation, try all possible values.
func search(values map[string]string) map[string]string {
	if values == nil {
		return nil // Already failed earlier.
	}
	solved := true
	for _, value := range values {
		if len(string(value)) != 1 {
			solved = false
		}
	}
	if solved {
		return values // Congrats!
	}

	// Choose the unfilled square with the fewest possbilities.
	min_square := ""
	min_length := len(digits) + 1
	for _, square := range squares {
		square_length := len(values[square])
		if (square_length > 1) && (min_length > square_length) {
			min_square = square
			min_length = square_length
		}
	}

	solution_chan := make(chan map[string]string)
	for _, d := range values[min_square] {
		go func(dd string) {
			values_copy := clone(values)
			val := search(assign(values_copy, min_square, dd))
			if val != nil {
				solution_chan <- val
			}
		}(string(d))
	}
	return <-solution_chan
}

func clone(values map[string]string) map[string]string {
	new_values := make(map[string]string, len(values))
	for key, val := range values {
		new_values[key] = val
	}
	return new_values
}

// Display this map of values as a 2-D grid.
func Display(values map[string]string) {
	for r, row := range rows {
		for c, col := range cols {
			if (c == 3) || (c == 6) {
				fmt.Printf("| ")
			}
			fmt.Printf("%v ", values[string(row)+string(col)])
		}
		fmt.Println()
		if (r == 2) || (r == 5) {
			fmt.Println("------+------+------")
		}
	}
	fmt.Println()
}

func Solve(grid string) map[string]string {
	return search(parseGrid(grid))
}
