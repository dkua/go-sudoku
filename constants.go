package sudoku

import "strings"

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
	units := make(map[string][][]string)

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
	peers := make(map[string][]string)

	for unit, unit_list := range units {
		peer := make(map[string]bool)
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
