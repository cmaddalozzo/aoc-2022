package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	start int
	end   int
}

func (a assignment) Equals(b assignment) bool {
	return a.start == b.start && a.end == b.end
}

func (a assignment) Contains(b assignment) bool {
	if a.Equals(b) {
		return true
	}
	return a.start <= b.start && a.end >= b.end
}

func (a assignment) OverlapsWith(b assignment) bool {
	if a.start >= b.start {
		return b.end >= a.start
	}
	if b.start >= a.start {
		return a.end >= b.start
	}
	return false
}

func buildAssignment(part string) *assignment {
	a := assignment{}
	bounds := strings.Split(part, "-")
	a.start, _ = strconv.Atoi(bounds[0])
	a.end, _ = strconv.Atoi(bounds[1])
	return &a
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
	numOverlapping := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		first := buildAssignment(parts[0])
		second := buildAssignment(parts[1])
		fmt.Printf("%d -> %d, %d -> %d\n", first.start, first.end, second.start, second.end)
		if first.OverlapsWith(*second) {
			numOverlapping += 1
			fmt.Printf("\t%s overlaps with %s\n", parts[0], parts[1])
		}
	}
	fmt.Printf("Num overlapping: %d\n", numOverlapping)
}
