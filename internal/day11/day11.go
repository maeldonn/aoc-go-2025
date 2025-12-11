package day11

import "strings"

func PartOne(input []string) any {
	paths := make(map[string][]string)
	for _, line := range input {
		devices := strings.Split(line, " ")
		for _, device := range devices[1:] {
			current := devices[0][:len(devices[0])-1]
			paths[current] = append(paths[current], device)
		}
	}

	visited := make(map[string]bool, len(paths))

	var visit func(curr string) int
	visit = func(curr string) int {
		if curr == "out" {
			return 1
		}

		visited[curr] = true
		defer func() {
			visited[curr] = false
		}()

		var total int
		for _, device := range paths[curr] {
			if visited[device] {
				continue
			}

			total += visit(device)
		}

		return total
	}

	return visit("you")
}

func PartTwo(input []string) any {
	paths := make(map[string][]string)
	for _, line := range input {
		devices := strings.Split(line, " ")
		for _, device := range devices[1:] {
			current := devices[0][:len(devices[0])-1]
			paths[current] = append(paths[current], device)
		}
	}

	type cacheKey struct {
		node   string
		hasDac bool
		hasFft bool
	}

	var (
		cache   = make(map[cacheKey]int)
		visited = make(map[string]bool, len(paths))
	)

	var visit func(curr string) int
	visit = func(curr string) int {
		key := cacheKey{curr, visited["dac"], visited["fft"]}
		if val, ok := cache[key]; ok {
			return val
		}

		if curr == "out" {
			if key.hasDac && key.hasFft {
				return 1
			}
			return 0
		}

		visited[curr] = true
		defer func() {
			visited[curr] = false
		}()

		var total int
		for _, device := range paths[curr] {
			if visited[device] {
				continue
			}

			total += visit(device)
		}

		cache[key] = total
		return total
	}

	return visit("svr")
}
