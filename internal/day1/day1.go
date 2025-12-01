package day1

import "strconv"

func PartOne(input []string) any {
	dial := 50

	var password int
	for _, line := range input {
		dir, rot := line[0], line[1:]
		rotN, _ := strconv.Atoi(rot)

		switch string(dir) {
		case "L":
			dial = (dial - rotN) % 100
		case "R":
			dial = (dial + rotN) % 100
		}

		if dial == 0 {
			password++
		}
	}

	return password
}

func PartTwo(input []string) any {
	dial := 50

	var password int
	for _, line := range input {
		dir, rot := line[0], line[1:]
		rotN, _ := strconv.Atoi(rot)

		// count the number of rotations
		if r := rotN / 100; r > 0 {
			password += r
			rotN = rotN % 100
		}

		switch string(dir) {
		case "L":
			sub := dial - rotN
			if (dial > 0 && sub <= 0) || (dial < 0 && sub <= -100) {
				password++
			}

			dial = sub % 100
		case "R":
			add := dial + rotN
			if (dial < 0 && add >= 0) || (dial > 0 && add >= 100) {
				password++
			}

			dial = add % 100
		}
	}

	return password
}
