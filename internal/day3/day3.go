package day3

import "math"

func PartOne(input []string) any {
	var total int

	for _, bank := range input {
		var first, second rune
		for i, b := range bank {
			if b > first && i != len(bank)-1 {
				first, second = b, '1'
			} else if b > second {
				second = b
			}
		}

		total += int(first-'0')*10 + int(second-'0')
	}

	return total
}

func PartTwo(input []string) any {
	var total int

	for _, bank := range input {
		var curr int

		for pos := 11; pos >= 0; pos-- {
			maxIdx := curr

			for i := curr; i < len(bank)-pos; i++ {
				if bank[i] > bank[maxIdx] {
					maxIdx = i
				}
			}

			curr = maxIdx + 1
			total += int(bank[maxIdx]-'0') * int(math.Pow10(pos))
		}
	}

	return total
}
