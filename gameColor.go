package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const RedMax = 12
const GreenMax = 13
const BluMax = 14

func GameMinNumber() int {
	file, err := os.Open("./gameInput")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	redNumber := 1
	greenNumber := 1
	blueNumber := 1
	toReturn := 0
	precNum := -1
	enter := false
	firstTime := 0

	for scanner.Scan() {
		value := scanner.Text()

		if strings.Contains(value, ":") {
			if firstTime > 2 {
				toReturn += (redNumber * greenNumber * blueNumber)
			}
			redNumber = 1
			greenNumber = 1
			blueNumber = 1
		}

		if num, err := strconv.Atoi(value); err == nil && !strings.Contains(value, ":") {
			enter = true
			precNum = num
		}

		if strings.Contains(value, "blue") && enter {
			if precNum > blueNumber {
				blueNumber = precNum
			}
			enter = false
		}

		if strings.Contains(value, "red") && enter {
			if precNum > redNumber {
				redNumber = precNum
			}
			enter = false
		}

		if strings.Contains(value, "green") && enter {
			if precNum > greenNumber {
				greenNumber = precNum
			}
			enter = false
		}
		firstTime++
	}

	toReturn += (redNumber * greenNumber * blueNumber)

	return toReturn

}
func GameColor() int {
	file, err := os.Open("./gameInput")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	toReturn := 0
	gameNum := 0
	precNum := -1
	enter := false
	itsOkay := true

	for scanner.Scan() {
		value := scanner.Text()

		if strings.Contains(value, ":") {
			if itsOkay {
				toReturn += gameNum
			}

			itsOkay = true
			gameNum = 0
			for i := 0; i < len(value)-1; i++ {
				gameNum = (gameNum * 10) + controllNumber(value[i])
			}
		}

		if num, err := strconv.Atoi(value); err == nil && !strings.Contains(value, ":") {
			enter = true
			precNum = num
		}

		if strings.Contains(value, "blue") && enter {
			if precNum <= BluMax && itsOkay {
				itsOkay = true
			} else {
				itsOkay = false
			}
			enter = false
		}

		if strings.Contains(value, "red") && enter {
			if precNum <= RedMax && itsOkay {
				itsOkay = true
			} else {
				itsOkay = false
			}
			enter = false
		}

		if strings.Contains(value, "green") && enter {
			if precNum <= GreenMax && itsOkay {
				itsOkay = true
			} else {
				itsOkay = false
			}
			enter = false
		}
	}
	if itsOkay {
		toReturn += gameNum
	}

	return toReturn

}
