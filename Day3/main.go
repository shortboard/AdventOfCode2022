package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("rucksacks.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	group := []string{}
	for scanner.Scan() {
		group = append(group, scanner.Text())

		if len(group) == 3 {
			firstTwo := findCommonCharacters(group[0], group[1])
			character := findCommonCharacters(firstTwo, group[2])
			sum += int(calcultePriority(character[0]))
			group = []string{}
		}
	}

	println(sum)
}

func findCommonCharacters(first, second string) string {
	commonCharacters := ""

	for _, firstHalfCharacter := range first {
		for _, secondHalfCharacter := range second {
			if firstHalfCharacter == secondHalfCharacter {
				commonCharacters += string(firstHalfCharacter)
			}
		}
	}

	return commonCharacters
}

func findCommonCharacterPriority(firstHalf, secondHalf string) byte {
	char := findCommonCharacters(firstHalf, secondHalf)[0]

	return calcultePriority(char)

}

func calcultePriority(char byte) byte {
	if char > 96 {
		return char - 96
	} else if char > 64 {
		return char - 64 + 26
	}
	return char
}
