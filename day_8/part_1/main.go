package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
)

type tree struct {
	height        int
	visibleLeft   bool
	visibleRight  bool
	visibleTop    bool
	visibleBottom bool
}

func (t tree) visible() bool {
	return t.visibleRight || t.visibleLeft || t.visibleTop || t.visibleBottom
}

func main() {

	trees := [][]tree{}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	i := 0
	for scanner.Scan() {
		trees = append(trees, []tree{})
		line := scanner.Text()
		for _, c := range line {
			height, _ := strconv.Atoi(string(c))
			trees[i] = append(trees[i], tree{height: height})
		}
		i += 1
	}
	for i, row := range trees {
		left := 1
		right := len(trees[i]) - 2
		leftHighWater := row[0].height
		rightHighWater := row[len(trees[i])-1].height
		trees[i][0].visibleLeft = true
		trees[i][len(trees[i])-1].visibleRight = true
		for left < len(trees[i]) {
			leftTree := &trees[i][left]
			if leftTree.height <= leftHighWater {
				leftTree.visibleLeft = false
			} else {
				leftTree.visibleLeft = true
				leftHighWater = leftTree.height
			}
			if trees[i][right].height <= rightHighWater {
				trees[i][right].visibleRight = false
			} else {
				trees[i][right].visibleRight = true
				rightHighWater = trees[i][right].height
			}
			left += 1
			right -= 1
		}
	}
	for i := range trees[0] {
		top := 1
		bottom := len(trees) - 2
		topHighWater := trees[0][i].height
		bottomHighWater := trees[len(trees)-1][i].height
		trees[0][i].visibleTop = true
		trees[len(trees)-1][i].visibleBottom = true
		for top < len(trees) {
			topTree := &trees[top][i]
			if topTree.height <= topHighWater {
				topTree.visibleTop = false
			} else {
				topTree.visibleTop = true
				topHighWater = topTree.height
			}
			bottomTree := &trees[bottom][i]
			if bottomTree.height <= bottomHighWater {
				bottomTree.visibleBottom = false
			} else {
				bottomTree.visibleBottom = true
				bottomHighWater = bottomTree.height
			}
			top += 1
			bottom -= 1
		}
	}
	numVisible := 0
	for _, row := range trees {
		for _, t := range row {
			out := 0
			if t.visible() {
				out = 1
				numVisible += 1
			}
			fmt.Printf("%d ", out)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Num visible: %d\n", numVisible)
}
