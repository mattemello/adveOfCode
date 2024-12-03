package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func moltNumbers(num1, num2 string) int {
	nu1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Println("error in the atoi ", err)
		os.Exit(2)
	}
	nu2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println("error in the atoi ", err)
		os.Exit(2)
	}

	return (nu1 * nu2)

}

func firstPart(scanner *bufio.Scanner) {

	lineCode := ""
	sum := 0

	for scanner.Scan() {

		lineCode = scanner.Text()

		theSum := strings.Split(lineCode, "mul")

		for _, cell := range theSum {
			if cell != "" {
				if cell[0] == '(' {
					number := ""
					var numbers = make([]string, 2)
					attualNum := 0

					for i := 1; i < len(cell); i++ {
						if cell[i] != ',' && cell[i] != ')' && !unicode.IsNumber(rune(cell[i])) {
							break

						} else {
							if unicode.IsNumber(rune(cell[i])) {
								number += string(cell[i])
							} else if cell[i] == ',' {
								if number == "" || attualNum != 0 {
									break
								}
								numbers[attualNum] = number

								number = ""
								attualNum++
							} else {
								if number == "" || attualNum != 1 {
									break
								}
								numbers[attualNum] = number
								number = ""

								sum += moltNumbers(numbers[0], numbers[1])

								break
							}
						}
					}
				}
			}

		}
	}

	fmt.Println(sum)
}

func secondPart(scanner *bufio.Scanner) {

	var cmd = regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	sum := 0
	do := true

	for scanner.Scan() {

		lineCode := scanner.Text()
		lineCodes := cmd.FindAllString(lineCode, -1)

		for _, line := range lineCodes {
			if line[0:3] == "mul" {

				if do {
					num1 := line[4:strings.Index(line, ",")]
					num2 := line[strings.Index(line, ",")+1 : len(line)-1]
					sum += moltNumbers(num1, num2)
				}

			} else {
				do = line == "do()"
			}

		}
	}

	fmt.Println(sum)
}

func main() {

	file, err := os.Open("day3")
	if err != nil {
		os.Exit(1)
	}

	var scanner = bufio.NewScanner(file)

	firstPart(scanner)

	_, err = file.Seek(0, io.SeekStart)

	var scanner2 = bufio.NewScanner(file)

	secondPart(scanner2)

}
