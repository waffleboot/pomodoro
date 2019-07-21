package main

import (
	"fmt"
	"os"
	"strconv"
)

func read(arg int) int {
	n, err := strconv.ParseInt(os.Args[arg], 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func p(t int) string {
	h := t / 60
	m := t % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

const (
	WORK = iota
	SMALL
	LARGE
)

func main() {
	var totalTime, workingTime, relaxingTime, workingCount, state int
	argWorkingTime, argSmallRelaxTime, argLargeRelaxTime, argLargeRelaxAfterEachWork := read(1), read(2), read(3), read(4)

	var hr, mi int
	fmt.Sscanf(os.Args[5], "%d:%d", &hr, &mi)
	m := hr*60 + mi

loop:
	for {
		switch state {
		case WORK:
			state = SMALL
			workingCount++
			if totalTime+argWorkingTime < m {
				workingTime += argWorkingTime
				totalTime += argWorkingTime
				if workingCount == argLargeRelaxAfterEachWork {
					workingCount = 0
					state = LARGE
				}
				fmt.Printf("%-20s%v\n", "рабочий интервал", p(totalTime))
			} else {
				workingTime += m - totalTime
				totalTime = m
				fmt.Printf("%-20s%v\n", "рабочий интервал", p(totalTime))
				break loop
			}
		case SMALL:
			state = WORK
			if totalTime+argSmallRelaxTime < m {
				relaxingTime += argSmallRelaxTime
				totalTime += argSmallRelaxTime
				fmt.Printf("%-20s%v\n", "короткий перерыв", p(totalTime))
			} else {
				break loop
			}
		case LARGE:
			state = WORK
			if totalTime+argLargeRelaxTime < m {
				relaxingTime += argLargeRelaxTime
				totalTime += argLargeRelaxTime
				fmt.Printf("%-20s%v\n", "БОЛЬШОЙ ПЕРЕРЫВ ---", p(totalTime))
			} else {
				break loop
			}
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v\n", "рабочий интервал:", p(workingTime))
	fmt.Printf("%-20s%v\n", "отдых:", p(relaxingTime))
}
