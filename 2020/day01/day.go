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
		d.process2(lines),
	)
}

func (d Day) process1(lines []string) int {
	numbers := []int{}
	for _, n := range lines {
		numbers = append(numbers, lib.Atoi(n))
	}

	var x, y int
out:
	for a := 0; a < len(numbers); a++ {
		for b := 0; b < len(numbers); b++ {
			if numbers[a]+numbers[b] == 2020 {
				x = numbers[a]
				y = numbers[b]
				break out
			}
		}
	}

	return x * y
}

func (d Day) process2(lines []string) int {
	numbers := []int{}
	for _, n := range lines {
		numbers = append(numbers, lib.Atoi(n))
	}

	var x, y, z int
out:
	for a := 0; a < len(numbers); a++ {
		for b := 0; b < len(numbers); b++ {
			for c := 0; c < len(numbers); c++ {
				if numbers[a]+numbers[b]+numbers[c] == 2020 {
					x = numbers[a]
					y = numbers[b]
					z = numbers[c]
					break out
				}
			}
		}
	}

	return x * y * z
}
