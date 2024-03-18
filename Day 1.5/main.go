package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// open the file using Open() function from os library
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(?P<one>one)|(?P<two>two)|(?P<three>three)|(?P<four>four)|(?P<five>five)|(?P<six>six)|(?P<seven>seven)|(?P<eight>eight)|(?P<nine>nine)|(?P<ten>ten)|(?P<digit>\d)`)

	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)

		sumRow(matches)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	fmt.Printf("Result is %v \n", totalSum)
}

var totalSum uint

func sumRow(slice []string) {
	num1 := convertToDigit(slice[0])
	num2 := convertToDigit(slice[len(slice)-1])

	totalSum += num1*10 + num2
}

func convertToDigit(digit string) uint {

	switch digit {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	default:
		log.Fatal("something went wrong")
	}
	return 0
}
