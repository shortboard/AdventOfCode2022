package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("strategy.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0

	for scanner.Scan() {
		fmt.Print(string(scanner.Text()[0]), string(scanner.Text()[2]))

		if string(scanner.Text()[2]) == "X" {
			if string(scanner.Text()[0]) == "A" {
				totalScore += 3
			}
			if string(scanner.Text()[0]) == "B" {
				totalScore += 1
			}
			if string(scanner.Text()[0]) == "C" {
				totalScore += 2
			}
		}
		if string(scanner.Text()[2]) == "Y" {
			if string(scanner.Text()[0]) == "A" {
				totalScore += 4
			}
			if string(scanner.Text()[0]) == "B" {
				totalScore += 5
			}
			if string(scanner.Text()[0]) == "C" {
				totalScore += 6
			}
		}
		if string(scanner.Text()[2]) == "Z" {
			if string(scanner.Text()[0]) == "A" {
				totalScore += 8
			}
			if string(scanner.Text()[0]) == "B" {
				totalScore += 9
			}
			if string(scanner.Text()[0]) == "C" {
				totalScore += 7
			}
		}
		fmt.Println("Running total:", totalScore)
	}
	fmt.Println("Total Score:", totalScore)
}
