package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type P struct {
	r int
	c int
	o string
}

type I struct {
	p  P
	sc int
	i  int
}

type PriorityQueue []*I

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].sc < pq[j].sc
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].i = i
	pq[j].i = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*I)
	item.i = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	item.i = -1    // for safety
	*pq = old[0 : n-1]
	return item
}

func findChar(m [][]string, ch string) P {
	for r, rr := range m {
		for c, cc := range rr {
			if cc == ch {
				return P{r, c, ">"}
			}
		}
	}

	return P{-1, -1, ""}
}

func partOne() {
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

	hist := map[P]*I{}

	ms := -1

	for pq.Len() > 0 {
		e := heap.Pop(&pq).(*I)

		if v, ok := hist[e.p]; ok && v.sc < e.sc {
			continue
		}

		if e.p.r == end.r && e.p.c == end.c {
			ms = e.sc
			break
		}

		if e.p.o == "<" {
			if m[e.p.r][e.p.c-1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c - 1, "<"}, sc: e.sc + 1})
			}
			if m[e.p.r-1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "^"}, sc: e.sc + 1000})
			}
			if m[e.p.r][e.p.c+1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, ">"}, sc: e.sc + 2000})
			}
			if m[e.p.r+1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "V"}, sc: e.sc + 1000})
			}
		} else if e.p.o == "^" {
			if m[e.p.r-1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r - 1, e.p.c, "^"}, sc: e.sc + 1})
			}
			if m[e.p.r][e.p.c+1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, ">"}, sc: e.sc + 1000})
			}
			if m[e.p.r+1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "V"}, sc: e.sc + 2000})
			}
			if m[e.p.r][e.p.c-1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "<"}, sc: e.sc + 1000})
			}
		} else if e.p.o == ">" {
			if m[e.p.r][e.p.c+1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c + 1, ">"}, sc: e.sc + 1})
			}
			if m[e.p.r+1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "V"}, sc: e.sc + 1000})
			}
			if m[e.p.r][e.p.c-1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "<"}, sc: e.sc + 2000})
			}
			if m[e.p.r-1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "^"}, sc: e.sc + 1000})
			}
		} else if e.p.o == "V" {
			if m[e.p.r+1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r + 1, e.p.c, "V"}, sc: e.sc + 1})
			}
			if m[e.p.r][e.p.c-1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "<"}, sc: e.sc + 1000})
			}
			if m[e.p.r-1][e.p.c] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, "^"}, sc: e.sc + 2000})
			}
			if m[e.p.r][e.p.c+1] != "#" {
				heap.Push(&pq, &I{p: P{e.p.r, e.p.c, ">"}, sc: e.sc + 1000})
			}
		}

		if v, ok := hist[e.p]; !ok || (ok && v.sc > e.sc) {
			hist[e.p] = e
		}
	}

	fmt.Println(ms)
}
