package day8

import (
	"flag"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type Point3 struct {
	X, Y, Z int
}

func PartOne(input []string) any {
	points := make([]Point3, 0, len(input))
	for _, line := range input {
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		points = append(points, Point3{x, y, z})
	}

	distances := make(map[int][]int)
	for i, p1 := range points {
		for j, p2 := range points[i+1:] {
			dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
			d := dx*dx + dy*dy + dz*dz
			distances[d] = []int{i, j + i + 1}
		}
	}

	keys := slices.Collect(maps.Keys(distances))
	slices.Sort(keys)

	nb := 1000
	if flag.Lookup("test.v") != nil {
		nb = 10
	}

	links := make(map[Point3][]Point3)
	for _, key := range keys[:nb] {
		distance := distances[key]
		p1, p2 := points[distance[0]], points[distance[1]]

		links[p1] = append(links[p1], p2)
		links[p2] = append(links[p2], p1)
	}

	seen := make(map[Point3]struct{})

	var visit func(point Point3) int
	visit = func(point Point3) int {
		if _, ok := seen[point]; ok {
			return 0
		}
		seen[point] = struct{}{}

		total := 1
		for _, linkedPoint := range links[point] {
			total += visit(linkedPoint)
		}

		return total
	}

	largest := make([]int, 3)
	for point := range links {
		switch total := visit(point); {
		case total >= largest[0]:
			largest[2], largest[1], largest[0] = largest[1], largest[0], total
		case total >= largest[1]:
			largest[2], largest[1] = largest[1], total
		case total >= largest[2]:
			largest[2] = total
		}
	}

	return largest[0] * largest[1] * largest[2]
}

func PartTwo(input []string) any {
	points := make([]Point3, 0, len(input))
	for _, line := range input {
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		points = append(points, Point3{x, y, z})
	}

	distances := make(map[int][]int)
	for i, p1 := range points {
		for j, p2 := range points[i+1:] {
			dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
			d := dx*dx + dy*dy + dz*dz
			distances[d] = []int{i, j + i + 1}
		}
	}

	keys := slices.Collect(maps.Keys(distances))
	slices.Sort(keys)

	links := make(map[Point3][]Point3)
	for _, key := range keys {
		distance := distances[key]
		p1, p2 := points[distance[0]], points[distance[1]]

		links[p1] = append(links[p1], p2)
		links[p2] = append(links[p2], p1)

		visited := make(map[Point3]struct{})

		var dfs func(point Point3)
		dfs = func(point Point3) {
			if _, ok := visited[point]; ok {
				return
			}
			visited[point] = struct{}{}

			for _, link := range links[point] {
				dfs(link)
			}
		}

		if dfs(p1); len(visited) == len(points) {
			return p1.X * p2.X
		}
	}

	return 0
}
