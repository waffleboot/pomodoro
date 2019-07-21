package main

type typ int

const (
	WORK typ = iota
	SMALL
	LARGE
)

type item struct {
	typ     typ
	elapsed int
}

func calc(cfg *config) []item {
	result := make([]item, 0, 10)
	var work, total, workCount int

	var remainder func(int, int) int

	if cfg.mode {
		remainder = func(work int, total int) int {
			return cfg.timelimit - work
		}
	} else {
		remainder = func(work int, total int) int {
			return cfg.timelimit - total
		}
	}

	for {
		workCount++
		if remainder := remainder(work, total); remainder <= cfg.work {
			result = append(result, item{
				typ:     WORK,
				elapsed: remainder,
			})
			return result
		}
		work += cfg.work
		total += cfg.work
		result = append(result, item{
			typ:     WORK,
			elapsed: cfg.work,
		})
		relaxtype := SMALL
		relaxperiod := cfg.small
		if workCount == cfg.worklimit {
			relaxperiod = cfg.large
			relaxtype = LARGE
			workCount = 0
		}
		if remainder(work, total+relaxperiod) <= 0 {
			return result
		}
		total += relaxperiod
		result = append(result, item{
			typ:     relaxtype,
			elapsed: relaxperiod,
		})

	}
}
