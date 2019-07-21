package main

import (
	"testing"
)

func TestTimeLimit(t *testing.T) {
	cfg := &config{
		work:      25,
		small:     5,
		large:     15,
		worklimit: 2,
		timelimit: parsehhmm("8:00"),
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
	if workingCount != 14 {
		t.Fail()
	}
	if workingTime != parsehhmm("5:50") {
		t.Fail()
	}
	if relaxingTime != parsehhmm("2:05") {
		t.Fail()
	}
}

func TestWorkLimit(t *testing.T) {
	cfg := &config{
		work:      25,
		small:     5,
		large:     15,
		worklimit: 2,
		mode:      true,
		timelimit: parsehhmm("8:00"),
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
	if workingCount != 20 {
		t.Fail()
	}
	if workingTime != parsehhmm("8:00") {
		t.Fail()
	}
	if relaxingTime != parsehhmm("3:05") {
		t.Fail()
	}
}
