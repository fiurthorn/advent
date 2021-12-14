package day04

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

type Field struct {
	value int
	drawn bool
}

func (b Field) String() string {
	return fmt.Sprintf("[%2d {%5t}]", b.value, b.drawn)
}

type Board struct {
	fields []Field
}

func (b *Board) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v",
		b.fields[0:5],
		b.fields[5:10],
		b.fields[10:15],
		b.fields[15:20],
		b.fields[20:25],
	)
}

func (b *Board) Sum() (sum int) {
	for i := 0; i < len(b.fields); i++ {
		if !b.fields[i].drawn {
			sum += b.fields[i].value
		}
	}
	return
}

func (b *Board) Draw(num int) {
	for i := 0; i < len(b.fields); i++ {
		if b.fields[i].value == num {
			b.fields[i].drawn = true
		}
	}
}

func (b *Board) Won() bool {
	result := false
	for i := 0; i < 5; i++ {
		result = b.or(result, b.check(i, 0, 0, 1))
		result = b.or(result, b.check(0, i, 1, 0))
	}
	return result
}

func (b *Board) check(x, y, dx, dy int) bool {
	result := true
	for i := 0; i < 5; i++ {
		result = b.and(result, b.fields[y*5+x].drawn)
		x, y = x+dx, y+dy
	}
	return result
}

func (b *Board) and(x, y bool) bool {
	return x && y
}

func (b *Board) or(x, y bool) bool {
	return x || y
}

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
	draws := lib.Numbers(lines[0])

	lines = lines[2:]
	boards := []*Board{}
	for i, length := 0, 1+len(lines)/6; i < length; i++ {
		values := lines[6*i : 6*i+5]
		board := d.parseBoard(strings.Join(values, " "))
		boards = append(boards, &board)
	}

	for _, draw := range draws {
		for i := 0; i < len(boards); i++ {
			boards[i].Draw(draw)
			if boards[i].Won() {
				sum := boards[i].Sum()
				return fmt.Sprintf("%v * %v = %v", draw, sum, draw*sum)
			}
		}
	}

	return ""
}

func (d Day) process2(lines []string) string {
	draws := lib.Numbers(lines[0])

	lines = lines[2:]
	won := d.last(lines, draws)

	boards := []*Board{}
	for i, length := 0, 1+len(lines)/6; i < length; i++ {
		values := lines[6*i : 6*i+5]
		board := d.parseBoard(strings.Join(values, " "))
		boards = append(boards, &board)
	}

	for _, draw := range draws {
		for i := 0; i < len(boards); i++ {
			boards[i].Draw(draw)
			if boards[i].Won() && i == won {
				sum := boards[i].Sum()
				return fmt.Sprintf("%v * %v = %v", draw, sum, draw*sum)
			}
		}
	}

	return ""

}

func (d Day) last(lines []string, draws []int) (current int) {
	boards := []*Board{}
	for i, length := 0, 1+len(lines)/6; i < length; i++ {
		values := lines[6*i : 6*i+5]
		board := d.parseBoard(strings.Join(values, " "))
		boards = append(boards, &board)
	}

	count := len(boards)
	for _, draw := range draws {
		for i := 0; i < len(boards); i++ {
			if boards[i] != nil {
				boards[i].Draw(draw)
				if boards[i].Won() {
					boards[i] = nil
					current = i
					count--
					if count == 0 {
						return
					}
				}
			}
		}
	}
	return
}

func (d Day) parseBoard(data string) (result Board) {
	values := strings.SplitAfter(data, " ")

	for _, value := range values {
		value = strings.TrimSpace(value)
		if len(value) > 0 {
			result.fields = append(result.fields, Field{value: lib.Atoi(value)})
		}
	}

	return
}
