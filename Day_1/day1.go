package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	filename := "day1.txt"
	// filename := "test.txt"
	fmt.Print("loading file\n")
	f, err := os.Open(filename)

	check(err)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	calibrationSum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		calibrationSum += getCalibration(line)
	}

	fmt.Printf("The total calibration is: %d\n", calibrationSum)
	f.Close()
}

func getCalibration(line string) (calibration int) {
	lastNum := 0
	numberFound := false
	secondNumberFound := false
	for _, ch := range line {
		if unicode.IsNumber(ch) && !numberFound {
			num, err := strconv.Atoi(string(ch))
			check(err)
			calibration = num
			numberFound = true
		} else if unicode.IsNumber(ch) && numberFound {
			num, err := strconv.Atoi(string(ch))
			check(err)
			lastNum = num
			secondNumberFound = true
		}
	}

	if secondNumberFound == true {
		calibration = (10 * calibration) + lastNum
	} else {
		calibration = (10 * calibration) + calibration
	}
	fmt.Printf(line+": %d\n", calibration)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
