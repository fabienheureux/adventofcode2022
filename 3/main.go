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

func findBadge(group [][]string) string {
	newGroup := [][]string{}
	// Condition d'arrêt, lorsque la liste ne contient plus qu'un badge, on le retourne
	if len(group) == 1 {
		return group[0][0]
	}

	// On itère sur les elfes du groupe (3)
	for i := 0; i < len(group)-1; i++ {
		elve := group[i]
		newGroup = append(newGroup, []string{})

		// On itère sur le sac à d
		for j := 0; j < len(elve); j++ {
			if find(elve[j], group[i+1]) {
				newGroup[i] = append(newGroup[i], elve[j])
			}
		}
	}

	return findBadge(newGroup)
}

func process(rounds []string) int {
	// Initialisation des variables
	groups := [][][]string{}
	group := [][]string{}
	badges := []string{}

	// Consitution des groupes
	for i := 0; i < len(rounds); i++ {
		group = append(group, strings.Split(rounds[i], ""))

		if i%3 == 2 {
			groups = append(groups, group)
			group = [][]string{}
		}
	}

	// Pour chaque groupe, détermination du badge
	for i := 0; i < len(groups); i++ {
		group = groups[i]
		badges = append(badges, findBadge(group))
	}

	return computePrioritiesSum(badges)
}

func main() {
	dat, err := os.ReadFile("./input")
	res := process(strings.Split(string(dat), "\n"))
	fmt.Println(res)

	check(err)
}
