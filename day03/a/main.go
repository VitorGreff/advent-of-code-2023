package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type part_number struct {
	value  uint64
	row    int
	column int
	size   int
	summed bool
}

func gear_ratios(m [][]string) {
	part_numbers := []part_number{}
	for i := range m {
		var str string
		for j := range m[i] {
			_, err := strconv.ParseUint(m[i][j], 10, 64)
			if err == nil {
				str += m[i][j]
			} else if str != "" {
				number, _ := strconv.ParseUint(str, 10, 64)
				part_numbers = append(part_numbers, part_number{number, i, j - 1, len(strconv.Itoa(int(number))), false})
				str = ""
			}
		}
		if str != "" {
			number, _ := strconv.ParseUint(str, 10, 64)
			part_numbers = append(part_numbers, part_number{number, i, len(m[i]) - 1, len(strconv.Itoa(int(number))), false})
		}
	}
	fmt.Println(check_parts(m, part_numbers))
}

func check_parts(m [][]string, arr []part_number) uint64 {
	var sum uint64
	for _, p := range arr {
		// above
		if p.row > 0 {
			for i := 0; i < p.size; i++ {
				if !p.summed {
					p.summed = contain_patters(m[p.row-1][p.column-i])
				}
			}
		}
		// below
		if p.row < len(m)-1 {
			for i := 0; i < p.size; i++ {
				if !p.summed {
					p.summed = contain_patters(m[p.row+1][p.column-i])
				}
			}
		}
		// right side
		if p.column < len(m[0])-1 {
			if !p.summed {
				p.summed = contain_patters(m[p.row][p.column+1])
			}
		}
		// left side
		if p.column-p.size > 0 {
			if !p.summed {
				p.summed = contain_patters(m[p.row][p.column-p.size])
			}
		}
		// upper left diagonal
		if p.row > 0 && p.column-p.size > 0 {
			if !p.summed {
				p.summed = contain_patters(m[p.row-1][p.column-p.size])
			}
		}
		// upper right diagonal
		if p.row > 0 && p.column < len(m[0])-1 {
			if !p.summed {
				p.summed = contain_patters(m[p.row-1][p.column+1])
			}
		}
		// lower left diagonal
		if p.row < len(m)-1 && p.column-p.size > 0 {
			if !p.summed {
				p.summed = contain_patters(m[p.row+1][p.column-p.size])
			}
		}
		// lower right diagonal
		if p.row < len(m)-1 && p.column < len(m[0])-1 {
			if !p.summed {
				p.summed = contain_patters(m[p.row+1][p.column+1])
			}
		}
		if p.summed {
			sum += p.value
		}
		fmt.Print(p, " \n")
	}
	return sum
}

func contain_patters(char string) bool {
	if char != "." {
		if _, err := strconv.ParseUint(char, 10, 64); err != nil {
			return true
		}
	}
	return false
}

func main() {
	var matrix [][]string
	file, err := os.Open("../3.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		matrix = append(matrix, line)
	}
	gear_ratios(matrix)
}
