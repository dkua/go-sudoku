package sudoku

import (
	"bytes"
	"math/rand"
)

func CreatePuzzle(n int) string {
	// Make a random puzzle with n or more assignments. Restart on contradictions.
	// Note the resulting puzzle is not guaranteed to be solvable, but empirically
	// about 99.8% of them are solvable. Some have multiple solutions.

	values := make(map[string]string, len(squares))
	for _, square := range squares {
		values[square] = digits
	}

	for _, square := range shuffle(squares) {
		if assign(values, square, randomChoice(values[square])) == nil {
			break
		}

		ds := make([]string, 0)
		for _, s := range squares {
			if len(values[s]) == 1 {
				ds = append(ds, values[s])
			}
		}
		if (len(ds) >= n) && (len(uniqueArray(ds)) >= 8) {
			var buffer bytes.Buffer
			for _, s := range squares {
				if len(values[s]) == 1 {
					buffer.WriteString(values[s])
				} else {
					buffer.WriteString(".")
				}
			}
			return buffer.String()
		}
	}
	return CreatePuzzle(n)
}

func randomChoice(seq string) string {
	random_choice := rand.Intn(len(seq))
	item := seq[random_choice]
	return string(item)
}

func shuffle(seq []string) []string {
	seq_length := len(seq)
	shuffled_array := make([]string, seq_length)
	permutation := rand.Perm(seq_length)
	for index, value := range permutation {
		shuffled_array[value] = seq[index]
	}
	return shuffled_array
}

func uniqueArray(seq []string) []string {
	unique_array := make([]string, 0)
	for _, item := range seq {
		unique_array = appendIfMissing(unique_array, item)
	}
	return unique_array
}

func appendIfMissing(seq []string, item string) []string {
	for _, ele := range seq {
		if ele == item {
			return seq
		}
	}
	return append(seq, item)
}
