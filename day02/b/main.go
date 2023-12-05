package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cube_conumdrum(line string) uint64 {
	var (
		maxR uint64
		maxG uint64
		maxB uint64
	)
	game := strings.Split(line, ": ")[1]
	sets := strings.Split(game, "; ")

	for _, b := range sets {
		cubes := strings.Split(b, ", ")
		for i := 0; i < len(cubes); i++ {
			aux, color := strings.Split(cubes[i], " ")[0], strings.Split(cubes[i], " ")[1]
			quantity, _ := strconv.ParseUint(aux, 10, 64)
			if color == "red" {
				if maxR < quantity {
					maxR = quantity
				}
			} else if color == "green" {
				if maxG < quantity {
					maxG = quantity
				}
			} else {
				if maxB < quantity {
					maxB = quantity
				}
			}
		}
	}
	return maxR * maxG * maxB
}
func main() {
	var (
		arr []uint64
		sum uint64
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
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Part 2: ", sum)
}
