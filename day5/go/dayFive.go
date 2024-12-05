package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func theRulesNumber(num []string) (int, int) {

	num1, err := strconv.Atoi(num[0])
	if err != nil {
		os.Exit(1)
	}
	num2, err := strconv.Atoi(num[1])
	if err != nil {
		os.Exit(1)
	}

	return num1, num2

}

func searchNumber(listNumber []int, element int) bool {

	for _, num := range listNumber {
		if num == element {
			return true
		}
	}

	return false

}

func firstPart(scanner *bufio.Scanner, dimensionRules int) {

	i := 0
	sum := 0
	sum2 := 0
	var rules = make(map[int][]int, dimensionRules)

	for scanner.Scan() {

		if i < dimensionRules {
			num1, num2 := theRulesNumber(strings.Split(scanner.Text(), "|"))

			var newArr = make([]int, len(rules[num1])+1)
			for pos, ele := range rules[num1] {
				newArr[pos] = ele
			}
			rules[num1] = newArr
			rules[num1][len(rules[num1])-1] = num2
		} else if scanner.Text() == "" {

		} else {

			pageUpdateString := strings.Split(scanner.Text(), ",")
			var pageUpdate = make([]int, len(pageUpdateString))
			var err error

			for pos, num := range pageUpdateString {
				pageUpdate[pos], err = strconv.Atoi(num)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}

			isNotOk := false
			for j := 0; j < len(pageUpdate)-1; j++ {
				elem, _ := rules[pageUpdate[j]]
				if !searchNumber(elem, pageUpdate[j+1]) {
					sum2 += secondPart(pageUpdate, rules)
					isNotOk = true
				}

				if isNotOk {
					break
				}
			}

			if !isNotOk {
				sum += pageUpdate[len(pageUpdate)/2]
			}
		}
		i++
	}

	fmt.Println(sum)
	fmt.Println(sum2)

}

func secondPart(pageUpdateWrong []int, rules map[int][]int) int {

	for i := 0; i < len(pageUpdateWrong)-1; i++ {
		elem, _ := rules[pageUpdateWrong[i]]
		if !searchNumber(elem, pageUpdateWrong[i+1]) {

			secondElem, isNotHere := rules[pageUpdateWrong[i+1]]

			if searchNumber(secondElem, pageUpdateWrong[i]) || isNotHere {

				temp := pageUpdateWrong[i]
				pageUpdateWrong[i] = pageUpdateWrong[i+1]
				pageUpdateWrong[i+1] = temp

				i = -1
			}
		}
	}

	return pageUpdateWrong[len(pageUpdateWrong)/2]

}

func main() {
	file, err := os.Open("../day5")
	if err != nil {
		os.Exit(1)
	}

	var scanner = bufio.NewScanner(file)

	dimension := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		dimension++
	}
	_, err = file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)

	firstPart(scanner, dimension)

}
