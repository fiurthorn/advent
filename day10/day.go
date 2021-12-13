package day10

import (
	_ "embed"
	"fmt"
	"log"
	"sort"

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
	return fmt.Sprintf("1:%v 2:%v",
		d.process1(lines),
		d.process2(lines),
	)
}

func (d Day) process1(lines []string) string {
	q := lib.NewQueue()

	sum := 0
	for _, line := range lines {
		q.Clear()
	out:
		for _, r := range line {
			switch r {
			case '(':
				q.Push(')')
			case '[':
				q.Push(']')
			case '{':
				q.Push('}')
			case '<':
				q.Push('>')
			case ')':
				op := q.Pop()
				if op != r {
					sum += 3
					q.Clear()
					break out
				}
			case ']':
				op := q.Pop()
				if op != r {
					sum += 57
					q.Clear()
					break out
				}
			case '}':
				op := q.Pop()
				if op != r {
					sum += 1197
					q.Clear()
					break out
				}
			case '>':
				op := q.Pop()
				if op != r {
					sum += 25137
					q.Clear()
					break out
				}
			}
		}
	}
	return fmt.Sprintf("%v", sum)
}

func (d Day) process2(lines []string) string {
	q := lib.NewQueue()

	sums := []int{}
	for _, line := range lines {
		q.Clear()
	out:
		for _, r := range line {
			switch r {
			case '(':
				q.Push(')')
			case '[':
				q.Push(']')
			case '{':
				q.Push('}')
			case '<':
				q.Push('>')
			case ')':
				op := q.Pop()
				if op != r {
					q.Clear()
					break out
				}
			case ']':
				op := q.Pop()
				if op != r {
					q.Clear()
					break out
				}
			case '}':
				op := q.Pop()
				if op != r {
					q.Clear()
					break out
				}
			case '>':
				op := q.Pop()
				if op != r {
					q.Clear()
					break out
				}
			}
		}
		score := 0
		if len(*q) > 0 {
			for len(*q) > 0 {
				el := q.Pop()
				score *= 5
				switch el {
				case ')':
					score += 1
				case ']':
					score += 2
				case '}':
					score += 3
				case '>':
					score += 4
				}
			}
			sums = append(sums, score)
		}
	}
	sort.Ints(sums)
	return fmt.Sprintf("%v", sums[len(sums)/2])
}
