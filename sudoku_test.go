package sudoku

import (
	"fmt"
	"strings"
	"testing"
)

func TestCross(t *testing.T) {
	squares := cross(rows, cols)
	expected := 81
	result := len(squares)
	if result != expected {
		t.Errorf("Wrong number of squares. Expected %v found %v\n", expected, result)
	}
	fmt.Printf("Squares: %v\n", squares)
}

func TestCreateUnitList(t *testing.T) {
	unitlist := createUnitList(rows, cols)
	expected := 27
	result := len(unitlist)
	if result != expected {
		t.Errorf("Wrong number of units. Expected %v found %v\n", expected, result)
	}
	fmt.Printf("Unitlist: %v\n", unitlist)
}

func TestCreateUnits(t *testing.T) {
	units := createUnits(unitlist, squares)
	expected := 3
	failed := false
	failed_keys := make([]string, 0)
	for key, val := range units {
		result := len(val)
		if result != expected {
			failed = true
			failed_keys = append(failed_keys, key)
		}
	}

	if failed {
		t.Errorf("Wrong number of units. Expected %v but didn't find in %v\n", expected, failed_keys)
	}
	fmt.Printf("Units: %v\n", units)
}

func TestCreatePeers(t *testing.T) {
	peers := createPeers(units)
	expected := 20
	failed := false
	failed_keys := make([]string, 0)
	for key, val := range peers {
		result := len(val)
		if result != expected {
			failed = true
			failed_keys = append(failed_keys, key)
		}
	}

	if failed {
		t.Errorf("Wrong number of peers. Expected %v but didn't find in %v\n", expected, failed_keys)
	}
	fmt.Printf("Peers: %v\n", peers)
}

func TestParseGrid(t *testing.T) {
	grid := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	values := parseGrid(grid)
	expected := 81
	result := len(values)
	if result != expected {
		t.Errorf("Wrong number of squares in values. Expected %v but didn't find in %v\n", expected, result)
	}
	fmt.Printf("Values: %v\n", values)
}

func TestCreatePuzzleLength(t *testing.T) {
	puzzle := CreatePuzzle(17)
	expected := 17
	result := len(strings.Replace(puzzle, ".", "", -1))
	if result != expected {
		t.Errorf("Wrong number of numbers in puzzle. Expected %v but found %v\n", expected, result)
	}
	fmt.Printf("Puzzle: %v\n", puzzle)
}

func TestRandomSolve(t *testing.T) {
	puzzle := CreatePuzzle(17)
	fmt.Printf("Puzzle: %v\n", puzzle)
	fmt.Println(Display(Solve(puzzle)))
}
