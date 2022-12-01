package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getMax(elves []string) int {
	max := 0
	for i := 0; i < len(elves); i++ {
		sum := 0
		elve := strings.Split(elves[i], "\n")

		for i := 0; i < len(elve); i++ {
			value, err := strconv.Atoi(elve[i])

			if err == nil {
				sum += value
			}
		}
		if sum >= max {
			max = sum
		}
	}

	return max
}

func main() {
	dat, err := os.ReadFile("./input")
	fmt.Print(getMax(strings.Split(string(dat), "\n\n")))

	if err != nil {
		log.Fatal(err)
	}
}
