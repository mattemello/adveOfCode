package main

import (
	"bufio"
	"log"
	"os"
)

func grabNumber(text string) map[int]bool {
	var numbers = make(map[int]bool)
	number := 0
	inNumber := false

	for i := 0; i < len(text); i++ {

		if IsNumber(text[i]) {
			number = (number * 10) + controllNumber(text[i])
			inNumber = true
		} else {
			if inNumber {
				numbers[number] = true
			}

			number = 0
			inNumber = false
		}

	}

	if inNumber {
		numbers[number] = true
	}

	return numbers
}

func pipeNumber() int {

	file, err := os.Open("./d4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		position := 0
		for text[position] != ':' {
			position++
		}

		text = text[position:]

		position = 0
		for text[position] != '|' {
			position++
		}

		winningNumber := grabNumber(text[:position])
		myNumber := grabNumber(text[position:])

		timeIn := 0

		for key, _ := range myNumber {
			_, isIn := winningNumber[key]

			if isIn {
				if timeIn == 0 {
					timeIn = 1
				} else {
					timeIn *= 2
				}
			}
		}

		sum += timeIn

	}

	return sum

}
