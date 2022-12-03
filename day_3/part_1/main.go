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

func getBadItem(bag string) rune {
	seen := map[rune]int{}
	for i, c := range bag {
		if i >= int(len(bag)/2) {
			_, present := seen[c]
			if present {
				return c
			}
		} else {
			seen[c] = 1
		}
	}
	return 'a'
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
	for scanner.Scan() {
		bag := scanner.Text()
		item := getBadItem(bag)
		priority := getPriority(item)
		sum += priority
		fmt.Printf("Bad item: %c -> %d with priority %d\n", item, int(item), priority)
	}
	fmt.Printf("Priority: %d\n", sum)
}
