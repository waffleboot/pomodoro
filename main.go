package main

import (
	"fmt"
	// "net/http"
	// "os"
	// "os/signal"
	// "syscall"
)

func p(t int) string {
	h := t / 60
	m := t % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

func main() {

	// graceful shutdown

	// srv := http.Server{}
	// term := make(chan os.Signal, 1)
	// exit := make(chan struct{})

	// signal.Notify(term, syscall.SIGTERM)
	// go func() {
	// 	<-term
	// 	srv.Shutdown(ctx)
	// 	close(exit)
	// }()

	// err := srv.ListenAndServe()
	// if err == http.ErrServerClosed {
	// 	<-exit
	// }

	cfg := newConfig()
	var totalTime, workCount, workTime, relaxTime, smallCount, largeCount int
	items := calc(cfg)
	for i := range items {
		item := &items[i]
		totalTime += item.elapsed
		switch item.typ {
		case WORK:
			workCount++
			workTime += item.elapsed
			if cfg.verbose {
				if item.elapsed == cfg.work {
					fmt.Printf("%-20s%v\n", "рабочий интервал", p(totalTime))
				} else {
					fmt.Printf("%-20s%v (%d)\n", "рабочий интервал", p(totalTime), item.elapsed)
				}
			}
		case SMALL:
			smallCount++
			relaxTime += item.elapsed
			if cfg.verbose {
				fmt.Printf("%-20s%v\n", "короткий перерыв", p(totalTime))
			}
		case LARGE:
			largeCount++
			relaxTime += item.elapsed
			if cfg.verbose {
				fmt.Printf("%-20s%v\n", "БОЛЬШОЙ ПЕРЕРЫВ ---", p(totalTime))
			}
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v (%d)\n", "работа:", p(workTime), workCount)
	fmt.Printf("%-20s%v (%d/%d)\n", "отдых:", p(relaxTime), smallCount, largeCount)
	fmt.Println("-------------------------")
	fmt.Printf("%-20s%v\n", "полное время:", p(workTime+relaxTime))
}
