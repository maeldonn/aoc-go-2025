package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/maeldonn/aoc-go-2025/internal/day1"
	"github.com/maeldonn/aoc-go-2025/internal/day10"
	"github.com/maeldonn/aoc-go-2025/internal/day2"
	"github.com/maeldonn/aoc-go-2025/internal/day3"
	"github.com/maeldonn/aoc-go-2025/internal/day4"
	"github.com/maeldonn/aoc-go-2025/internal/day5"
	"github.com/maeldonn/aoc-go-2025/internal/day6"
	"github.com/maeldonn/aoc-go-2025/internal/day7"
	"github.com/maeldonn/aoc-go-2025/internal/day8"
	"github.com/maeldonn/aoc-go-2025/internal/day9"
	aocgoclient "github.com/maeldonn/aoc-go-client"
)

type DaySolver struct {
	PartOne func([]string) any
	PartTwo func([]string) any
}

var days = []DaySolver{
	{},
	{day1.PartOne, day1.PartTwo},
	{day2.PartOne, day2.PartTwo},
	{day3.PartOne, day3.PartTwo},
	{day4.PartOne, day4.PartTwo},
	{day5.PartOne, day5.PartTwo},
	{day6.PartOne, day6.PartTwo},
	{day7.PartOne, day7.PartTwo},
	{day8.PartOne, day8.PartTwo},
	{day9.PartOne, day9.PartTwo},
	{day10.PartOne, day10.PartTwo},
}

func main() {
	dayStr, ok := os.LookupEnv("DAY")
	if !ok {
		panic("DAY env var not set")
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		panic(err)
	}

	if day < 1 || day >= len(days) {
		panic(fmt.Sprintf("invalid day %d for year 2025", day))
	}

	client, err := aocgoclient.NewClient()
	if err != nil {
		panic(err)
	}

	input, err := client.GetInput(2025, day)
	if err != nil {
		panic(err)
	}

	solver := days[day]
	if solver.PartOne == nil || solver.PartTwo == nil {
		panic(fmt.Sprintf("no solver registered for day %d", day))
	}

	fmt.Printf("########## Day %d ##########\n", day)

	start1 := time.Now()
	part1 := solver.PartOne(input)
	fmt.Printf("Solution of part 1: %v (took %s)\n", part1, time.Since(start1))

	start2 := time.Now()
	part2 := solver.PartTwo(input)
	fmt.Printf("Solution of part 2: %v (took %s)\n", part2, time.Since(start2))
}
