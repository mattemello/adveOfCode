package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const RedMax = 12
const GreenMax = 12
const BluMax = 12

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
			if precNum < BluMax && itsOkay {
				itsOkay = true
			} else {
				itsOkay = false
			}
			enter = false
		}

		if strings.Contains(value, "red") && enter {
			if precNum < RedMax && itsOkay {
				itsOkay = true
			} else {
				itsOkay = false
			}
			enter = false
		}

		if strings.Contains(value, "green") && enter {
			if precNum < GreenMax && itsOkay {
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
