package main

import "fmt"

func CollectionsTest() {
	howDoArraysWork()
	howDoSlicesWork()
	simpleSliceActivity()
}

func howDoArraysWork() {
	var simpleArray [10]string
	simpleArray[0] = "This"
	simpleArray[1] = "is"
	simpleArray[2] = "an"
	simpleArray[3] = "advanced"
	simpleArray[4] = "version"
	simpleArray[5] = "of"
	simpleArray[6] = "the"
	simpleArray[7] = "Hello"
	simpleArray[8] = "World"
	simpleArray[9] = "application!!"
	fmt.Println(simpleArray)
}

func howDoSlicesWork() {
	simpleSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(simpleSlice)
	fmt.Println(simpleSlice[1:4])
	fmt.Println(simpleSlice[:3])
	fmt.Println(simpleSlice[5:])
}

func simpleSliceActivity() {
	var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
		"Emma", "Isabella", "Emily", "Madison",
		"Ava", "Olivia", "Sophia", "Abigail",
		"Elizabeth", "Chloe", "Samantha",
		"Addison", "Natalie", "Mia", "Alexis"}

	output := [][]string{}
	for _, name := range names {
		target := output[len(name)]
		target = append(target, name)
	}
	fmt.Println(output)
}
