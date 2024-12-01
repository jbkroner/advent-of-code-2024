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

	// build a frequncy mapp for arr_r
	// https://gobyexample.com/maps
	arr_r_map := make(map[int]int)
	for _, num := range arr_r {
		_, exists := arr_r_map[num]
		if exists {
			arr_r_map[num] += 1
		} else {
			arr_r_map[num] = 1
		}
	}

	// calculate similiarty:
	// multiply values in left list by their frequency in right list
	// then add it to the sum
	sum := 0
	// for i := 0; i < len(arr_l); i++
	for i := range arr_l {
		freq, exists := arr_r_map[arr_l[i]]
		if !exists {
			freq = 0
		}
		sim := arr_l[i] * freq
		sum += sim
	}

	fmt.Println("similarity: ", sum)

}
