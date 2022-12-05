package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	stackRows := []string{}
	moves := []string{}
	stacksDone := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			stacksDone = true
			continue
		}
		if stacksDone {
			moves = append(moves, line)
		} else {
			stackRows = append(stackRows, line)
		}
	}
	// Last stack row contains names
	stackNames := strings.Split(strings.TrimSpace(stackRows[len(stackRows)-1]), "   ")
	stackRows = stackRows[:len(stackRows)-1]
	stacks := make(map[string]*list.List, len(stackNames))
	stackPosToName := make([]string, len(stackNames))
	for i, name := range stackNames {
		stackPosToName[i] = name
		stacks[name] = list.New()
	}
	for _, stackRow := range stackRows {
		fmt.Printf("Processing %s\n", stackRow)
		stackIndex := 0
		for i := 1; i < len(stackRow); i += 4 {
			if stackRow[i] != ' ' {
				stackName := stackPosToName[stackIndex]
				stacks[stackName].PushBack(stackRow[i])
				fmt.Printf("Adding %c to stack %s\n", stackRow[i], stackName)
			}
			stackIndex += 1
		}
	}
	for _, stackName := range stackPosToName {
		fmt.Printf("Stack %s length: %d\n", stackName, stacks[stackName].Len())
	}
	for _, move := range moves {
		parts := strings.Split(move, " ")
		count, _ := strconv.Atoi(parts[1])
		from := parts[3]
		to := parts[5]
		fmt.Printf("Should move %d items from queue %s to queue %s\n", count, from, to)
		for i := 0; i < count; i++ {
			val := stacks[from].Front()
			fmt.Printf("Move %c to %s\n", val.Value, to)
			stacks[to].PushFront(val.Value)
			stacks[from].Remove(val)
		}
	}
	for _, stackName := range stackNames {
		fmt.Printf("%c", stacks[stackName].Front().Value)
	}
}
