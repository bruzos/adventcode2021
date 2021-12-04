package main

import (
	"codeadvent/tools"
	"fmt"
)

func DepthSum(depths []int) int {
	var depthsIncreasedCount int
	for i, depth := range depths {
		if i > 0 {
			if depth > depths[i-1] {
				depthsIncreasedCount++
			}
		}
	}
	return depthsIncreasedCount
}

func averagedDepths(depths []int) []int {
	var averagedSlice []int
	positionLimit := len(depths) - 2
	for i := range depths {
		if i >= positionLimit {
			continue
		}
		partially := partialArray(i, depths)
		summable := sumArray(partially)
		averagedSlice = append(
			averagedSlice,
			summable,
		)
	}
	return averagedSlice
}

func partialArray(from int, depths []int) []int {
	to := from + 3
	partial := depths[from:to]
	return partial
}

func sumArray(depths []int) int {
	result := 0
	for _, depth := range depths {
		result += depth
	}
	return result
}

func main() {
	lines, err := tools.ReadFile("cmd/day-01/input")
	if err != nil {
		panic(err)
	}

	var depths []int
	depths, err = tools.ReadIntLines(lines)
	if err != nil {
		panic(err)
	}
	measurementsCount := DepthSum(depths)
	valueDepths := averagedDepths(depths)
	depthsIncreasedCount := DepthSum(valueDepths)
	fmt.Printf("Day-01 Part 1 - %v measurements that are larger than the previous.\n", measurementsCount)
	fmt.Printf("Day-01 Part 2 - %v sums that are larger than the previous sum.\n", depthsIncreasedCount)

}
