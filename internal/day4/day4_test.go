package day4

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 13, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 43, got)
}
