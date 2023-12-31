package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cube_conumdrum(line string) bool {
	var (
		cR uint64
		cG uint64
		cB uint64
	)
	game := strings.Split(line, ": ")[1]
	sets := strings.Split(game, "; ")

	for _, b := range sets {
		cubes := strings.Split(b, ", ")
		for i := 0; i < len(cubes); i++ {
			aux, color := strings.Split(cubes[i], " ")[0], strings.Split(cubes[i], " ")[1]
			quantity, _ := strconv.ParseUint(aux, 10, 64)
			if color == "red" {
				cR += quantity
			} else if color == "green" {
				cG += quantity
			} else {
				cB += quantity
			}
			if cR > 12 || cG > 13 || cB > 14 {
				return false
			}
		}
		cR, cB, cG = 0, 0, 0
	}
	return true
}
func main() {
	var (
		arr []bool
		sum int
	)
	file, err := os.Open("../2.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, cube_conumdrum(line))
	}
	for i, b := range arr {
		if b {
			sum += i + 1
		}
	}
	fmt.Println("Part 1: ", sum)
}
