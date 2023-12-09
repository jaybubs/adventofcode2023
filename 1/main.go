package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	// just digits
	/*
	for scanner.Scan() {
		input := scanner.Text()
		lastDigit := regexp.MustCompile(`.*(?P<digit>\d).*?`)
		firstDigit := regexp.MustCompile(`(?P<digit>\d).*`)
		lastMatch := lastDigit.FindStringSubmatch(input)
		firstMatch := firstDigit.FindStringSubmatch(input)
		lastIndex := lastDigit.SubexpIndex("digit")
		firstIndex := firstDigit.SubexpIndex("digit")
		lastConv := lastMatch[lastIndex]
		firstConv := firstMatch[firstIndex]
		concat, _ := strconv.Atoi(firstConv+lastConv)
		total += concat
	}
	*/

	intMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"zero": 0,
	}

	for scanner.Scan() {
		input := scanner.Text()
		lastDigit := regexp.MustCompile(`.*(?P<digit>\d|one|two|three|four|five|six|seven|eight|nine|zero).*?`)
		firstDigit := regexp.MustCompile(`(?P<digit>\d|one|two|three|four|five|six|seven|eight|nine|zero).*`)
		lastMatch := lastDigit.FindStringSubmatch(input)
		firstMatch := firstDigit.FindStringSubmatch(input)
		lastIndex := lastDigit.SubexpIndex("digit")
		firstIndex := firstDigit.SubexpIndex("digit")
		lastConv, err := strconv.Atoi(lastMatch[lastIndex])
		if err != nil {
			lastConv = intMap[lastMatch[lastIndex]]
		}
		firstConv, err := strconv.Atoi(firstMatch[firstIndex])
		if err != nil {
			firstConv = intMap[firstMatch[firstIndex]]
		}
		// try to convert straight to int
		concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstConv, lastConv))
		total += concat
	}

	fmt.Println(total)
	

}
