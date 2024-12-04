package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// open the file
	// https://gobyexample.com/reading-files
	file, err := os.Open("./solutions/02/input.txt")
	check(err)
	defer file.Close()

	// get a list of reports
	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		ints := make([]int, len(parts))
		for i, s := range parts {
			num, err := strconv.Atoi(s)
			check(err)
			ints[i] = num
		}
		reports = append(reports, ints)
	}

	// score the reports
	numValid := 0
	for _, report := range reports {
		all_decreasing := true
		all_increasing := true
		all_diff_ok := true

		for i := 0; i < len(report)-1; i++ {
			prev, curr := report[i], report[i+1]
			all_decreasing = all_decreasing && (prev > curr)
			all_increasing = all_increasing && (prev < curr)
			all_diff_ok = all_diff_ok && (1 <= AbsInt(prev-curr) && AbsInt(prev-curr) <= 3)
		}

		safe := (all_increasing != all_decreasing) && all_diff_ok
		if safe {
			numValid += 1
		}

		fmt.Println(numValid)

	}

}
