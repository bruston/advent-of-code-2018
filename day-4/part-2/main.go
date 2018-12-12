package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const (
	actionWake byte = iota
	actionSleep
	actionStartShift
)

type action struct {
	actionID byte
	time     time.Time
	guard    int
}

var errBadInput = errors.New("bad input")

func parseAction(s string) (action, error) {
	const timestampLength = 18
	if len(s) <= timestampLength {
		return action{}, errBadInput
	}
	t, err := time.Parse("[2006-01-02 15:04]", s[:timestampLength])
	if err != nil {
		return action{}, fmt.Errorf("bad date format: %v", err)
	}
	a := action{time: t}
	if s[timestampLength:] == " falls asleep" {
		a.actionID = actionSleep
		return a, nil
	}
	if s[timestampLength:] == " wakes up" {
		a.actionID = actionWake
		return a, nil
	}
	var id int
	if _, err := fmt.Sscanf(s[timestampLength:], " Guard #%d begins shift", &id); err != nil {
		return action{}, errBadInput
	}
	a.actionID = actionStartShift
	a.guard = id
	return a, nil
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	actions := []action{}
	for s.Scan() {
		a, err := parseAction(s.Text())
		if err != nil {
			log.Fatalf("error while scanning input file: %v", err)
		}
		actions = append(actions, a)
	}
	if err != nil {
		log.Fatalf("error while reading input file: %v", err)
	}

	sort.Slice(actions, func(i, j int) bool { return actions[i].time.Before(actions[j].time) })

	var lastID int
	naps := make(map[int][]int)
	for i, v := range actions {
		if v.actionID == actionStartShift {
			lastID = actions[i].guard
		} else {
			actions[i].guard = lastID
		}
		if v.actionID == actionWake {
			if actions[i-1].actionID != actionSleep {
				log.Fatalf("input format error, guard %d cannot wake up if they were not asleep!", v.guard)
			}
			if naps[actions[i].guard] == nil {
				naps[actions[i].guard] = make([]int, 60)
			}
			for minute := actions[i-1].time.Minute(); minute < v.time.Minute(); minute++ {
				naps[actions[i].guard][minute]++
			}
		}
	}
	var mostNapped int
	var bestChance int
	var sleepyGuard int
	for k, v := range naps {
		for i, v := range v {
			if v > mostNapped {
				mostNapped = v
				bestChance = i
				sleepyGuard = k
			}
		}
	}

	fmt.Println(sleepyGuard * bestChance)
}
