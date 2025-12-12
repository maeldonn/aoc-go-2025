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
	points := make([]image.Point, 0, len(input))
	for _, line := range input {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, image.Pt(x, y))
	}

	var maxi int
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			p3, p4, p5, p6 := image.Pt(min(p1.X, p2.X), min(p1.Y, p2.Y)), image.Pt(min(p1.X, p2.X), max(p1.Y, p2.Y)), image.Pt(max(p1.X, p2.X), min(p1.Y, p2.Y)), image.Pt(max(p1.X, p2.X), max(p1.Y, p2.Y))

			for _, p := range []image.Point{p3, p4, p5, p6} {
				if p == p1 || p == p2 {
					continue
				}

				if !PointInPolygon(p, points) {
					continue
				}
			}

			// scanner et si un point sur un des segments alors c'est pas ok

			dx := math.Abs(float64(p2.X-p1.X)) + 1
			dy := math.Abs(float64(p1.Y-p2.Y)) + 1

			if area := int(dx * dy); area > maxi {
				maxi = area
			}
		}
	}

	return maxi
}

func PointInPolygon(p image.Point, poly []image.Point) bool {
	inside := false
	j := len(poly) - 1

	for i := range poly {
		pi := poly[i]
		pj := poly[j]

		intersect := ((pi.Y > p.Y) != (pj.Y > p.Y)) && (p.X < (pj.X-pi.X)*(p.Y-pi.Y)/(pj.Y-pi.Y)+pi.X)
		if intersect {
			inside = !inside
		}

		j = i
	}

	return inside
}
