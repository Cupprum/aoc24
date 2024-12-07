package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func partOne() {
	i := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	ms := re.FindAllString(i, -1)

	s := 0

	for _, m := range ms {
		m = strings.TrimPrefix(m, "mul(")
		m = strings.TrimSuffix(m, ")")
		vs := strings.Split(m, ",")
		l, err := strconv.Atoi(vs[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(vs[1])
		if err != nil {
			panic(err)
		}
		s += l * r
	}

	fmt.Println(s)
}
