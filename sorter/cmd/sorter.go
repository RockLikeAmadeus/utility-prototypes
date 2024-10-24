package main

import (
	"bufio"
	"fmt"
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

	for _, item := range inputList {
		fmt.Println(item)
	}
}