package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type slice struct{ sort.IntSlice }

func (s *slice) Pop() interface{} {
	val := 0
	if len(s.IntSlice) > 0 {
		lastIndex := len(s.IntSlice) - 1
		val = s.IntSlice[lastIndex]
		(*s).IntSlice = s.IntSlice[:lastIndex]
	}
	return val
}

func (s *slice) Push(x interface{}) {
	(*s).IntSlice = append((*s).IntSlice, x.(int))
}

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var accum int = 0
	topK := &slice{sort.IntSlice{0, 0, 0}}
	heap.Init(topK)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if accum > topK.IntSlice[0] {
				heap.Pop(topK)
				heap.Push(topK, accum)
			}
			accum = 0
		} else {
			curr, _ := strconv.Atoi(text)
			accum += curr
		}
	}
	if accum > topK.IntSlice[0] {
		heap.Pop(topK)
		heap.Push(topK, accum)
	}
	total := 0
	for _, val := range topK.IntSlice {
		fmt.Printf("Val: %d\n", val)
		total += val
	}
	fmt.Printf("Max sum is: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
