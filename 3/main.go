package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find(searchString string, searchedSlice []string) bool {
	found := false

	for i := 0; i < len(searchedSlice); i++ {
		if searchString == searchedSlice[i] {
			found = true
			break
		}
	}
	return found
}

func alphaNum(letter string) int {
	alphabet := "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter)
}

func computePrioritiesSum(priorities []string) int {
	count := 0
	for i := 0; i < len(priorities); i++ {
		letter := priorities[i]
		count += alphaNum(letter)
	}

	return count
}

func getRuckSackContent(rawRuckSack string) int {
	var comp1 []string
	var comp2 []string
	var priorities []string

	items := strings.Split(rawRuckSack, "")

	for i := 0; i < len(items); i++ {
		item := items[i]

		if i < len(items)/2 {
			comp1 = append(comp1, item)
		} else {
			comp2 = append(comp2, item)
			if find(item, comp1) && !find(item, priorities) {
				priorities = append(priorities, item)
			}
		}
	}

	return computePrioritiesSum(priorities)
}

func process(rounds []string) int {
	total := 0

	for i := 0; i < len(rounds); i++ {
		if rounds[i] != "" {
			total += getRuckSackContent(rounds[i])
		}
	}

	return total
}

func main() {
	dat, err := os.ReadFile("./input")
	fmt.Println(process(strings.Split(string(dat), "\n")))

	check(err)
}
