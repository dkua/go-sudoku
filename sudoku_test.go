package sudoku

import (
	"fmt"
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
