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
		"B": 1,
		"C": 2,
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
		desired_outcome := parts[1]
		var my_move int
		if desired_outcome == "X" {
			// Is this loss?
			if opponent_move == 0 {
				my_move = num_moves - 1
			} else {
				my_move = opponent_move - 1
			}
		} else if desired_outcome == "Y" {
			// Draw
			my_move = opponent_move
			score += 3
		} else {
			// Win
			my_move = (opponent_move + 1) % num_moves
			score += 6
		}
		score += moves[my_move]
	}
	fmt.Printf("Final score: %d\n", score)
}
