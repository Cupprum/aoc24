package main

import (
	"fmt"
	"os"
	"strconv"
)

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := string(f)

	r := []int{}
	c := 0
	for i := 0; i < len(in); i++ {
		n, err := strconv.Atoi(string(in[i]))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				r = append(r, c)
			}
			c++
		} else if i%2 == 1 {
			for j := 0; j < n; j++ {
				r = append(r, -1)
			}
		} else {
			panic("this should never happen")
		}
	}

	for i := c; i > 0; i-- {
		count := 0
		loc := -1
		for j, v := range r {
			if v == i {
				if loc == -1 {
					loc = j
				}
				count++
			}
		}

		cnt := 0
		for j := 0; j < len(r); j++ {
			if j > loc {
				break
			}

			if r[j] == -1 {
				cnt++
			} else {
				cnt = 0
			}
			if cnt == count {
				for k := 0; k < count; k++ {
					r[j-k] = i
					r[loc+k] = -1
				}
				break
			}
		}

	}

	s := 0
	for i, n := range r {
		if n == -1 {
			continue
		}
		s += i * n
	}
	fmt.Println(s)
}
