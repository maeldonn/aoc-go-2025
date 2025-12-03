package day3

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 357, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 3121910778619, got)
}
