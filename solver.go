package sudoku

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Convert grid to a map of possible values, {square: digits}, or return nil
// if a contradiction is detected.
func parseGrid(grid string) (map[string]string, error) {
	values := make(map[string]string, len(squares))
	for _, square := range squares {
		values[square] = digits
	}

	// Check validity of given grid
	regex := regexp.MustCompile(`[0-9]*\.*`)
	grid = strings.Join(regex.FindAllString(grid, -1), "")
	valid_grid, err := gridValues(grid)
	if err != nil {
		return nil, err
	}

	for s, d := range valid_grid {
		if strings.Contains(digits, d) {
			values, err = assign(values, s, d)
			if err != nil {
				return nil, err
			}
		}
	}
	return values, nil
}

// Convert grid into a map of {string: char} with "0" or "." for empties.
func gridValues(grid string) (map[string]string, error) {
	grid_values := make(map[string]string, len(grid))

	if len(grid) != 81 {
		return grid_values, errors.New("Puzzle does not contain 81 valid characters")
	}

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
	return grid_values, nil
}

// Eliminate all the other values (except d) from values[s] and propagate.
// Return values, except return nil if a contradiction is detected.
func assign(values map[string]string, s string, d string) (map[string]string, error) {

	// If square s has no possible values then fail.
	if len(values[s]) < 1 {
		return nil, fmt.Errorf("Value at %v has no possible digit left.", s)
	}

	other_values := strings.Replace(values[s], d, "", -1)
	if len(other_values) > 0 {
		for _, val := range other_values {
			if _, err := eliminate(values, s, string(val)); err != nil {
				return nil, err
			}
		}
	}
	return values, nil
}

// Eliminate d from values[s]; propagate when values or places <= 2.
// Return values, except return nil if a contradiction is detected.
func eliminate(values map[string]string, s string, d string) (map[string]string, error) {
	if !strings.Contains(values[s], d) {
		return values, nil // Digit d already eliminated from values[s]
	}

	values[s] = strings.Replace(values[s], d, "", -1) // Remove digit from values[s]

	// If a square s is eliminated to one value, then eliminate that value from the peers.
	value_length := len(values[s])
	if value_length == 0 {
		// Contradiction: Removed last possible value for square s
		return nil, fmt.Errorf("Cannot eliminate %v from %v because there are no more potential digits left.", d, s)
	} else if value_length == 1 {
		// One possibility left; therefore it is the solution to the square
		for _, peer := range peers[s] {
			if _, err := eliminate(values, peer, values[s]); err != nil {
				return nil, err // Contradiction: removed last value
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
			// Contradiction - no place for this value.
			return nil, fmt.Errorf("There is no place in unit %v to place %v", unit, d)
		} else if dplaces_length == 1 {
			// d can only be in one place in the unit, so assign it there.
			_, err := assign(values, dplaces[0], d)
			if err != nil {
				// Cannot solve puzzle so backtrack or it's unsolvable.
				return nil, err
			}
		}
	}
	return values, nil
}

// Using depth-first search and propagation, try all possible values.
func search(values map[string]string, ch chan map[string]string, err error) error {
	// Already failed so exit
	if err != nil {
		return err
	}

	// All squares have one possibility, puzzled solved so exit.
	solved := true
	for _, s := range squares {
		if len(values[s]) != 1 {
			solved = false
		}
	}
	if solved {
		ch <- values
		return nil // Congrats!
	}

	// Choose the unfilled square with the fewest possbilities.
	sq := ""
	min := len(digits) + 1
	for _, square := range squares {
		square_length := len(values[square])
		if (square_length > 1) && (min > square_length) {
			sq = square
			min = square_length

			if min == 2 {
				break
			}
		}
	}

	// For every possibility, recursively call search() and try to assign it to square s
	for _, d := range values[sq] {
		d := string(d)
		values_copy := clone(values)
		assigned_values, err := assign(values_copy, sq, d)
		if err != nil {
			continue
		}
		err = search(assigned_values, ch, err)
		if err == nil {
			return nil
		}
	}
	return errors.New("The search for solutions failed on this branch.")
}

func clone(values map[string]string) map[string]string {
	new_values := make(map[string]string, len(values))
	for key, val := range values {
		new_values[key] = val
	}
	return new_values
}

func Solve(grid string) (map[string]string, error) {
	ch := make(chan map[string]string)
	values, err := parseGrid(grid)
	if err != nil {
		return values, err
	}
	go search(values, ch, nil)
	return <-ch, nil
}
