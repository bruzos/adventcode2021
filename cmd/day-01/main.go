package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(name string) (nums []int, err error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

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
	depths, err := readFile("cmd/day-01/input")
	if err != nil {
		panic(err)
	}

	depthsIncreasedCount := DepthSum(depths)
	fmt.Printf("Day-01 Part 1 - %v measurements that are larger than the previous.\n", depthsIncreasedCount)

	valueDepths := averagedDepths(depths)
	depthsIncreasedCount = DepthSum(valueDepths)
	fmt.Printf("Day-01 Part 2 - %v sums that are larger than the previous sum.\n", depthsIncreasedCount)

}
