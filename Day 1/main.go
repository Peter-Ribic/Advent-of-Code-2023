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
	//	var totalSum uint64

	for scanner.Scan() {

		re := regexp.MustCompile(`\d`)
		matches := re.FindAllString(scanner.Text(), -1)
		fmt.Println(scanner.Text(), matches)
		number := 10 * uint64(matches[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

}
