package day6

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 4277556, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 3263827, got)
}
