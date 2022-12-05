package main_1

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
	for scanner.Scan() {
		text := scanner.Text()
		firstHalf := text[:len(text)/2]
		secondHalf := text[len(text)/2:]

		commonCharacter := findCommonCharacterPriority(firstHalf, secondHalf)

		sum += int(commonCharacter)

		println(commonCharacter)
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
