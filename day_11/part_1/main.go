package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	_ "strings"
)

type monkey struct {
	items           []int
	operation       string
	divisor         int
	target_positive int
	target_negative int
}

func parseMonkey(s string) monkey {
	m := monkey{}
	lines := strings.Split(s, "\n")
	for _, s := range strings.Split(lines[1][18:], ", ") {
		item, _ := strconv.Atoi(s)
		m.items = append(m.items, item)
	}
	m.operation = lines[2][19:]
	m.divisor, _ = strconv.Atoi(lines[3][21:])
	m.target_positive, _ = strconv.Atoi(lines[4][29:])
	m.target_negative, _ = strconv.Atoi(lines[5][30:])
	return m
}

func main() {
	monkeys := []monkey{}

	byte_content, _ := os.ReadFile(os.Args[1])
	content := string(byte_content)
	for _, txt := range strings.Split(content, "\n\n") {
		monkeys = append(monkeys, parseMonkey(strings.Trim(txt, "\n")))
	}
	fmt.Println(monkeys)
}
