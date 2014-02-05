package sudoku

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
  "fmt"
)

func readPuzzle(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	bf := bufio.NewReader(f)

	puzzles := make([]string, 0)
	for {
		line, isPrefix, err := bf.ReadLine()

		// End of file so leave loop
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if isPrefix {
			log.Fatal("Error: Line is unexpectedly too long to read.", f.Name())
		}

		puzzles = append(puzzles, string(line))
	}
	return puzzles
}

func benchmarkSolve(puzzle string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Println(Display(Solve(puzzle)))
	}
}

func BenchmarkSolveHardest(b *testing.B) {
	puzzles := readPuzzle("test_files/hardest.txt")
	for _, puzzle := range puzzles {
		benchmarkSolve(puzzle, b)
	}
}

func BenchmarkSolveTop95(b *testing.B) {
	puzzles := readPuzzle("test_files/top95.txt")
	for _, puzzle := range puzzles {
		benchmarkSolve(puzzle, b)
	}
}
