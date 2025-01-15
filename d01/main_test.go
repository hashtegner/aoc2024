package main

import "testing"

var sample = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSum(t *testing.T) {
	left, right := parse(sample)
	sum := sum(left, right)

	if sum != 11 {
		t.Errorf("Expected 11, got %d", sum)
	}
}

func TestScore(t *testing.T) {
	left, right := parse(sample)
	score := score(left, right)

	if score != 31 {
		t.Errorf("Expected 31, got %d", score)
	}
}
