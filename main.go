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

func main() {
	cfg := newConfig()
	var workingTime, relaxingTime int
	for _, item := range calc(cfg) {
		switch item.typ {
		case WORK:
			workingTime += item.elapsed
			fmt.Printf("%-20s%v\n", "рабочий интервал", p(item.totaltime))
		case SMALL:
			relaxingTime += item.elapsed
			fmt.Printf("%-20s%v\n", "короткий перерыв", p(item.totaltime))
		case LARGE:
			relaxingTime += item.elapsed
			fmt.Printf("%-20s%v\n", "БОЛЬШОЙ ПЕРЕРЫВ ---", p(item.totaltime))
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v\n", "рабочий интервал:", p(workingTime))
	fmt.Printf("%-20s%v\n", "отдых:", p(relaxingTime))
}
