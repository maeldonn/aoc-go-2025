package day7

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 21, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 40, got)
}
