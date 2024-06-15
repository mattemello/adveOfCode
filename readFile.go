package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func controllNumber(cha byte) int {
	switch cha {
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case '0':
		return 0
	}
	return -1
}

func controllWord(word string) int {
	if strings.Contains(word, "one") {
		return 1
	} else if strings.Contains(word, "two") {
		return 2
	} else if strings.Contains(word, "three") {
		return 3
	} else if strings.Contains(word, "four") {
		return 4
	} else if strings.Contains(word, "five") {
		return 5
	} else if strings.Contains(word, "six") {
		return 6
	} else if strings.Contains(word, "seven") {
		return 7
	} else if strings.Contains(word, "eight") {
		return 8
	} else if strings.Contains(word, "nine") {
		return 9
	}

	return -1
}

func decimal(firstN int, secondN int) int {
	return ((firstN * 10) + secondN)
}

func takeTheNumber() int {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		firstNumber := -1
		secondNumber := -1

		for i := 0; i < len(text); i++ {
			firstNumber = controllWord(text[0:i])
			if firstNumber != -1 {
				break
			}
			firstNumber = controllNumber(text[i])
			if firstNumber != -1 {
				break
			}
		}

		for i := len(text) - 1; i >= 0; i-- {
			secondNumber = controllWord(text[i:])
			if secondNumber != -1 {
				break
			}
			secondNumber = controllNumber(text[i])
			if secondNumber != -1 {
				break
			}
		}

		sum = sum + decimal(firstNumber, secondNumber)
	}

	return sum
}
