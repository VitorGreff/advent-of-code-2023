package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func gear_ratios(m [][]string) {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == "*" {
				check_around(i, j, m)
			}
		}
	}
}

func check_around(i, j int, m [][]string) {
	counter := 0
	// left side
	if j > 0 {
		if is_number(m[i][j-1]) {
			pointer := j - 1
			str := ""
			for pointer >= 0 && is_number(m[i][pointer]) {
				str = m[i][pointer] + str
				pointer--
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
	// right side
	if j < len(m[0])-1 {
		if is_number(m[i][j+1]) {
			pointer := j + 1
			str := ""
			for pointer <= len(m[0])-1 && is_number(m[i][pointer]) {
				str = m[i][pointer] + str
				pointer++
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
	// above refactor
	if i > 0 {
		if is_number(m[i-1][j]) {
			str := m[i-1][j]
			if is_number(m[i-1][j-1]) {
				str = m[i-1][j-1] + str
			}
			if is_number(m[i-1][j+1]) {
				str += m[i-1][j+1]
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
	// below refactor
	if i < len(m)-1 {
		if is_number(m[i+1][j]) {
			str := m[i+1][j]
			if is_number(m[i+1][j-1]) {
				str = m[i+1][j-1] + str
			}
			if is_number(m[i+1][j+1]) {
				str += m[i+1][j+1]
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
	// upper left diagonal
	if i > 0 && j > 0 {
		if is_number(m[i-1][j-1]) {
			str := ""
			pointer := j - 1
			for pointer >= 0 && is_number(m[i-1][pointer]) {
				str = m[i-1][pointer] + str
				pointer--
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
	// upper right diagonal
	if i > 0 && j < len(m[0])-1 {
		if is_number(m[i-1][j+1]) {
			str := ""
			pointer := j + 1
			for pointer < len(m[0]) && is_number(m[i-1][pointer]) {
				str += m[i-1][pointer]
				pointer++
			}
			pointer = j + 1
			for pointer < len(m[0]) && is_number(m[i-1][pointer]) {
				str += m[i-1][pointer]
				pointer--
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
	// lower left diagonal
	if i < len(m[0])-1 && j > 0 {
		if is_number(m[i+1][j-1]) {
			str := ""
			pointer := j - 1
			for pointer >= 0 && is_number(m[i+1][pointer]) {
				str = m[i+1][pointer] + str
				pointer--
			}
			pointer = j - 1
			for pointer >= 0 && is_number(m[i+1][pointer]) {
				str = m[i+1][pointer] + str
				pointer++
			}
			fmt.Println("i: ", i, "j: ", j, "value: ", str)
		}
		counter++
	}
}

func is_number(char string) bool {
	if _, err := strconv.ParseUint(char, 10, 64); err == nil {
		return true
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
