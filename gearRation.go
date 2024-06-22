package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func IsPoint(char byte) bool {
	return char == '.'
}

func IsNumber(char byte) bool {
	return controllNumber(char) != -1
}

func IsSimble(char byte) bool {
	return !(IsNumber(char) || IsPoint(char))
}

func IsMolt(char byte) bool {
	return char == '*'
}

/*
[-1,-1][0,-1][1.-1]
[-1,0] [0,0] [1.0]
[-1,1][0,1][1.1]
*/
func takeBlock(file [][]byte, x, y int) int {
	var table [][]int
	var startCountx int
	var endCountx int
	var startCounty int
	var endCounty int

	if startCountx = -1; x == 0 {
		startCountx = 0
	}
	if startCounty = -1; x == 0 {
		startCounty = 0
	}
	if endCounty = 3; x == len(file) {
		endCounty = 2
	}
	if endCountx = 3; x == len(file) {
		endCountx = 2
	}

	table = make([][]int, endCounty)

	for i := startCounty; i < endCounty-1; i++ {
		table[i+1] = make([]int, 3)
		for j := startCountx; j < endCountx-1; j++ {
			if IsNumber(file[y+i][x+j]) {
				table[i+1][j+1] = 1
			} else if IsMolt(file[y+i][x+j]) {
				table[i+1][j+1] = 2
			} else {
				table[i+1][j+1] = 0
			}
		}
	}

	molt := 1
	countEnter := 0
	for i := 0; i < endCounty; i++ {
		IsOne := false
		for j := 0; j < endCounty; j++ {
			if table[i][j] == 1 {
				IsOne = true
			} else if table[i][j] != 1 && IsOne {
				countEnter++
				num := takeNumberFromTable(file, x+(j-1+startCountx), y+(i+startCounty))
				molt *= num
				IsOne = false
			}
		}
		if IsOne {
			countEnter++
			num := takeNumberFromTable(file, x+(endCountx+startCountx-1), y+(i+startCounty))
			molt *= num
			IsOne = false
		}
	}

	if countEnter > 1 {
		return molt
	} else {
		return 0
	}
}

func takeNumberFromTable(file [][]byte, x, y int) int {
	num := 0
	count := 0

	for IsNumber(file[y][x-count]) {
		valueNum := controllNumber(file[y][x-count])
		if count == 0 {
			num = valueNum
		} else {
			num = (valueNum * (int(math.Pow10(count)))) + num
		}
		count++
		if x-count < 0 {
			break
		}
	}

	count = 1

	if x+count <= len(file[y]) {
		for IsNumber(file[y][x+count]) {
			num = (num * 10) + controllNumber(file[y][x+count])
			count++
			if x+count >= len(file[y]) {
				break
			}
		}
	}

	return num
}

func GearRation(file [][]byte) int {
	sum := 0

	for y := 0; y < len(file); y++ {
		for x := 0; x < len(file[y]); x++ {
			if IsMolt(file[y][x]) {
				sum += takeBlock(file, x, y)

			}
		}
	}

	return sum

}

func controllTheConten(file [][]byte) int {
	sum := 0

	for y := 0; y < len(file); y++ {
		for x := 0; x < len(file[y]); x++ {
			if IsPoint(file[y][x]) {
				continue
			} else if IsNumber(file[y][x]) {
				numTo := controllNumber(file[y][x])
				numberN := 1
				for IsNumber(file[y][x+numberN]) {
					numTo = (numTo * 10) + controllNumber(file[y][x+numberN])
					numberN++
					if x+numberN > len(file[y])-1 {
						break
					}
				}

				contrAll := false
				if y == 0 {
					for i := 0; i < 2; i++ {
						if x == 0 {
							for j := 0; j < numberN+1; j++ {
								if x+j < len(file[y]) {
									contrAll = contrAll || IsSimble(file[y+i][x+j])
									if contrAll {
										break
									}
								}
							}
						} else {
							for j := -1; j < numberN+1; j++ {
								if x+j < len(file[y]) {
									contrAll = contrAll || IsSimble(file[y+i][x+j])
									if contrAll {
										break
									}
								}
							}
						}

					}
				} else {
					for i := -1; i < 2; i++ {
						if y+i < len(file) {
							if x == 0 {
								for j := 0; j < numberN+1; j++ {
									if x+j < len(file[y]) {
										contrAll = contrAll || IsSimble(file[y+i][x+j])
										if contrAll {
											break
										}
									}
								}
							} else {
								for j := -1; j < numberN+1; j++ {
									if x+j < len(file[y]) {
										contrAll = contrAll || IsSimble(file[y+i][x+j])
										if contrAll {
											break
										}
									}
								}
							}

						}
					}
				}

				if contrAll {
					sum += numTo
				}
				x += numberN
			}
		}
	}

	return sum

}

//[i-1][j-1] [i-1][j] ... [i-1][j+numberN+1]
//[i][j-1] [i][j](number)...[i][j+numberN+1]
//[i+1][j-1] [i+1][j] ... [i+1][j+numberN+1]

func takeFromTheFile() int {
	var fileContent = [][]byte{}
	dimensioFile := 0

	file, err := os.Open("./d3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dimensioFile++
	}

	file.Close()

	file, err = os.Open("./d3")

	scanner2 := bufio.NewScanner(file)

	y := 0

	fileContent = make([][]byte, dimensioFile)
	for scanner2.Scan() {
		text := scanner2.Text()

		fileContent[y] = make([]byte, len(text))
		for x := 0; x < len(text); x++ {
			fileContent[y][x] = text[x]
		}
		y++
	}

	//return controllTheConten(fileContent)
	return GearRation(fileContent)
}
