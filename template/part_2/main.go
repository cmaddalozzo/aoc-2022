package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	_ "strconv"
	_ "strings"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		fmt.Printf("Line: %s\n", scanner.Text())
	}
}
