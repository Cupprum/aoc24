package main

import (
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

func mark(m [][]string, hist map[P][]I, cur P, st P) {
	m[cur.r][cur.c] = "O"

	if cur == st {
		return
	}

	for _, i := range hist[cur] {
		mark(m, hist, i.p, st)
	}
}

func partTwo() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lns := strings.Split(string(in), "\n")

	m := [][]string{}
	for _, ln := range lns {
		m = append(m, strings.Split(ln, ""))
	}

	st := findChar(m, "S")
	end := findChar(m, "E")

	pq := PriorityQueue{}
	heap.Init(&pq)
	pq.Push(&I{p: st, sc: 0})

	vstd := map[P]int{}
	hist := map[P][]I{}

	addHist := func(hist map[P][]I, f P, c I) {
		if v, ok := hist[f]; !ok || (ok && v[0].sc > c.sc) {
			hist[f] = []I{c}
		} else if ok && v[0].sc == c.sc && !slices.Contains(hist[f], c) {
			hist[f] = append(hist[f], c)
		}
	}

	for pq.Len() > 0 {
		e := heap.Pop(&pq).(*I)

		if v, ok := vstd[e.p]; ok && v < e.sc {
			continue
		}

		if e.p.o == "<" {
			if m[e.p.r][e.p.c-1] != "#" {
				addHist(hist, P{e.p.r, e.p.c - 1, "<"}, I{p: e.p, sc: e.sc + 1})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c - 1, "<"}, sc: e.sc + 1})
			}
			if m[e.p.r-1][e.p.c] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "^"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "^"}, sc: e.sc + 1000})
			}
			if m[e.p.r][e.p.c+1] != "#" {
				addHist(hist, P{e.p.r, e.p.c, ">"}, I{p: e.p, sc: e.sc + 2000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, ">"}, sc: e.sc + 2000})
			}
			if m[e.p.r+1][e.p.c] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "V"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "V"}, sc: e.sc + 1000})
			}
		} else if e.p.o == "^" {
			if m[e.p.r-1][e.p.c] != "#" {
				addHist(hist, P{e.p.r - 1, e.p.c, "^"}, I{p: e.p, sc: e.sc + 1})
				heap.Push(&pq, &I{p: P{e.p.r - 1, e.p.c, "^"}, sc: e.sc + 1})
			}
			if m[e.p.r][e.p.c+1] != "#" {
				addHist(hist, P{e.p.r, e.p.c, ">"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, ">"}, sc: e.sc + 1000})
			}
			if m[e.p.r+1][e.p.c] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "V"}, I{p: e.p, sc: e.sc + 2000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "V"}, sc: e.sc + 2000})
			}
			if m[e.p.r][e.p.c-1] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "<"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "<"}, sc: e.sc + 1000})
			}
		} else if e.p.o == ">" {
			if m[e.p.r][e.p.c+1] != "#" {
				addHist(hist, P{e.p.r, e.p.c + 1, ">"}, I{p: e.p, sc: e.sc + 1})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c + 1, ">"}, sc: e.sc + 1})
			}
			if m[e.p.r+1][e.p.c] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "V"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "V"}, sc: e.sc + 1000})
			}
			if m[e.p.r][e.p.c-1] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "<"}, I{p: e.p, sc: e.sc + 2000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "<"}, sc: e.sc + 2000})
			}
			if m[e.p.r-1][e.p.c] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "^"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "^"}, sc: e.sc + 1000})
			}
		} else if e.p.o == "V" {
			if m[e.p.r+1][e.p.c] != "#" {
				addHist(hist, P{e.p.r + 1, e.p.c, "V"}, I{p: e.p, sc: e.sc + 1})
				heap.Push(&pq, &I{p: P{e.p.r + 1, e.p.c, "V"}, sc: e.sc + 1})
			}
			if m[e.p.r][e.p.c-1] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "<"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "<"}, sc: e.sc + 1000})
			}
			if m[e.p.r-1][e.p.c] != "#" {
				addHist(hist, P{e.p.r, e.p.c, "^"}, I{p: e.p, sc: e.sc + 2000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "^"}, sc: e.sc + 2000})
			}
			if m[e.p.r][e.p.c+1] != "#" {
				addHist(hist, P{e.p.r, e.p.c, ">"}, I{p: e.p, sc: e.sc + 1000})
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, ">"}, sc: e.sc + 1000})
			}
		}

		if v, ok := vstd[e.p]; !ok || (ok && v > e.sc) {
			vstd[e.p] = e.sc
		}
	}

	m[st.r][st.c] = "."
	m[end.r][end.c] = "."

	mark(m, hist, P{end.r, end.c, "^"}, st)
	s := 0
	for _, ln := range m {
		for _, c := range ln {
			if c == "O" {
				s++
			}
		}
	}
	fmt.Println(s)
}
