package day2

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var input = []string{
	"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
}

func TestPartOne(t *testing.T) {
	t.Parallel()

	got := PartOne(input)
	assert.Equal(t, 1227775554, got)
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	got := PartTwo(input)
	assert.Equal(t, 4174379265, got)
}
