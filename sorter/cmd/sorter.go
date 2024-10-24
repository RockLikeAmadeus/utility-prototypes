package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	file, err := os.Open("todo.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	inputList := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	// for _, item := range inputList {
	// 	fmt.Println(item)
	// }
	sortInput(inputList)
}

func sortInput(inputList []string) {
	printInput(inputList)
	fmt.Println("Removed: ", removeRandomElement(&inputList))
	printInput(inputList)

	// s1 = inputList[len(inputList)-1]
	// topHalf := make([]string, 0)
	// bottomHalf := make([]string, 0)
	fmt.Println("Choose the larger of")
	fmt.Println(rand.Intn(len(inputList)))
}

func removeElementAtIndex(index int, inputList *[]string) string {
	res := (*inputList)[index]
	(*inputList) = append((*inputList)[:index], (*inputList)[index + 1:]...)
	return res
}


func removeRandomElement(inputList *[]string) string {
	return removeElementAtIndex(rand.Intn(len(*inputList)), inputList)
}



func printInput(inputList []string) {
	fmt.Println()
	fmt.Println("Printing input list: ")
	for _, item := range inputList {
		fmt.Println(item)
	}
}