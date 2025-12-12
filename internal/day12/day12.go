package day12

import "fmt"

func PartOne(input []string) any {
	var total int
	for _, line := range input[30:] {
		var x, y, a, b, c, d, e, f int
		fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &x, &y, &a, &b, &c, &d, &e, &f)

		shapes := a + b + c + d + e + f
		possible := (x / 3) * (y / 3)

		if possible >= shapes {
			total++
		}
	}

	return total
}

func PartTwo(input []string) any {
	return 0
}
