package day12

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"0:",
	"###",
	"##.",
	"##.",
	"",
	"1:",
	"###",
	"##.",
	".##",
	"",
	"2:",
	".##",
	"###",
	"##.",
	"",
	"3:",
	"##.",
	"###",
	"##.",
	"",
	"4:",
	"###",
	"#..",
	"###",
	"",
	"5:",
	"###",
	".#.",
	"###",
	"",
	"4x4: 0 0 0 0 2 0",
	"12x5: 1 0 1 0 2 2",
	"12x5: 1 0 1 0 3 2",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 0, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 0, got)
}
