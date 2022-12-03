package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	_ "strconv"
	_ "strings"
)

func getPriority(char rune) int {
	num := int(char)
	if num <= int('Z') {
		return (num - int('A')) + 1 + 26
	} else {
		return (num - int('a')) + 1
	}
}

func getBadgeItem(group []string) rune {
	seen := [2]map[rune]int{}
	var badge rune
	for i := 0; i < 2; i++ {
		bag := group[i]
		seen[i] = make(map[rune]int)
		for j := 0; j < int(len(bag)); j++ {
			c := rune(bag[j])
			seen[i][c] += 1
		}
	}
	for _, c := range group[2] {
		_, in_first := seen[0][c]
		_, in_second := seen[1][c]
		if in_first && in_second {
			return c
		}
	}
	return badge
}

func main() {

	//seen := map[string]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sum := 0
	var group []string
	count := 0
	for scanner.Scan() {
		bag := scanner.Text()
		group = append(group, bag)
		count += 1
		if count%3 == 0 {
			fmt.Printf("Length group: %d\n", len(group))
			badge := getBadgeItem(group)
			fmt.Printf("Badge: %c\n", badge)
			sum += getPriority(badge)
			group = []string{}
		}
	}
	fmt.Printf("Priority: %d\n", sum)
}
