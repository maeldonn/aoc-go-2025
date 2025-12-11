package day10

import (
	"container/list"
	"strings"
)

type Node struct {
	state string
	dist  int
}

func (n Node) NewState(button string) string {
	newLights := []rune(n.state)
	for pos := range strings.SplitSeq(button, ",") {
		if pos == "" {
			continue
		}

		pos := int(pos[0] - '0')

		state := '#'
		if newLights[pos] == '#' {
			state = '.'
		}

		newLights[pos] = state
	}

	return string(newLights)
}

func PartOne(input []string) any {
	var presses int
	for _, line := range input {
		split := strings.Split(line, " ")
		lights := split[0][1 : len(split[0])-1]
		start := strings.Repeat(".", len(lights))

		var buttons []string
		for _, button := range split[1 : len(split)-1] {
			button = button[1 : len(button)-1]
			buttons = append(buttons, button)
		}

		visited := make(map[string]bool)

		q := list.New()
		q.PushBack(Node{start, 0})
		visited[start] = true

		for q.Len() > 0 {
			current := q.Front().Value.(Node)
			q.Remove(q.Front())

			if current.state == lights {
				presses += current.dist
				break
			}

			for _, button := range buttons {
				newState := current.NewState(button)

				if visited[newState] {
					continue
				}
				visited[newState] = true

				neighbor := Node{newState, current.dist + 1}
				q.PushBack(neighbor)
			}
		}
	}

	return presses
}

func PartTwo(input []string) any {
	return 0
}
