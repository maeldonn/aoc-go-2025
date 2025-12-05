package day5

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) any {
	var (
		ranges [][2]int
		ids    []int
	)

	for _, line := range input {
		split := strings.Split(line, "-")
		if len(split) == 2 {
			min, _ := strconv.Atoi(split[0])
			max, _ := strconv.Atoi(split[1])
			ranges = append(ranges, [2]int{min, max})
		}

		id, _ := strconv.Atoi(line)
		ids = append(ids, id)
	}

	var total int
	for _, id := range ids {
		for _, r := range ranges {
			if r[0] <= id && id <= r[1] {
				total++
				break
			}
		}
	}

	return total
}

func PartTwo(input []string) any {
	var ranges [][2]int
	for _, line := range input {
		if len(line) == 0 {
			break
		}

		split := strings.Split(line, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])
		ranges = append(ranges, [2]int{min, max})
	}

	slices.SortFunc(ranges, func(a, b [2]int) int {
		minA, minB, maxA, maxB := a[0], b[0], a[1], b[1]

		if minA == minB {
			return maxA - maxB
		}
		return minA - minB
	})

	total, lastMax := 0, -1
	for _, r := range ranges {
		min, max := r[0], r[1]

		if lastMax >= max {
			continue
		}

		if lastMax >= min {
			min = lastMax + 1
		}

		total += max - min + 1
		lastMax = max
	}

	return total
}
