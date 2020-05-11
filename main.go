package main

import (
	"fmt"
	"time"
)

func p(t int) string {
	return fmt.Sprintf("%02d:%02d", t/60, t%60)
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
		if cfg.verbose {
			prefix(item, totalTime)
		}
		switch item.typ {
		case WORK:
			workCount++
			workTime += item.elapsed
			if cfg.verbose {
				if item.full {
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
		if cfg.verbose {
			fmt.Println()
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v\n", "полное время:", p(workTime+relaxTime))
	fmt.Printf("%-20s%v\n", "работа:", p(workTime))
	fmt.Printf("%-20s%v\n", "отдых:", p(relaxTime))
	fmt.Printf("%-20s%2v по %2d минут\n", "рабочих интервалов", workCount, cfg.work.start)
	fmt.Printf("%-20s%2v по %2d минут\n", "коротких перерывов", smallCount, cfg.small.start)
	if largeCount > 0 {
		fmt.Printf("%-20s%2v по %2d минут\n", "больших перерывов", largeCount, cfg.large.start)
	}
}
