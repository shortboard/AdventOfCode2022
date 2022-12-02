package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("calories.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentElf := 0
	maxElf := []int{0, 0, 0}

	for scanner.Scan() {
		if scanner.Text() == "" {
			insertIntoTop3(maxElf, currentElf)
			currentElf = 0
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			currentElf += calories
		}
	}

	totalMax := 0
	for _, elf := range maxElf {
		totalMax += elf
	}

	fmt.Println("Max elf: ", maxElf)
	fmt.Println("Total max: ", totalMax)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func insertIntoTop3(arr []int, val int) []int {
	if val > arr[0] {
		arr[0] = val
	} else if val > arr[1] {
		arr[1] = val
	} else if val > arr[2] {
		arr[2] = val
	}
	return arr
}
