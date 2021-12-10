package main

import (
	"log"
	"time"
)

type Solution interface {
	Run()
}

var (
	solutions = [25]Solution{}
)

func init() {
	log.SetFlags(0)
}

func main() {
	gstart := time.Now()
	for index, solution := range solutions {
		if solution != nil {
			log.Printf("start: %2d", index+1)
			start := time.Now()
			solution.Run()
			log.Printf(" end : %2d in %v", index+1, time.Since(start))
		}
	}
	log.Printf(" all : %v", time.Since(gstart))
}
