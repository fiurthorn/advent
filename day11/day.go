package day11

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
	return fmt.Sprintf("1:%v 2:%v",
		d.process1(lines),
		d.process2(lines),
	)
}

type Field struct {
	Points  [][]*Point
	RowSize int
}

func (f *Field) String() string {
	sb := strings.Builder{}
	for _, line := range f.Points {
		for _, p := range line {
			sb.WriteString(p.String())
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (f *Field) positions(x, y int) (result [][2]int) {
	check := [2]int{x, y - 1}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x - 1, y}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x + 1, y}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x, y + 1}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}

	check = [2]int{x - 1, y - 1}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x + 1, y + 1}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x - 1, y + 1}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x + 1, y - 1}
	if check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}

	return
}

type Point struct {
	Level   int
	Flashed bool
	Flashes int
}

func (p *Point) String() string {
	a, z := ' ', ' '
	if p.Flashed {
		a, z = '(', ')'
	}
	return fmt.Sprintf("%c%d%c", a, p.Level, z)
}

func (p *Point) Incr() {
	p.Level++
}

func (p *Point) Flash() {
	p.Flashed = true
}

func (p *Point) Reset() {
	if p.Level > 9 {
		p.Level = 0
	}
	if p.Flashed {
		p.Flashes++
		p.Flashed = false
	}
}

func (d Day) process1(lines []string) string {
	f := &Field{Points: [][]*Point{}, RowSize: len(lines[0])}
	for _, line := range lines {
		p := []*Point{}
		for _, r := range line {
			p = append(p, &Point{Level: int(r - '0')})
		}
		f.Points = append(f.Points, p)
	}

	for i := 0; i < 100; i++ {
		for x := 0; x < f.RowSize; x++ {
			for y := 0; y < len(f.Points); y++ {
				curr := f.Points[y][x]
				curr.Incr()
			}
		}
		check := true
		for check {
			check = false
			for x := 0; x < f.RowSize; x++ {
				for y := 0; y < len(f.Points); y++ {
					curr := f.Points[y][x]
					if curr.Level > 9 && !curr.Flashed {
						curr.Flashed = true
						check = true
						for _, pos := range f.positions(x, y) {
							f.Points[pos[1]][pos[0]].Incr()
						}
					}
				}
			}
		}
		for x := 0; x < f.RowSize; x++ {
			for y := 0; y < len(f.Points); y++ {
				(f.Points[y][x]).Reset()
			}
		}
	}
	flashes := 0
	for x := 0; x < f.RowSize; x++ {
		for y := 0; y < len(f.Points); y++ {
			flashes += (f.Points[y][x]).Flashes
		}
	}

	return fmt.Sprintf("%v", flashes)
}

func (d Day) process2(lines []string) string {
	f := &Field{Points: [][]*Point{}, RowSize: len(lines[0])}
	for _, line := range lines {
		p := []*Point{}
		for _, r := range line {
			p = append(p, &Point{Level: int(r - '0')})
		}
		f.Points = append(f.Points, p)
	}

	i := 1
out:
	for ; ; i++ {
		for x := 0; x < f.RowSize; x++ {
			for y := 0; y < len(f.Points); y++ {
				curr := f.Points[y][x]
				curr.Incr()
			}
		}
		check := true
		for check {
			check = false
			for x := 0; x < f.RowSize; x++ {
				for y := 0; y < len(f.Points); y++ {
					curr := f.Points[y][x]
					if curr.Level > 9 && !curr.Flashed {
						curr.Flashed = true
						check = true
						for _, pos := range f.positions(x, y) {
							f.Points[pos[1]][pos[0]].Incr()
						}
					}
				}
			}
		}
		count := len(f.Points) * f.RowSize
		for x := 0; x < f.RowSize; x++ {
			for y := 0; y < len(f.Points); y++ {
				(f.Points[y][x]).Reset()
				if f.Points[y][x].Level == 0 {
					count--
				}
			}
		}
		if count == 0 {
			log.Println(f)
			break out
		}
	}

	return fmt.Sprintf("%v", i)
}
