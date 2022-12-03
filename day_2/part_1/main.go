package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	_ "strconv"
	"strings"
)

func main() {

	moves := [3]int{1, 2, 3}
	num_moves := 3
	aliases := map[string]int{
		"A": 0,
		"X": 0,
		"B": 1,
		"Y": 1,
		"C": 2,
		"Z": 2,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	score := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		opponent_move := aliases[parts[0]]
		my_move := aliases[parts[1]]
		if my_move == (opponent_move+1)%num_moves {
			score += 6
		} else if my_move == opponent_move {
			score += 3
		}
		score += moves[my_move]
	}
	fmt.Printf("Final score: %d\n", score)
}
