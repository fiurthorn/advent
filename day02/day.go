package day02

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

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
	return fmt.Sprintf("simple:%v, extended: %v",
		d.process1(lines),
		d.process2(lines),
	)
}

func (d Day) process1(lines []string) string {
	var depth, horizontal int

	for _, line := range lines {
		values := strings.Split(line, " ")
		cmd, value := values[0], lib.Atoi(values[1])
		switch cmd {
		case "up":
			depth -= value
		case "down":
			depth += value
		case "forward":
			horizontal += value
		}
	}

	return fmt.Sprintf("%d", depth*horizontal)
}

func (d Day) process2(lines []string) string {
	var depth, horizontal, aim int

	for _, line := range lines {
		values := strings.Split(line, " ")
		cmd, value := values[0], lib.Atoi(values[1])
		switch cmd {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			horizontal += value
			depth += aim * value
		}
	}

	return fmt.Sprintf("%d*%d=%d", depth, horizontal, depth*horizontal)
}
