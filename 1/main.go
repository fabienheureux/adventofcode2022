package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getMax(elves []string) int {
	var sums []int
	for i := 0; i < len(elves); i++ {
		sum := 0
		elve := strings.Split(elves[i], "\n")

		for i := 0; i < len(elve); i++ {
			value, err := strconv.Atoi(elve[i])

			if err == nil {
				sum += value
			}
		}

		if len(sums) > 3 {
			sums[3] = sum
		} else {
			sums = append(sums, sum)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	}

	topThreeSum := 0

	for i := 0; i < 3; i++ {
		topThreeSum += sums[i]
	}

	return topThreeSum
}

func main() {
	dat, err := os.ReadFile("./input")
	fmt.Print(getMax(strings.Split(string(dat), "\n\n")))

	if err != nil {
		log.Fatal(err)
	}
}
