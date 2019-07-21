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

func makeRemainder(cfg *config) func(int, int) int {
	if cfg.mode {
		return func(work int, total int) int {
			return cfg.timelimit - work
		}
	}
	return func(work int, total int) int {
		return cfg.timelimit - total
	}
}

func calc(cfg *config) []item {
	items := make([]item, 0, 10)
	var work, total, workCount int

	remainder := makeRemainder(cfg)

	for {
		workCount++
		if remainder := remainder(work, total); remainder <= cfg.work {
			items = append(items, item{
				typ:     WORK,
				elapsed: remainder,
			})
			return items
		}
		work += cfg.work
		total += cfg.work
		items = append(items, item{
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
		total += relaxperiod
		if remainder(work, total) <= 0 {
			return items
		}
		items = append(items, item{
			typ:     relaxtype,
			elapsed: relaxperiod,
		})

	}
}
