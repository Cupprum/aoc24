package main

import (
	"fmt"
	"os"
	"strings"
)

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := strings.Split(string(f), "\n")

	s := 0

	for r := 1; r < len(in)-1; r++ {
		for c := 1; c < len(in[0])-1; c++ {
			if string(in[r][c]) == "A" {
				w1 := string(in[r-1][c-1]) + string(in[r][c]) + string(in[r+1][c+1])
				w2 := string(in[r-1][c+1]) + string(in[r][c]) + string(in[r+1][c-1])
				if (w1 == "SAM" || w1 == "MAS") && (w2 == "SAM" || w2 == "MAS") {
					s++
				}
			}
		}
	}

	fmt.Println(s)
}
