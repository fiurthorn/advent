package main

import (
	"flag"
	"log"
	"time"

	"github.com/fiurthorn/advent/day01"
	"github.com/fiurthorn/advent/day02"
)

type Solution interface {
	Run()
}

var (
	day       int
	solutions = []Solution{
		day01.Day{},
		day02.Day{},
	}
)

func init() {
	log.SetFlags(0)
	flag.IntVar(&day, "day", -1, "day to choose")
}

func main() {
	flag.Parse()

	gstart := time.Now()
	if day > 0 && day <= len(solutions) {
		if solution := solutions[day-1]; solution != nil {
			log.Printf("start: %2d", day)
			start := time.Now()
			solution.Run()
			log.Printf(" end : %2d in %v", day, time.Since(start))
		}
	} else {
		for index, solution := range solutions {
			if solution != nil {
				log.Printf("start: %2d", index+1)
				start := time.Now()
				solution.Run()
				log.Printf(" end : %2d in %v", index+1, time.Since(start))
			}
		}
	}

	log.Printf(" all : %v", time.Since(gstart))
}