package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
)

type coordinate struct {
	x int
	y int
}

type tree struct {
	coordinate
	height     int
	viewLeft   int
	viewRight  int
	viewTop    int
	viewBottom int
}

func (t tree) view() int {
	return t.viewLeft * t.viewRight * t.viewTop * t.viewBottom
}

func calcView(trees *[][]tree, t *tree) {
	left := t.x - 1
	right := t.x + 1
	top := t.y - 1
	bottom := t.y + 1
	viewLeft := 1
	for left > 0 && t.height > (*trees)[t.y][left].height {
		viewLeft += 1
		left -= 1
	}
	if t.x > 0 {
		t.viewLeft = viewLeft
	}
	viewRight := 1
	for right < len((*trees)[t.y])-1 && t.height > (*trees)[t.y][right].height {
		viewRight += 1
		right += 1
	}
	if t.x < len((*trees)[t.y])-1 {
		t.viewRight = viewRight
	}
	viewTop := 1
	for top > 0 && t.height > (*trees)[top][t.x].height {
		viewTop += 1
		top -= 1
	}
	if t.y > 0 {
		t.viewTop = viewTop
	}
	viewBottom := 1
	for bottom < len((*trees)[t.y])-1 && t.height > (*trees)[bottom][t.x].height {
		viewBottom += 1
		bottom += 1
	}
	if t.y < len((*trees))-1 {
		t.viewBottom = viewBottom
	}
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
		for j, c := range line {
			height, _ := strconv.Atoi(string(c))
			trees[i] = append(trees[i], tree{height: height, coordinate: coordinate{x: j, y: i}})
		}
		i += 1
	}
	for i := range trees {
		for j := range trees[i] {
			calcView(&trees, &trees[i][j])
		}
	}
	bestView := 0
	for _, row := range trees {
		for _, t := range row {
			view := t.view()
			if view > bestView {
				bestView = view
			}
			fmt.Printf("%d ", view)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Best view: %d\n", bestView)
}
