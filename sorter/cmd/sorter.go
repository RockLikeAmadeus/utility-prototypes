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

	// for _, item := range inputList {
	// 	fmt.Println(item)
	// }
	sortInput(&inputList)
}

func sortInput(inputList *[]string) {
	s1 := removeRandomElement(inputList)
	s2 := removeRandomElement(inputList)
	err := promptToSortValues(s1, s2)
	for err != nil {
		fmt.Println("Error: Please enter either '1' or '2'")
		err = promptToSortValues(s1, s2)
	}
}

func promptToSortValues( s1 string, s2 string) error {
	reader := bufio.NewReader(os.Stdin)
	// for {
		// topHalf := make([]string, 0)
		// bottomHalf := make([]string, 0)
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
			return nil
		} else if input == '2' {
			fmt.Println("good choice")
			return nil
		} else {
			return errors.New("invalid input")
		}
	// }
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
	fmt.Println()
	fmt.Println("Printing input list: ")
	for _, item := range inputList {
		fmt.Println(item)
	}
}