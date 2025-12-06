package day6

import (
	"strconv"
	"strings"
)

func PartOne(input []string) any {
	operators := strings.Fields(input[len(input)-1])

	nums := make([]int, len(operators))
	for _, line := range input[:len(input)-1] {
		for i, f := range strings.Fields(line) {
			switch operators[i] {
			case "+":
				n, _ := strconv.Atoi(f)
				nums[i] += n
			case "*":
				if nums[i] == 0 {
					nums[i] = 1
				}

				n, _ := strconv.Atoi(f)
				nums[i] *= n
			}
		}
	}

	var total int
	for _, n := range nums {
		total += n
	}

	return total
}

func PartTwo(input []string) any {
	operators := strings.Fields(input[len(input)-1])
	curr := len(operators) - 1

	nums := make([]int, len(operators))
	for i := len(input[0]) - 1; i >= 0; i-- {
		var num []byte
		for j := range len(input) - 1 {
			if n := input[j][i]; n != 32 {
				num = append(num, n)
			}
		}

		if len(num) == 0 {
			curr--
			continue
		}

		switch operators[curr] {
		case "+":
			n, _ := strconv.Atoi(string(num))
			nums[curr] += n
		case "*":
			if nums[curr] == 0 {
				nums[curr] = 1
			}

			n, _ := strconv.Atoi(string(num))
			nums[curr] *= n
		}
	}

	var total int
	for _, n := range nums {
		total += n
	}

	return total
}
