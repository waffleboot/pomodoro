package main

import (
	"flag"
	"fmt"
	"strconv"
)

type config struct {
	work      int
	small     int
	large     int
	worklimit int
	timelimit int
	mode      bool
	verbose   bool
}

func parsehhmm(s string) int {
	var hr, mi int
	fmt.Sscanf(s, "%d:%d", &hr, &mi)
	return hr*60 + mi
}

func newConfig() *config {

	c := &config{}

	flag.BoolVar(&c.mode, "w", false, "mode")
	flag.BoolVar(&c.verbose, "v", false, "verbose")
	flag.Parse()

	read := func(arg int) int {
		v, err := strconv.ParseInt(flag.Arg(arg), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(v)
	}
	c.work = read(0)
	c.small = read(1)
	c.large = read(2)
	c.worklimit = read(3)
	c.timelimit = parsehhmm(flag.Arg(4))

	return c
}

type typ int

const (
	WORK typ = iota
	SMALL
	LARGE
)

type item struct {
	typ       typ
	elapsed   int
	totaltime int
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
				typ:       WORK,
				elapsed:   remainder,
				totaltime: total + remainder,
			})
			return result
		}
		work += cfg.work
		total += cfg.work
		result = append(result, item{
			typ:       WORK,
			elapsed:   cfg.work,
			totaltime: total,
		})
		if workCount == cfg.worklimit {
			workCount = 0
			if remainder(work, total+cfg.large) <= 0 {
				return result
			}
			total += cfg.large
			result = append(result, item{
				typ:       LARGE,
				elapsed:   cfg.large,
				totaltime: total,
			})
		} else {
			if remainder(work, total+cfg.small) <= 0 {
				return result
			}
			total += cfg.small
			result = append(result, item{
				typ:       SMALL,
				elapsed:   cfg.small,
				totaltime: total,
			})
		}
	}
}
