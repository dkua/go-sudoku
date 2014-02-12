package sudoku

import (
	"bytes"
	"math/rand"
	"time"
)

// Make a random puzzle with n or more assignments. Restart on contradictions.
// Note the resulting puzzle is not guaranteed to be solvable, but empirically
// about 99.8% of them are solvable. Some have multiple solutions.
func CreatePuzzle(n int) string {
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
	rand.Seed(time.Now().UnixNano())
	random_choice := rand.Intn(len(seq))
	item := seq[random_choice]
	return string(item)
}

// Implementation of the "Knuth Shuffle" for string arrays
func shuffle(seq []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := len(seq) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		seq[i], seq[j] = seq[j], seq[i]
	}
	return seq
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
