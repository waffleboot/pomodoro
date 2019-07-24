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

func prefix(item *item, totalTime int) {
	switch item.typ {
	case WORK:
		fmt.Printf("%-20s%v", "рабочий интервал", p(totalTime))
	case SMALL:
		fmt.Printf("%-20s%v", "короткий перерыв", p(totalTime))
	case LARGE:
		fmt.Printf("%-20s%v", "БОЛЬШОЙ ПЕРЕРЫВ ---", p(totalTime))
	}
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
		prefix(item, totalTime)
		switch item.typ {
		case WORK:
			workCount++
			workTime += item.elapsed
			if cfg.verbose {
				if item.elapsed == cfg.work {
					if cfg.time {
						fmt.Printf(" | %s", t)
					}
				} else {
					if cfg.time {
						fmt.Printf(" | %s +%d", t, item.elapsed)
					} else {
						fmt.Printf(" +%d", item.elapsed)
					}
				}
			}
		case SMALL:
			smallCount++
			relaxTime += item.elapsed
			if cfg.verbose && cfg.time {
				fmt.Printf(" | %s", t)
			}
		case LARGE:
			largeCount++
			relaxTime += item.elapsed
			if cfg.verbose && cfg.time {
				fmt.Printf(" | %s", t)
			}
		}
		fmt.Println()
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
