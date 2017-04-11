// https://projecteuler.net/problem=13
// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const stringMaxLength = 10

func readTextFile(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	result := strings.Split(str, "\n")

	return result
}

func getSignificantDigit(partialSum uint64) uint64 {
	partialResult := strconv.FormatUint(partialSum, 10)
	lastPos := len(partialResult) - stringMaxLength
	result, _ := strconv.ParseUint(partialResult[0:lastPos], 10, 64)
	return result
}

func isLastIteration(i int) bool {
	if i > 10 {
		return false
	}
	return true
}

func main() {
	var lines = readTextFile("data.txt")
	var partialSum uint64

	for i := len(lines[0]); i > 0; i -= stringMaxLength {
		for _, line := range lines {
			startPos := i - stringMaxLength
			if startPos < 0 {
				startPos = 0
			}

			lineValue, _ := strconv.ParseUint(line[startPos:i], 10, 64)
			partialSum += lineValue
		}

		result := getSignificantDigit(partialSum)

		if isLastIteration(i) {
			partialSum += result
		} else {
			partialSum = result
		}
	}

	fmt.Printf("Result: %s\n", strconv.FormatUint(partialSum, 10)[0:10])
}
