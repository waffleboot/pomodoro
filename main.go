package main

import (
	"fmt"
	"time"
)

func p(t int) string {
	h := t / 60
	m := t % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

func main() {

	cfg := newConfig()
	var totalTime, workCount, workTime, relaxTime, smallCount, largeCount int
	now := time.Now()
	items := calc(cfg)
	for i := range items {
		item := &items[i]
		totalTime += item.elapsed
		t := now.Add(time.Duration(totalTime) * time.Minute).Format("15:04")
		switch item.typ {
		case WORK:
			workCount++
			workTime += item.elapsed
			if cfg.verbose {
				if item.elapsed == cfg.work {
					fmt.Printf("%-20s%v | %s\n", "рабочий интервал", p(totalTime), t)
				} else {
					fmt.Printf("%-20s%v | %s +%d\n", "рабочий интервал", p(totalTime), t, item.elapsed)
				}
			}
		case SMALL:
			smallCount++
			relaxTime += item.elapsed
			if cfg.verbose {
				fmt.Printf("%-20s%v | %s\n", "короткий перерыв", p(totalTime), t)
			}
		case LARGE:
			largeCount++
			relaxTime += item.elapsed
			if cfg.verbose {
				fmt.Printf("%-20s%v | %s\n", "БОЛЬШОЙ ПЕРЕРЫВ ---", p(totalTime), t)
			}
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v\n", "полное время:", p(workTime+relaxTime))
	fmt.Printf("%-20s%v\n", "работа:", p(workTime))
	fmt.Printf("%-20s%v\n", "отдых:", p(relaxTime))
	fmt.Printf("%-20s%v\n", "рабочих интервалов", workCount)
	fmt.Printf("%-20s%v\n", "коротких перерывов", smallCount)
	if largeCount > 0 {
		fmt.Printf("%-20s%v\n", "больших перерывов", largeCount)
	}
}
