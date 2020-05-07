package main

import "fmt"

var test1 = []int{1, 1, 3}
var test2 = []int{1, 1, 1}
var test3 = []int{1, 2, 3}
var test4 = []int{1, 2, 3, 4, 4, 2, 3, 1, 1, 4, 1}

func DuplicatesTest() {
	fmt.Println(runTest(test1))
	fmt.Println(runTest(test2))
	fmt.Println(runTest(test3))
	fmt.Println(runTest(test4))
}

func runTest(input []int) int {
	counts := make(map[int]int)
	output := 0
	for _, v := range input {
		counts[v]++
		nc := counts[v]
		pc := nc - 1
		output -= (pc * (pc - 1)) / 2
		output += (nc * (nc - 1)) / 2
	}

	return output
}
