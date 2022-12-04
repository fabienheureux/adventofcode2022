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

func getScore(play string) int {
	if play == "A" || play == "X" {
		return 1
	} else if play == "B" || play == "Y" {
		return 2
	} else if play == "C" || play == "Z" {
		return 3
	}
	return 0
}

func checkWinner(opponent string, self string) bool {
	if opponent == "A" && self == "Z" {
		return false
	} else if opponent == "C" && self == "Y" {
		return false
	} else if opponent == "B" && self == "X" {
		return false
	}

	return true
}

func checkPlay(opponent string, self string) int {
	if getScore(opponent) == getScore(self) {
		return getScore(self) + 3
	}
	if checkWinner(opponent, self) {
		return getScore(self) + 6
	}
	return getScore(self)
}

func process(rounds []string) int {
	total := 0
	for i := 0; i < len(rounds); i++ {
		round := rounds[i]
		if round != "" {
			opponent := strings.Split(round, " ")[0]
			self := strings.Split(round, " ")[1]
			total += checkPlay(opponent, self)
		}
	}

	return total
}

func main() {
	dat, err := os.ReadFile("./input")
	fmt.Print(process(strings.Split(string(dat), "\n")))

	check(err)
}
