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

	calibration_sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		calibration_sum += get_calibration(line)
	}

	fmt.Printf("The total calibration is: %d", calibration_sum)
	f.Close()
}

func get_calibration(line string) (calibration int) {
	calibration = 0
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
	fmt.Println(line+": %d", calibration)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
