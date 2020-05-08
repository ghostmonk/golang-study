package main

import "fmt"

var matrix = [][]int {	{10, 20, 30, 40},
                      	{15, 25, 35, 45},
                      	{27, 29, 37, 48},
                      	{32, 33, 39, 50},
											}

func MatricesTest() {
	runAll(matrix, runMatricesTestBrute, 100)
	runAll(matrix, runMatricesTestBetter, 100)
}

func runAll(data [][]int, test func([][]int, int)(int, int, int), failCase int) {
	for _, stack := range data {
		for _, val := range stack {
			fmt.Println(test(data, val))
		}
	}
	fmt.Println(test(data, failCase))
}

func runMatricesTestBetter(haystack [][]int, needle int) (int, int, int) {
	totalIterations := 0
	i := 0
	j := 0
	for ; i < len(haystack); i++ {
		innerStack := haystack[i]
		for j = len(innerStack) - 1; j >= 0; j-- {
			totalIterations++
			target := innerStack[j]
			
			if needle > target {
				break
			}

			if needle == target {
				return i, j, totalIterations
			}
		}
	}
	return i, j, totalIterations	
}

func runMatricesTestBrute(haystack [][]int, needle int) (int, int, int) {
	i := 0
	for x, innerStack := range haystack {
		for y, val := range innerStack {
			i++
			if val == needle {
				return x, y, i
			}
		} 
	}
	return -1, -1, i
}
