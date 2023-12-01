package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to open file ")
	}
	defer f.Close()

	var calVals []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		calVals = append(calVals, calcCalibration(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int
	for _, v := range calVals {
		result += v
	}
	fmt.Printf("Result: %v\n", result)

}

func calcCalibration(input string) int {
	var first, last int

	for _, ch := range input {
		i, err := strconv.Atoi(string(ch))
		if err != nil {
			continue
		}
		if first == 0 {
			first = i
			continue
		}
		last = i
	}
	if last == 0 {
		last = first
	}

	result, _ := strconv.Atoi(fmt.Sprintf("%v%v", first, last))
	log.Printf("Got first + last = %v + %v = %v", first, last, result)
	return result
}
