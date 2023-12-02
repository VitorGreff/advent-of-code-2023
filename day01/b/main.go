package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func replace_words(line *string) {
	*line = strings.ReplaceAll(*line, "one", "on1e")
	*line = strings.ReplaceAll(*line, "two", "tw2o")
	*line = strings.ReplaceAll(*line, "three", "thre3e")
	*line = strings.ReplaceAll(*line, "four", "fou4r")
	*line = strings.ReplaceAll(*line, "five", "fiv5e")
	*line = strings.ReplaceAll(*line, "six", "si6x")
	*line = strings.ReplaceAll(*line, "seven", "seve7n")
	*line = strings.ReplaceAll(*line, "eight", "eigh8t")
	*line = strings.ReplaceAll(*line, "nine", "nin9e")
}

func eval_first(line string) byte {
	for i := 0; i < len(line); i++ {
		if _, err := strconv.ParseUint(string(line[i]), 10, 64); err != nil {
			continue
		}
		return line[i]
	}
	return ' '
}

func eval_last(line string) byte {
	for i := len(line) - 1; i >= 0; i-- {
		if _, err := strconv.ParseUint(string(line[i]), 10, 64); err != nil {
			continue
		}
		return line[i]
	}
	return ' '
}

func main() {
	var (
		sum   uint64
		value uint64
		line  string
	)

	file, err := os.Open("../1.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		replace_words(&line)
		value, _ = strconv.ParseUint(string([]byte{eval_first(line), eval_last(line)}), 10, 64)
		sum += value
	}
	fmt.Println(sum)
}
