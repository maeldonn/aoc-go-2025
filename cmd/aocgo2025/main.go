package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	aocgoclient "github.com/maeldonn/aoc-go-client"
)

func main() {
	dayStr, ok := os.LookupEnv("DAY")
	if !ok {
		panic("DAY env var not set")
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		panic(err)
	}

	client, err := aocgoclient.NewClient()
	if err != nil {
		panic(err)
	}

	input, err := client.GetInput(2025, day)
	if err != nil {
		panic(err)
	}

	var partOne, partTwo func([]string) any

	fmt.Printf("########## Day %d ##########\n", day)

	start1 := time.Now()
	part1 := partOne(input)
	fmt.Printf("Solution of part 1: %v (took %s)\n", part1, time.Since(start1))

	start2 := time.Now()
	part2 := partTwo(input)
	fmt.Printf("Solution of part 2: %v (took %s)\n", part2, time.Since(start2))
}
