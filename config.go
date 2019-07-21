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

func newConfig() *config {

	c := &config{}

	flag.BoolVar(&c.mode, "w", false, "mode")
	flag.BoolVar(&c.verbose, "v", false, "verbose")
	flag.Parse()

	c.work = read(0)
	c.small = read(1)
	c.large = read(2)
	c.worklimit = read(3)
	c.timelimit = parsehhmm(flag.Arg(4))

	return c
}

func parsehhmm(s string) int {
	var hr, mi int
	fmt.Sscanf(s, "%d:%d", &hr, &mi)
	return hr*60 + mi
}

func read(arg int) int {
	v, err := strconv.ParseInt(flag.Arg(arg), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}
