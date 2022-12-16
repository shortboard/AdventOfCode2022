package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := make([]Stack, 9)
	setup := true
	var firstLines Stack
	for scanner.Scan() {
		text := scanner.Text()
		if setup {
			if len(text) != 0 {
				firstLines.Push(scanner.Text())
				continue
			} else {
				setup = false

				for i := 0; i < 9; i++ {
					var line = firstLines.Pop()
					if i == 0 {
						continue
					}
					vals := []string{line[1:2], line[5:6], line[9:10], line[13:14], line[17:18], line[21:22], line[25:26], line[29:30], line[33:34]}

					fmt.Println(vals)
					for s, val := range vals {
						if val != " " {
							stacks[s].Push(val)
						}
					}
				}
			}
		} else {
			crane := Stack{}

			s := strings.Split(text, " ")
			var from, to int
			howMany, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println(err)
			}
			from, err = strconv.Atoi(s[3])
			if err != nil {
				fmt.Println(err)
			}
			to, err = strconv.Atoi(s[5])
			if err != nil {
				fmt.Println(err)
			}

			for i := 0; i < howMany; i++ {
				crane.Push(stacks[from-1].Pop())
			}
			for len(crane) > 0 {
				stacks[to-1].Push(crane.Pop())
			}

		}
	}

	fmt.Println(stacks)

	for i, stack := range stacks {
		fmt.Printf("Stack %d: ", i+1)
		for len(stack) > 0 {
			fmt.Printf("%s ", stack.Pop())
		}
		fmt.Println()
	}

}
