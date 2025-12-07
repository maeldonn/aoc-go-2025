package day7

import (
	"image"
	"strings"
)

func PartOne(input []string) any {
	visited := make(map[image.Point]struct{})

	var visit func(pos image.Point)
	visit = func(pos image.Point) {
		if pos.Y >= len(input) || pos.X < 0 || pos.X >= len(input[0]) {
			return
		}

		if input[pos.Y][pos.X] != '^' {
			visit(pos.Add(image.Pt(0, 1)))
			return
		}

		if _, ok := visited[pos]; !ok {
			visited[pos] = struct{}{}
			visit(pos.Add(image.Pt(-1, 1)))
			visit(pos.Add(image.Pt(1, 1)))
		}
	}

	start := image.Pt(strings.IndexByte(input[0], 'S'), 0)
	visit(start)

	return len(visited)
}

func PartTwo(input []string) any {
	cache := make(map[image.Point]int)

	var visit func(pos image.Point) int
	visit = func(pos image.Point) int {
		if pos.Y >= len(input) || pos.X < 0 || pos.X >= len(input[0]) {
			return 1
		}

		if input[pos.Y][pos.X] != '^' {
			return visit(pos.Add(image.Pt(0, 1)))
		}

		if count, ok := cache[pos]; ok {
			return count
		}

		left := visit(pos.Add(image.Pt(-1, 1)))
		right := visit(pos.Add(image.Pt(1, 1)))

		cache[pos] = left + right
		return left + right
	}

	start := image.Pt(strings.IndexByte(input[0], 'S'), 0)
	return visit(start)
}
