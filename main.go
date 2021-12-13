package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fiurthorn/advent/day01"
	"github.com/fiurthorn/advent/day02"
	"github.com/fiurthorn/advent/day03"
	"github.com/fiurthorn/advent/day04"
	"github.com/fiurthorn/advent/day05"
	"github.com/fiurthorn/advent/day06"
	"github.com/fiurthorn/advent/day07"
	"github.com/fiurthorn/advent/day08"
	"github.com/fiurthorn/advent/day09"
	"github.com/fiurthorn/advent/day10"
	"github.com/fiurthorn/advent/day11"
	"github.com/fiurthorn/advent/day12"
	"github.com/fiurthorn/advent/day13"
	"github.com/fiurthorn/advent/lib"
)

type Solution interface {
	Run()
}

var (
	day  int
	all  bool
	args = flag.NewFlagSet("AoC", flag.ExitOnError)

	solutions = []Solution{
		day01.Day{}, day02.Day{}, day03.Day{}, day04.Day{}, day05.Day{},
		day06.Day{}, day07.Day{}, day08.Day{}, day09.Day{}, day10.Day{},
		day11.Day{}, day12.Day{}, day13.Day{},
	}
)

func init() {
	log.SetFlags(0)
	args.Func("day", fmt.Sprintf("day to choose [1-%d|*]", len(solutions)), parseDay)
	args.BoolVar(&all, "all", false, "day to choose")
	args.Parse(os.Args[1:])
}

func main() {
	gstart := time.Now()
	defer log.Printf(" all : %v", time.Since(gstart))
	if all {
		for index, solution := range solutions {
			if solution != nil {
				log.Printf("start: %2d", index+1)
				start := time.Now()
				solution.Run()
				log.Printf(" end : %2d in %v", index+1, time.Since(start))
			}
		}
	} else if day > 0 && day <= len(solutions) {
		if solution := solutions[day-1]; solution != nil {
			log.Printf("start: %2d", day)
			start := time.Now()
			solution.Run()
			log.Printf(" end : %2d in %v", day, time.Since(start))
		}
	} else {
		args.Usage()
	}
}

func parseDay(value string) error {
	if value == "all" || value == "*" {
		all = true
		return nil
	}

	day = lib.Atoi(value)
	if day <= 0 || day > len(solutions) {
		return flag.ErrHelp
	}

	return nil
}
