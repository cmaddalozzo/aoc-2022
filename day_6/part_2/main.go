package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	_ "strconv"
	_ "strings"
)

func findPacketMarker(message string) int {
	marker_len := 14
	char_positions := map[rune]int{}
	unique_start := 0
	for i, c := range message {
		pos, exists := char_positions[c]
		if exists && pos >= unique_start {
			unique_start = pos + 1
		}
		char_positions[c] = i
		if (i-unique_start)+1 == marker_len {
			return i + 1
		}
	}
	return -1
}

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		marker := findPacketMarker(scanner.Text())
		fmt.Printf("Packet marker: %d\n", marker)
	}
}
