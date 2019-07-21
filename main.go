package main

import (
	"fmt"
)

func p(t int) string {
	h := t / 60
	m := t % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

func main() {
	cfg := newConfig()
	var workCount, workTime, relaxTime int
	for _, item := range calc(cfg) {
		switch item.typ {
		case WORK:
			workCount++
			workTime += item.elapsed
			fmt.Printf("%-20s%v\n", "рабочий интервал", p(item.totaltime))
		case SMALL:
			relaxTime += item.elapsed
			fmt.Printf("%-20s%v\n", "короткий перерыв", p(item.totaltime))
		case LARGE:
			relaxTime += item.elapsed
			fmt.Printf("%-20s%v\n", "БОЛЬШОЙ ПЕРЕРЫВ ---", p(item.totaltime))
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v (%d)\n", "работа:", p(workTime), workCount)
	fmt.Printf("%-20s%v\n", "отдых:", p(relaxTime))
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v\n", "полное время:", p(workTime+relaxTime))
}
