package day11

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input1 = []string{
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input1)
	assert.Equal(t, 5, got)
}

var input2 = []string{
	"svr: aaa bbb",
	"aaa: fft",
	"fft: ccc",
	"bbb: tty",
	"tty: ccc",
	"ccc: ddd eee",
	"ddd: hub",
	"hub: fff",
	"eee: dac",
	"dac: fff",
	"fff: ggg hhh",
	"ggg: out",
	"hhh: out",
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input2)
	assert.Equal(t, 2, got)
}
