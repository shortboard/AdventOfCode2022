package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, ",")
		low1, high1 := getHighAndLow(s[0])
		low2, high2 := getHighAndLow(s[1])

		fmt.Println(low1, high1, low2, high2)

		if low1 <= low2 && high1 >= high2 {
			sum++
		} else if low2 <= low1 && high2 >= high1 {
			sum++
		} else if low1 <= high2 && high1 >= high2 {
			sum++
		} else if low2 <= high1 && high2 >= high1 {
			sum++
		}

	}

	fmt.Println(sum)

}

func getHighAndLow(input string) (int, int) {
	s := strings.Split(input, "-")
	low, _ := strconv.Atoi(s[0])
	high, _ := strconv.Atoi(s[1])
	return low, high
}
