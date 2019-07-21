package main

import (
	"fmt"
	"os"
	"strconv"
)

type config struct {
	work      int
	small     int
	large     int
	worklimit int
	timelimit int
}

func newConfig() *config {
	c := &config{}
	read := func(arg int) int {
		v, err := strconv.ParseInt(os.Args[arg], 10, 64)
		if err != nil {
			panic(err)
		}
		return int(v)
	}
	c.work = read(1)
	c.small = read(2)
	c.large = read(3)
	c.worklimit = read(4)

	var hr, mi int
	fmt.Sscanf(os.Args[5], "%d:%d", &hr, &mi)
	c.timelimit = hr*60 + mi
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
	var total, workCount int
	for {
		if total+cfg.work < cfg.timelimit {
			total += cfg.work
			result = append(result, item{
				typ:       WORK,
				elapsed:   cfg.work,
				totaltime: total,
			})
		} else {
			remainder := cfg.timelimit - total
			total += remainder
			result = append(result, item{
				typ:       WORK,
				elapsed:   remainder,
				totaltime: total,
			})
			return result
		}
		workCount++
		if workCount == cfg.worklimit {
			workCount = 0
			total += cfg.large
			if total >= cfg.timelimit {
				return result
			}
			result = append(result, item{
				typ:       LARGE,
				elapsed:   cfg.large,
				totaltime: total,
			})
		} else {
			total += cfg.small
			if total >= cfg.timelimit {
				return result
			}
			result = append(result, item{
				typ:       SMALL,
				elapsed:   cfg.small,
				totaltime: total,
			})
		}
	}
}
