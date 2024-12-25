package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func find(m map[string][]string, a []string) []string {
	for _, i := range a {
		for _, j := range a {
			if i == j {
				continue
			}
			if !slices.Contains(m[j], i) {
				b := append([]string{}, a[:slices.Index(a, j)]...)
				b = append(b, a[slices.Index(a, j)+1:]...)
				return find(m, b)
			}
		}
	}

	sort.Strings(a)
	return a
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	prs := strings.Split(string(f), "\n")

	m := map[string][]string{}
	for _, pr := range prs {
		cs := strings.Split(pr, "-")
		c1 := cs[0]
		c2 := cs[1]

		if _, ok := m[c1]; !ok {
			m[c1] = []string{c2}
		} else {
			m[c1] = append(m[c1], c2)
		}

		if _, ok := m[c2]; !ok {
			m[c2] = []string{c1}
		} else {
			m[c2] = append(m[c2], c1)
		}
	}

	ml := 0
	mv := ""

	for k, v := range m {
		t := append([]string{k}, v...)

		a := find(m, t)
		if len(a) > ml {
			ml = len(a)
			mv = strings.Join(a, ",")
		}
	}
	fmt.Println(mv)
}
