package main

import (
	"bufio"
	"errors"
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
	sortInput(&inputList)
}

func sortInput(inputList *[]string) {
	winners := make([]string, 0, len(*inputList)/2 + 1)
	losers := make([]string, 0, len(*inputList)/2 + 1)

	for len(*inputList) > 1 {
		s1 := removeRandomElement(inputList)
		s2 := removeRandomElement(inputList)
		higher, lower, err := promptToSortTwoInputs(s1, s2)
		for err != nil {
			fmt.Println("Error: Please enter either '1' or '2'")
			higher, lower, err = promptToSortTwoInputs(s1, s2)
		}
		winners = append(winners, higher)
		losers = append(losers, lower)
	}
	if len(*inputList) == 1 {
		losers = append(losers, (*inputList)[0])
	}


	fmt.Println("Winners: ")
	printSlice(winners)
	fmt.Println("Losers: ")
	printSlice(losers)
}

// promptToSortTwoInputs prompts the user to enter the higher of the inputs, then returns them in order
func promptToSortTwoInputs( s1 string, s2 string) (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose the larger of: ")
	fmt.Println("1. ", s1)
	fmt.Println("2. ", s2)
	fmt.Print("Enter '1' or '2' -> ")
	input, _, err := reader.ReadRune()

	fmt.Println()
	if err != nil {
		panic(err)
	}
	if input == '1' {
		fmt.Println("good choice")
		return s1, s2, nil
	} else if input == '2' {
		fmt.Println("good choice")
		return s2, s1, nil
	} else {
		return "", "", errors.New("invalid input")
	}
}

func removeRandomElement(inputList *[]string) string {
	return removeElementAtIndex(rand.Intn(len(*inputList)), inputList)
}

func removeElementAtIndex(index int, inputList *[]string) string {
	res := (*inputList)[index]
	(*inputList) = append((*inputList)[:index], (*inputList)[index + 1:]...)
	return res
}

func printSlice(inputList []string) {
	// fmt.Println("Printing slice: ")
	for _, item := range inputList {
		fmt.Println(item)
	}
}