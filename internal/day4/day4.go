package day4

import "image"

func PartOne(input []string) any {
	var total int

	directions := []image.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for y, line := range input {
		for x, c := range line {
			if c != '@' {
				continue
			}

			var rolls int
			for _, dir := range directions {
				next := image.Pt(x, y).Add(dir)

				if next.Y < 0 || next.Y >= len(input) || next.X < 0 || next.X >= len(input[y]) {
					continue
				}

				if input[next.Y][next.X] == '@' {
					rolls++
				}
			}

			if rolls < 4 {
				total++
			}
		}
	}

	return total
}

func PartTwo(input []string) any {
	var total int

	directions := []image.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	seen := make(map[image.Point]struct{})
	visit := func() {
		for y, line := range input {
			for x, c := range line {
				if c != '@' {
					continue
				}

				curr := image.Pt(x, y)
				if _, ok := seen[curr]; ok {
					continue
				}

				var rolls int
				for _, dir := range directions {
					next := curr.Add(dir)

					if next.Y < 0 || next.Y >= len(input) || next.X < 0 || next.X >= len(input[y]) {
						continue
					}

					if _, ok := seen[next]; ok {
						continue
					}

					if input[next.Y][next.X] == '@' {
						rolls++
					}
				}

				if rolls < 4 {
					seen[curr] = struct{}{}
					total++
				}
			}
		}
	}

	for prev := -1; prev != total; {
		prev = total
		visit()
	}

	return total
}
