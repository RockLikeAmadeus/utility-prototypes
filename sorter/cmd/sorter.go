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
	// mergeInsertionSort(&inputList) //put back in
	testSortedList := []string{"a", "c", "e"}
	testUnsortedList := []string{"b", "f", "d", "g"}
	insertionSort(&testSortedList, &testUnsortedList)

	fmt.Println("Sorted list: ")
	printSlice(testSortedList)

	// fmt.Println("Sorted list: ") // put back in
	// printSlice(inputList)
}

// mergeInsertionSort performs the merge-insertion sort: https://en.wikipedia.org/wiki/Merge-insertion_sort
func mergeInsertionSort(inputList *[]string) {
	if len(*inputList) == 1 {
		return
	}

	// Step 1 and 2
	winners := make([]string, 0, len(*inputList)/2 + 1)
	losers := make([]string, 0, len(*inputList)/2 + 1)

	for len(*inputList) > 1 {
		s1 := removeRandomElement(inputList)
		s2 := removeRandomElement(inputList)
		higher, lower := promptToSortTwoInputs(s1, s2)
		winners = append(winners, higher)
		losers = append(losers, lower)
	}
	if len(*inputList) == 1 {
		losers = append(losers, (*inputList)[0])
	}

	// fmt.Println("Winners: ")
	// printSlice(winners)
	// fmt.Println("Losers: ")
	// printSlice(losers)

	// Step 3
	mergeInsertionSort(&winners)

	// Step 4
	// skip for now

	// Step 5
	insertionSort(&winners, &losers)
}

func insertionSort(sortedList *[]string, unsortedList *[]string) {
	// newList := make([]string, 0, len(*sortedList) + len(*unsortedList))
	for len(*unsortedList) != 0 {
		toSort := removeRandomElement(unsortedList)
		*sortedList = append([]string{toSort}, *sortedList...)
		fmt.Println("State:")
		printSlice(*sortedList)
		for i := range len(*sortedList)-1 {
			higher, _ := promptToSortTwoInputs((*sortedList)[i], (*sortedList)[i+1])
			if toSort == higher {
				// Out of order, need to swap
				(*sortedList)[i] = (*sortedList)[i+1]
				(*sortedList)[i+1] = toSort
			} else {
				// In correct order, no need to continue through the sorted list
				break
			}
		}
	}

	// for idx, value := range *sortedList {
	// 	higher, lower := promptToSortTwoInputs(value, toSort)
	// 	// newList = append(newList, higher)
	// 	// newList = append(newList, lower)
	// 	if toSort == higher {
	// 		sortedList = append(append((*sortedList)[:idx], toSort), (*sortedList)[idx:])
	// 	} else {
	// 		// We're done with this one, no need to go through the rest of the sorted list
	// 		// sortedList = /
	// 		break
	// 	}
	// }

}



// promptToSortTwoInputs prompts the user to enter the higher of the inputs, then returns them in order
func promptToSortTwoInputs(s1 string, s2 string) (string, string) {
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
		return s1, s2
	} else if input == '2' {
		return s2, s1
	} else {
		fmt.Println("Error: Please enter either '1' or '2'")
		return promptToSortTwoInputs(s1, s2)
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