package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var accum int = 0
	var max int = 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if accum > max {
				max = accum
			}
			accum = 0
		} else {
			curr, _ := strconv.Atoi(text)
			accum += curr
		}
	}
	if accum > max {
		max = accum
	}
	fmt.Printf("Max sum is: %d\n", max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
