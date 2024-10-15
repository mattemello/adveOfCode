package main

import (
	"bufio"
	"fmt"
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

func PipeNumber2() int {

	file, err := os.Open("./test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var n = make(map[int]int)

	i := 0
	for scanner.Scan() {
		i++
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
				timeIn += 1
			}
		}

		n[i] = 0
		if timeIn > 0 {
			n[i] = timeIn
		}

	}

	return sum(n)
}

/*
	tutte le istanze sono 1

	  1     2     3     4     5     6
	[ 1 ] [ 1 ] [ 1 ] [ 1 ] [ 1 ] [ 1 ]
	  4
	[ 1 ] [ 1+1 ] [ 1+1 ] [ 1+1 ] [ 1+1 ] [ 1 ]
			2
	[ 1 ] [ 1+1 ] [ 1+1+1+1 ] [ 1+1+1+1 ] [ 1+1 ] [ 1 ]
				  2
	[ 1 ] [ 1+1 ] [ 1+1+1+1 ] [ 1+1+1+1 ] [ 1 ] [ 1 ]

	[ 1 ] [ 1 ] [ 1 ] [ 1 ] [ 1 ] [ 1 ]
*/

/*
	map[int]int
	[1]->4
	[2]->2    ->2
	[3]->2    ->4
	[4]->1    ->8
	[5]->0    ->
	[6]->0
*/

func sum(n map[int]int) int {

	sum := 0
	dim := len(n)
	p := make([]int, dim)
	for i, _ := range p {
		p[i] = 0
	}

	for i, m := range n {
		fmt.Println("[ ", i, " ] -> ", m)
		if p[i] > 0 {

			for c := 0; c < p[i]; c++ {
				for j := 0; j < m; j++ {
					p[i+m-1] += 1
				}
			}
		}
		for m >= 0 {
			p[i+m-1] += 1
			m--
		}
	}

	for _, i := range p {
		sum += p[i]
	}

	return sum

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
