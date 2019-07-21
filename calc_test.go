package main

import (
	"testing"
)

func TestTimeLimit(t *testing.T) {
	cfg := &config{
		work:      30,
		small:     5,
		large:     15,
		worklimit: 2,
		timelimit: parsehhmm("2:10"),
	}
	var workingCount, workingTime, relaxingTime int
	for _, item := range calc(cfg) {
		if item.typ == WORK {
			workingCount++
			workingTime += item.elapsed
		} else {
			relaxingTime += item.elapsed
		}
	}
	if workingCount != 4 {
		t.Error(workingCount)
	}
	if workingTime != parsehhmm("1:45") {
		t.Error(workingTime)
	}
	if relaxingTime != parsehhmm("0:25") {
		t.Error(relaxingTime)
	}
}

func TestWorkLimit(t *testing.T) {
	cfg := &config{
		work:      30,
		small:     5,
		large:     15,
		worklimit: 2,
		timelimit: parsehhmm("2:10"),
		mode:      true,
	}
	var workingCount, workingTime, relaxingTime int
	for _, item := range calc(cfg) {
		if item.typ == WORK {
			workingCount++
			workingTime += item.elapsed
		} else {
			relaxingTime += item.elapsed
		}
	}
	if workingCount != 5 {
		t.Error(workingCount)
	}
	if workingTime != parsehhmm("2:10") {
		t.Error(workingTime)
	}
	if relaxingTime != parsehhmm("0:40") {
		t.Error(relaxingTime)
	}
}
