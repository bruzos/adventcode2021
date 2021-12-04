package main

import (
	"codeadvent/tools"
	"fmt"
	"strconv"
	"strings"
)

type PositionExtended struct {
	X     int
	Depth int
	Aim   int
}

func (p *PositionExtended) forward(n int) {
	p.X += n
	p.Depth += p.Aim * n
}

func (p *PositionExtended) down(n int) {
	p.Aim += n
}

func (p *PositionExtended) up(n int) {
	p.Aim -= n
}

func (p *PositionExtended) PositionValue() int {
	return p.X * p.Depth
}

func main() {
	// Example
	position := PositionExtended{0, 0, 0}
	position.forward(5)
	position.down(5)
	position.forward(8)
	position.up(3)
	position.down(8)
	position.forward(2)

	valuePosition := position.PositionValue()
	fmt.Printf("Day-02 Example %v value position.\n", valuePosition)

	// Part 1
	position = PositionExtended{0, 0, 0}
	lines, err := tools.ReadFile("cmd/day-02/input")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		measure := strings.Fields(line)
		action := measure[0]
		steps, err := strconv.Atoi(measure[1])
		if err != nil {
			fmt.Println(steps, "measure error")
		}

		switch action {
		case "forward":
			position.forward(steps)
		case "down":
			position.down(steps)
		case "up":
			position.up(steps)
		}
	}

	fmt.Printf("Day-02 Part 2 %v value position.\n", position.PositionValue())
}
