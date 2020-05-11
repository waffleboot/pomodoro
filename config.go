package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type interval struct {
	start int
	end   int
}

type config struct {
	work      interval
	small     interval
	large     interval
	worklimit int
	timelimit int
	mode      bool
	verbose   bool
	time      bool
}

func usage() {
	fmt.Println("usage: pomodoro work small large count limit")
	fmt.Println("usage: pomodoro work limit")
	fmt.Println("usage: pomodoro limit")
}

func defaultConfig() *config {
	c := &config{}
	c.work = interval{25, 25}
	c.small = interval{5, 5}
	c.large = interval{25, 25}
	c.worklimit = 2
	c.timelimit = parsehhmm("8:00")
	return c
}

func newConfig() *config {

	c := defaultConfig()

	var help bool
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&c.mode, "w", false, "mode")
	flag.BoolVar(&c.verbose, "v", false, "verbose")
	flag.BoolVar(&c.time, "t", false, "show time")
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}
	if flag.NArg() == 0 {
		usage()
	} else if flag.NArg() == 1 {
		c.timelimit = parsehhmm(flag.Arg(0))
	} else if flag.NArg() == 2 {
		c.work = readInterval(0)
		c.timelimit = parsehhmm(flag.Arg(1))
	} else if flag.NArg() == 5 {
		c.work = readInterval(0)
		c.small = readInterval(1)
		c.large = readInterval(2)
		c.worklimit = read(3)
		c.timelimit = parsehhmm(flag.Arg(4))
	} else {
		usage()
		os.Exit(1)
	}

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

func readInterval(arg int) interval {
	var i interval
	if _, err := fmt.Sscanf(flag.Arg(arg), "%d-%d", &i.start, &i.end); err != nil {
		i.start = read(arg)
		i.end = i.start
	}
	return i
}
