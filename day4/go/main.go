package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func controlXmas(m, a, s byte) bool {

	if m == 'M' && a == 'A' && s == 'S' {
		return true
	}

	return false
}

func firstPart(scanner *bufio.Scanner, dimension int) {

	var file = make([]string, dimension)
	position := 0

	for scanner.Scan() {
		file[position] = scanner.Text()
		position++
	}

	result := 0
	for i, line := range file {
		for j, charac := range line {
			if charac == 'X' {
				if i >= 3 {
					if j >= 3 {
						if controlXmas(file[i-1][j-1], file[i-2][j-2], file[i-3][j-3]) {
							result += 1
						}
					}
					if len(line)-j > 3 {
						if controlXmas(file[i-1][j+1], file[i-2][j+2], file[i-3][j+3]) {
							result += 1
						}
					}
					if controlXmas(file[i-1][j], file[i-2][j], file[i-3][j]) {
						result += 1
					}
				}
				if dimension-i > 3 {
					if j >= 3 {
						if controlXmas(file[i+1][j-1], file[i+2][j-2], file[i+3][j-3]) {
							result += 1
						}
					}
					if len(line)-j > 3 {
						if controlXmas(file[i+1][j+1], file[i+2][j+2], file[i+3][j+3]) {
							result += 1
						}
					}
					if controlXmas(file[i+1][j], file[i+2][j], file[i+3][j]) {
						result += 1
					}
				}
				if j >= 3 {
					if controlXmas(file[i][j-1], file[i][j-2], file[i][j-3]) {
						result += 1
					}

				}
				if len(line)-j > 3 {
					if controlXmas(file[i][j+1], file[i][j+2], file[i][j+3]) {
						result += 1
					}

				}

			}
		}
	}

	fmt.Println(result)

}

func controlMas(m, s byte) bool {

	if (m == 'M' && s == 'S') || (m == 'S' && s == 'M') {
		return true
	}

	return false
}

func secondPart(scanner *bufio.Scanner, dimension int) {

	var file = make([]string, dimension)
	position := 0

	for scanner.Scan() {
		file[position] = scanner.Text()
		position++
	}

	result := 0
	for i, line := range file {
		for j, charac := range line {
			if charac == 'A' {
				if i >= 1 && dimension-i > 1 {
					if len(line)-j > 1 && j >= 1 {
						if controlMas(file[i-1][j-1], file[i+1][j+1]) && controlMas(file[i-1][j+1], file[i+1][j-1]) {
							result += 1
						}
					}
				}
			}
		}
	}

	fmt.Println(result)

}

func main() {
	file, err := os.Open("../day4")
	if err != nil {
		os.Exit(1)
	}

	var scanner = bufio.NewScanner(file)

	dimension := 0
	for scanner.Scan() {
		dimension++
	}
	_, err = file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)

	firstPart(scanner, dimension)

	_, err = file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)

	secondPart(scanner, dimension)
}
