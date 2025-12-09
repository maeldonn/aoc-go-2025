package day9

import (
	"fmt"
	"image"
	"math"
)

func PartOne(input []string) any {
	points := make([]image.Point, 0, len(input))
	for _, line := range input {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, image.Pt(x, y))
	}

	var max int
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			dx := math.Abs(float64(p2.X-p1.X)) + 1
			dy := math.Abs(float64(p1.Y-p2.Y)) + 1

			if area := int(dx * dy); area > max {
				max = area
			}
		}
	}

	return max
}

func PartTwo(input []string) any {
	return 0
}
