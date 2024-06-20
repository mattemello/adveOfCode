package main

import (
	"bufio"
	"fmt"
	"log"
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

func TakeNumber(file [][]byte, x, y int) int {
	preNum := -1
	if x+preNum >= 0 {
		for IsNumber(file[y][x+preNum]) {
			preNum -= 1
		}
	}

	numTo := 0
	numberN := 1

	if x+preNum+numberN < len(file[y]) {
		for IsNumber(file[y][x+preNum+numberN]) {
			numTo = (numTo * 10) + controllNumber(file[y][x+preNum+numberN])
			numberN++
			if x+preNum+numberN > len(file[y])-1 {
				break
			}
		}
	}

	return numTo
}

func GearRation(file [][]byte) int {
	sum := 0

	for y := 0; y < len(file); y++ {
		for x := 0; x < len(file[y]); x++ {
			if IsMolt(file[y][x]) {

				firstNum := 1
				timeEnter := 0
				if y == 0 {
					if x == 0 {
						for i := 0; i < 2; i++ {
							oneNumber := false
							for j := 0; j < 2; j++ {
								if IsNumber(file[y+i][x+j]) && !oneNumber {
									timeEnter++
									oneNumber = true
									firstNum *= TakeNumber(file, y+i, x+j)
								} else {
									oneNumber = false
								}
							}
						}
					} else {
						for i := 0; i < 2; i++ {
							oneNumber := false
							for j := -1; j < 2; j++ {
								if IsNumber(file[y+i][x+j]) && !oneNumber {
									timeEnter++
									oneNumber = true
									firstNum *= TakeNumber(file, y+i, x+j)
									fmt.Println(file[x+i][y+j])
									fmt.Println(firstNum)
								} else {
									oneNumber = false
								}
							}
						}
					}

				} else {
					if x == 0 {
						for i := -1; i < 2; i++ {
							oneNumber := false
							for j := 0; j < 2; j++ {
								if IsNumber(file[y+i][x+j]) && !oneNumber {
									timeEnter++
									oneNumber = true
									firstNum *= TakeNumber(file, y+i, x+j)
									fmt.Println(file[x+i][y+j])
									fmt.Println(firstNum)
								} else {
									oneNumber = false
								}
							}
						}
					} else {
						for i := -1; i < 2; i++ {
							oneNumber := false
							for j := -1; j < 2; j++ {
								if IsNumber(file[y+i][x+j]) && !oneNumber {
									timeEnter++
									oneNumber = true
									firstNum *= TakeNumber(file, y+i, x+j)
									fmt.Println(string(file[x+i][y+j]))
									fmt.Println(firstNum)
								} else {
									oneNumber = false
								}
							}
						}

					}

				}

				if timeEnter > 1 {
					sum += firstNum
				}

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
