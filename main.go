package main

import (
	"flag"
	"log"
	"os"
	"time"

	year2021 "github.com/fiurthorn/advent/2021"
	"github.com/fiurthorn/advent/lib"
)

var (
	day, year int

	all  bool
	args = flag.NewFlagSet("AoC", flag.ExitOnError)

	solutions = map[int][]lib.Solution{
		2021: year2021.Solutions,
	}
)

func init() {
	log.SetFlags(0)
	args.IntVar(&year, "year", 2021, "year to choose [2021]")
	args.Func("day", "day to choose [1-25|*]", parseDay)
	args.Parse(os.Args[1:])
}

func main() {
	gstart := time.Now()
	defer log.Printf("%d all : %v", year, time.Since(gstart))
	if all {
		for index, solution := range solutions[year] {
			if solution != nil {
				log.Printf("start: %2d", index+1)
				start := time.Now()
				solution.Run()
				log.Printf(" end : %2d in %v", index+1, time.Since(start))
			}
		}
	} else if day > 0 && day <= len(solutions[year]) {
		if solution := solutions[year][day-1]; solution != nil {
			log.Printf("start: %2d", day)
			start := time.Now()
			solution.Run()
			log.Printf("%d end : %2d in %v", year, day, time.Since(start))
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
