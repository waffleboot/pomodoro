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

func (cfg *config) remainder(work int, total int) int {
	return cfg.timelimit - total
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
	var work, total, workCount int
	for {
		workCount++
		if remainder := cfg.remainder(work, total); remainder <= cfg.work {
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
			if cfg.remainder(work, total+cfg.large) <= 0 {
				return result
			}
			total += cfg.large
			result = append(result, item{
				typ:       LARGE,
				elapsed:   cfg.large,
				totaltime: total,
			})
		} else {
			if cfg.remainder(work, total+cfg.small) <= 0 {
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
