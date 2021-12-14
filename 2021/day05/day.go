package day05

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

// - count max x und max y
// - create array with [size=x*y]int
// - search only for x1==x2 || y1==y2
// - a2-a1
//   - 0 -> 0
//   - >1 -> +1
//   - <1 -> -1
// increment each point
// count overlaps -> point >= 2

type Point struct {
	x1, y1 int
	x2, y2 int
}

func NewPoint(line string) *Point {
	point := strings.Split(line, " -> ")
	left := strings.Split(point[0], ",")
	right := strings.Split(point[1], ",")

	return &Point{
		x1: lib.Atoi(left[0]),
		y1: lib.Atoi(left[1]),
		x2: lib.Atoi(right[0]),
		y2: lib.Atoi(right[1]),
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("%d:%d -> %d:%d", p.x1, p.y1, p.x2, p.y2)
}

func (p *Point) Equals(x, y int) bool {
	return x == p.x2 && y == p.y2
}

func (p *Point) Dx() int {
	if p.x2 > p.x1 {
		return 1
	}
	if p.x1 > p.x2 {
		return -1
	}
	return 0
}

func (p *Point) Dy() int {
	if p.y2 > p.y1 {
		return 1
	}
	if p.y1 > p.y2 {
		return -1
	}
	return 0
}

type Field struct {
	Size       int
	MaxX, MaxY int
	Points     []*Point
	Landscape  []int
}

func NewField(lines []string, second bool) *Field {
	f := Field{}
	for _, line := range lines {
		p := NewPoint(line)
		if p.x1 > f.MaxX {
			f.MaxX = p.x1
		}
		if p.x2 > f.MaxX {
			f.MaxX = p.x2
		}
		if p.y1 > f.MaxY {
			f.MaxY = p.y1
		}
		if p.y2 > f.MaxY {
			f.MaxY = p.y2
		}
		f.Points = append(f.Points, p)
	}
	f.Size = f.MaxX * (f.MaxY + 1)
	f.Landscape = make([]int, f.Size)
	f.insertData()
	if second {
		return f.insertData2()
	}
	return (&f)
}

func (f *Field) insertData() *Field {
	for _, p := range f.Points {
		if p.x1 == p.x2 || p.y1 == p.y2 {
			for x, y := p.x1, p.y1; !p.Equals(x, y); x, y = x+p.Dx(), y+p.Dy() {
				f.Landscape[y*f.MaxX+x]++
			}
			f.Landscape[p.y2*f.MaxX+p.x2]++
		}
	}
	return f
}

func (f *Field) diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (f *Field) insertData2() *Field {
	for _, p := range f.Points {
		if f.diff(p.x1, p.x2) == f.diff(p.y1, p.y2) {
			for x, y := p.x1, p.y1; !p.Equals(x, y); x, y = x+p.Dx(), y+p.Dy() {
				f.Landscape[y*f.MaxX+x]++
			}
			f.Landscape[p.y2*f.MaxX+p.x2]++
		}
	}
	return f
}

func (f *Field) String() string {
	sb := strings.Builder{}

	sb.WriteRune('\n')
	for i := 0; i < f.Size; i += f.MaxX {
		sb.WriteString(fmt.Sprintf("%v", f.Landscape[i:i+f.MaxX]))
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (d Day) Run() {
	log.Printf("Example:  %v", d.process(dayExample))
	log.Printf("Solution: %v", d.process(dayData))
}

func (d Day) process(data string) string {
	lines := lib.Lines(data)

	return fmt.Sprintf("1:%v, 2:%v",
		d.process1(lines),
		d.process2(lines),
	)
}

func (d Day) process1(lines []string) string {
	f := NewField(lines, false)
	count := 0
	for _, v := range f.Landscape {
		if v >= 2 {
			count++
		}
	}
	return fmt.Sprintf("2-count: %d", count)
}

func (d Day) process2(lines []string) string {
	f := NewField(lines, true)
	count := 0
	for _, v := range f.Landscape {
		if v >= 2 {
			count++
		}
	}
	return fmt.Sprintf("2-count: %d", count)
}
