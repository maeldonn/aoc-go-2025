package day2

import (
	"strconv"
	"strings"
)

func PartOne(input []string) any {
	var invalid int

	for r := range strings.SplitSeq(input[0], ",") {
		ids := strings.Split(r, "-")

		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		for n := start; n <= end; n++ {
			str := strconv.Itoa(n)
			if string(str[:(len(str)/2)]) == string(str[(len(str)/2):]) {
				invalid += n
			}
		}

	}

	return invalid
}

func PartTwo(input []string) any {
	var invalid int

	for r := range strings.SplitSeq(input[0], ",") {
		ids := strings.Split(r, "-")

		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		compare := func(n int) {
			str := strconv.Itoa(n)

			for size := 1; size <= len(str)/2; size++ {
				if len(str)%size != 0 {
					continue
				}

				var valid = true
				for i := size; i < len(str); i += size {
					if str[i:i+size] != str[:size] {
						valid = false
						break
					}
				}

				if valid {
					invalid += n
					return
				}
			}
		}

		for n := start; n <= end; n++ {
			compare(n)
		}
	}

	return invalid
}
