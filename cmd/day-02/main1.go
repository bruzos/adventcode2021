package main

import (
	"codeadvent/tools"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	X     int
	Depth int
}

func (p *Position) forward(n int) {
	p.X += n
}

func (p *Position) down(n int) {
	p.Depth -= n
}

func (p *Position) up(n int) {
	p.Depth += n
}

func (p *Position) PositionValue() float64 {
	position := float64(p.X * p.Depth)
	return math.Abs(position)
}

func main() {
	// Example
	position := Position{0, 0}
	position.forward(5)
	position.down(5)
	position.forward(8)
	position.up(3)
	position.down(8)
	position.forward(2)

	valuePosition := position.PositionValue()
	fmt.Printf("Day-02 Example %v value position.\n", valuePosition)

	// Part 1
	position = Position{0, 0}
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

	valuePosition = position.PositionValue()
	fmt.Printf("Day-02 Part 1 %v value position.\n", int(valuePosition))
}
