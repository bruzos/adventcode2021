package tools

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(name string) (lines []string, err error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines = strings.Split(string(b), "\n")
	return lines, nil
}

func ReadIntLines(lines []string) (nums []int, err error) {
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

//func ReadMeasurements(lines []string){
//	measurements := make([]string, 0, len(lines))
//	for _, line := range lines {
//		measurements := strings.Fields(line)
//		measurements = append(measurements, measurement)
//	}
//}
