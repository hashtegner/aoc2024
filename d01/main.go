package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var reg = regexp.MustCompile(`\s+`)

func parse(str string) (left []int, right []int) {
	lines := strings.Split(str, "\n")

	left = make([]int, len(lines))
	right = make([]int, len(lines))

	for i, line := range lines {
		splitted := reg.Split(line, 2)
		if len(splitted) != 2 {
			continue
		}

		a, _ := strconv.Atoi(splitted[0])
		b, _ := strconv.Atoi(splitted[1])

		left[i] = a
		right[i] = b
	}

	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	var sum int
	for i := 0; i < len(left); i++ {
		sum += abs(right[i] - left[i])
	}

	return sum
}

func frequencies(list []int) (freq map[int]int) {
	freq = make(map[int]int)

	for _, v := range list {
		freq[v]++
	}

	return
}

func score(left []int, right []int) (score int) {
	freq := frequencies(right)

	for i := 0; i < len(left); i++ {
		n := left[i]
		f, ok := freq[n]
		if !ok {
			continue
		}

		score += f * n
	}

	return
}

func main() {
	left, right := parse(input)

	fmt.Println("sum", sum(left, right))
	fmt.Println("score", score(left, right))
}
