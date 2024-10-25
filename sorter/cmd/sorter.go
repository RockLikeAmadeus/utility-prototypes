package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
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
	mergeInsertionSortAscending(&inputList)
	fmt.Println("Sorted list: ")
	slices.Reverse(inputList)
	printSlice(inputList)
}

// mergeInsertionSortAscending performs the merge-insertion sort: https://en.wikipedia.org/wiki/Merge-insertion_sort
func mergeInsertionSortAscending(inputList *[]string) {
	if len(*inputList) == 1 {
		return
	}

	// Steps 1 and 2
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

	// Step 3
	mergeInsertionSortAscending(&winners)

	// Step 4
	// skip for now, TODO, add step 4 to reduce comparisons

	// Step 5
	// TODO insertion sort is not the optimal way to sort the remainder of the list
	// Instead, I need to insert using a *binary search*, as described by wikipedia, above.
	// Also see: https://www.geeksforgeeks.org/binary-insertion-sort/
	binaryInsertionSort(&winners, &losers)
	// insertionSort(&winners, &losers)

	*inputList = winners
}

func insertionSortAscending(sortedList *[]string, unsortedList *[]string) {
	for len(*unsortedList) != 0 {
		toSort := removeRandomElement(unsortedList)
		*sortedList = append([]string{toSort}, *sortedList...)
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
}

func binaryInsertionSort(sortedList *[]string, unsortedList *[]string) {
	fmt.Println("\nSorted list: ")
	printSlice(*sortedList)
	fmt.Println("\nUnsorted list: ")
	printSlice(*unsortedList)

	for len(*unsortedList) != 0 {
		toSort := removeRandomElement(unsortedList)
		if (len(*sortedList) == 0) {
			*sortedList = []string{toSort}
		} else {
			locationToInsert := determineSortedLocationViaBinarySearch(toSort, *sortedList)
			fmt.Println("\nInserting: ", toSort)
			fmt.Println("at location : ", locationToInsert)
			//(*sortedList) = append(append((*sortedList)[:locationToInsert], toSort), (*sortedList)[locationToInsert+1:]...)
			(*sortedList) = slices.Insert((*sortedList), locationToInsert, toSort)
		}
		fmt.Println("\nSorted list: ")
		printSlice(*sortedList)
		fmt.Println("\nUnsorted list: ")
		printSlice(*unsortedList)
	}
}

func determineSortedLocationViaBinarySearch(newItem string, sortedList []string) int {
	// Base case
	if len(sortedList) == 0 {
		return 0
	} else if len(sortedList) == 1 {
		fmt.Println("Base case: ", newItem)
		higher, _ := promptToSortTwoInputs(newItem, sortedList[0])
		if newItem == higher {
			fmt.Println("1")
			return 1
		} else {
			fmt.Println("0")
			return 0
		}
	}

	var middleIndex int = len(sortedList)/2
	fmt.Println("Middle index: ", middleIndex)
	fmt.Println("of sorted list: ")
	printSlice(sortedList)
	fmt.Println("with length ", len(sortedList))
	higher, _ := promptToSortTwoInputs(newItem, sortedList[middleIndex])
	if newItem == higher {
		return middleIndex + determineSortedLocationViaBinarySearch(newItem, sortedList[middleIndex+1:])
	} else {
		return determineSortedLocationViaBinarySearch(newItem, sortedList[:middleIndex])
	}
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
	fmt.Println("-------End of list")
}