package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var np = []string{"789", "456", "123", " 0A"}
var dp = []string{" ^A", "<v>"}

type Point struct {
	r int
	c int
}

func findChar(grid []string, target string) Point {
	for r, rr := range grid {
		for c, cc := range strings.Split(rr, "") {
			if cc == target {
				return Point{r, c}
			}
		}
	}
	return Point{-1, -1}
}

func findPaths(grid []string, from string, to string) string {
	start := findChar(grid, from)
	end := findChar(grid, to)
	var bestPath string
	minChanges := -1

	var dfs func(p Point, path string)
	dfs = func(p Point, path string) {
		if p == end {
			newPath := path + "A"
			changes := 0
			for i := 0; i < len(newPath)-1; i++ {
				if newPath[i] != newPath[i+1] {
					changes++
				}
			}
			if minChanges == -1 || changes < minChanges {
				minChanges = changes
				bestPath = newPath
			}
			return
		}

		if end.c < p.c && p.c > 0 && grid[p.r][p.c-1] != ' ' {
			dfs(Point{p.r, p.c - 1}, path+"<")
		}
		if end.r < p.r && p.r > 0 && grid[p.r-1][p.c] != ' ' {
			dfs(Point{p.r - 1, p.c}, path+"^")
		}
		if end.r > p.r && p.r < len(grid)-1 && grid[p.r+1][p.c] != ' ' {
			dfs(Point{p.r + 1, p.c}, path+"v")
		}
		if end.c > p.c && p.c < len(grid[p.r])-1 && grid[p.r][p.c+1] != ' ' {
			dfs(Point{p.r, p.c + 1}, path+">")
		}
	}

	dfs(start, "")
	return bestPath
}

type CK struct {
	s string
	l int
}

type CV struct {
	path   string
	length int
}

var cacheOne = make(map[CK]CV)

func solveOne(s string, l int) int {
	key := CK{s, l}
	if val, ok := cacheOne[key]; ok {
		return val.length
	}

	if l > 2 {
		return len(s)
	}

	grid := dp
	if l == 0 {
		grid = np
	}

	sum := 0

	pts := strings.Split("A"+s, "")
	for i, c := range pts {
		if i == 0 {
			continue
		}
		ll := pts[i-1]

		nextPath := findPaths(grid, ll, c)
		sum += solveOne(nextPath, l+1)
	}

	cacheOne[key] = CV{path: s, length: sum}
	return sum
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	cds := strings.Split(string(f), "\n")

	s := 0
	for _, cd := range cds {
		n, err := strconv.Atoi(strings.TrimSuffix(cd, "A"))
		if err != nil {
			panic(err)
		}

		s += solveOne(cd, 0) * n
	}
	fmt.Println(s)
}
