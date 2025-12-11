package day10

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 7, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 0, got)
}
