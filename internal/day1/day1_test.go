package day1

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
	"R1000",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 3, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 16, got)
}
