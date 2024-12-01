package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	file, err := os.Open("./solutions/01/input.txt")
	check(err)
	defer file.Close()

	// process the file into two lists
	scanner := bufio.NewScanner(file)
	arr_l := []int{}
	arr_r := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		// https://gobyexample.com/string-functions
		parts := strings.Split(line, "   ")
		l, r := parts[0], parts[1]

		// convert to int
		l_int, _ := strconv.Atoi(l)
		r_int, _ := strconv.Atoi(r)

		arr_l = append(arr_l, l_int)
		arr_r = append(arr_r, r_int)
	}

	// sort
	slices.Sort(arr_l)
	slices.Sort(arr_r)

	// calculate distances
	sum := 0
	// for i := 0; i < len(arr_l); i++
	for i := range arr_l {
		distance := AbsInt(arr_l[i] - arr_r[i]) // math.Abs works on floats
		sum += distance
	}

	fmt.Println("Total distance: ", sum)

}
