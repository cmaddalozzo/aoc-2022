package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op     string
	cycles int
	args   []string
}

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := []instruction{}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		op := parts[0]
		cycles := 1
		if op == "addx" {
			cycles = 2
		}
		instructions = append(instructions, instruction{op: op, cycles: cycles, args: parts[1:]})
	}
	period := 40
	offset := 20
	register := 1
	i := 0
	cost := 0
	var ins instruction
	total_signal := 0
	for ticks := 1; ticks <= 220; ticks += 1 {
		if (ticks-offset)%period == 0 {
			signal_strength := register * ticks
			total_signal += signal_strength
			fmt.Println("Signal at ", ticks, " is ", signal_strength)
		}
		if cost > 0 {
			cost -= 1
			if cost == 0 {
				operand, _ := strconv.Atoi(ins.args[0])
				register += operand
				fmt.Println("Adding", operand, "=", register, "at", ticks)
			}
		} else {
			if instructions[i].op != "noop" {
				ins = instructions[i]
				cost = 1
			}
			i += 1
		}
	}
	fmt.Println("Total signal", total_signal)
}
