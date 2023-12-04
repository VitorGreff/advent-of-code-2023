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
		sR uint64
		sG uint64
		sB uint64
	)
	game := strings.Split(line, ": ")[1]
	sets := strings.Split(game, "; ")

	for _, b := range sets {
		cubes := strings.Split(b, ", ")
		for i := 0; i < len(cubes); i++ {
			aux, color := strings.Split(cubes[i], " ")[0], strings.Split(cubes[i], " ")[1]
			quantity, _ := strconv.ParseUint(aux, 10, 64)
			if color == "red" {
				sR += quantity

			} else if color == "green" {
				sG += quantity

			} else {
				sB += quantity

			}
			if sR > 12 || sG > 13 || sB > 14 {
				return false
			}
		}
		sR, sB, sG = 0, 0, 0
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
		fmt.Println(line)
		arr = append(arr, cube_conumdrum(line))
	}
	for i, b := range arr {
		if b {
			sum += i + 1
		}
	}
	fmt.Println("Part 1: ", sum)
}
