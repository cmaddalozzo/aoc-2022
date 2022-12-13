package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	register := 1
	i := 0
	cost := 0
	var ins instruction
	for ticks := 0; ticks < 240; ticks += 1 {
		pixel := ticks % period
		if math.Abs(float64(pixel-register)) <= 1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		if (ticks+1)%period == 0 {
			fmt.Printf("\n")
		}
		if cost > 0 {
			cost -= 1
			if cost == 0 {
				operand, _ := strconv.Atoi(ins.args[0])
				register += operand
			}
		} else {
			if instructions[i].op != "noop" {
				ins = instructions[i]
				cost = 1
			}
			i += 1
		}
	}
}
