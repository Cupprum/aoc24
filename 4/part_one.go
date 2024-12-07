package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := strings.Split(string(f), "\n")

	s := 0

	for _, l := range in {
		// Horizontal
		s += strings.Count(l, "XMAS")
		// Horizontal backwards
		s += strings.Count(l, "SAMX")
	}

	for c := 0; c < len(in[0]); c++ {
		col := ""
		for r := 0; r < len(in); r++ {
			col += string(in[r][c])
		}
		// Vertical
		s += strings.Count(col, "XMAS")
		// Vertical backwards
		s += strings.Count(col, "SAMX")
	}

	// Diagonal

	// Left bottom triangle
	// XMAS doesnt fit in last three rows
	for r := 0; r < len(in)-3; r++ {
		d := ""

		for i := 0; i < len(in)-r; i++ {
			d += string(in[r+i][i])
		}
		// Diagonal
		s += strings.Count(d, "XMAS")
		// Diagonal backwards
		s += strings.Count(d, "SAMX")
	}

	// Right top triangle
	// XMAS doesnt fit in last three columns
	for c := 1; c < len(in[0])-3; c++ {
		d := ""

		for i := 0; i < len(in[0])-c; i++ {
			d += string(in[i][c+i])
		}
		// Diagonal
		s += strings.Count(d, "XMAS")
		// Diagonal backwards
		s += strings.Count(d, "SAMX")
	}

	// Diagonal reversed
	// Right bottom triangle
	// XMAS doesnt fit in last three rows
	for r := 0; r < len(in)-3; r++ {
		d := ""

		for i := 0; i < len(in)-r; i++ {
			d += string(in[r+i][len(in)-i-1])
		}
		// Diagonal
		s += strings.Count(d, "XMAS")
		// Diagonal backwards
		s += strings.Count(d, "SAMX")
	}

	// Left top triangle
	// XMAS doesnt fit in last three columns
	for c := 3; c < len(in[0])-1; c++ {
		d := ""

		for i := 0; i <= c; i++ {
			d += string(in[i][c-i])
		}
		// Diagonal
		s += strings.Count(d, "XMAS")
		// Diagonal backwards
		s += strings.Count(d, "SAMX")
	}

	fmt.Println(s)
}
