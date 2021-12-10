package day01

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/fiurthorn/advent/lib"
)

type Day struct{}

//go:embed example.txt
var dayExample string

//go:embed data.txt
var dayData string

func (d Day) Run() {
	log.Printf("Example:  %v", d.process(dayExample))
	log.Printf("Solution: %v", d.process(dayData))
}

func (d Day) process(data string) string {
	lines := lib.Lines(data)
	return fmt.Sprintf("count:%v, count3:%v",
		d.process1(lines),
		d.process3(lines),
	)
}

func (d Day) process1(lines []string) int {
	prev, count := -1, 0
	for _, line := range lines {
		current := lib.Atoi(line)
		if prev > 0 && current > prev {
			count++
		}
		prev = current
	}
	return count
}

func (d Day) process3(lines []string) int {
	prev, count := -1, 0
	for i, len := 0, len(lines)-2; i < len; i++ {
		current := lib.Atoi(lines[i]) + lib.Atoi(lines[i+1]) + lib.Atoi(lines[i+2])
		if prev > 0 && current > prev {
			count++
		}
		prev = current
	}
	return count
}
