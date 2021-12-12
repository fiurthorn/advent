package day09

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

func (d Day) process2(lines []string) string {
	f := NewField(lines)
	m := f.SearchBasin()
	return fmt.Sprintf("%v:%v", m, m[0]*m[1]*m[2])
}

func (d Day) process1(lines []string) string {
	f := NewField(lines)
	result := 0
	for _, pos := range f.Search() {
		result += f.Points[pos.Y][pos.X].Value + 1
	}
	return fmt.Sprintf("%v", result)
}

type Field struct {
	Points  [][]Point
	RowSize int
}

type Point struct {
	Value int
	Index int
	Mark  int
}

func (p Point) String() string {
	return fmt.Sprintf("(%2d{%2d}[%2d])", p.Value, p.Mark, p.Index)
}

func (f Field) String() string {
	sb := strings.Builder{}
	for _, l := range f.Points {
		for _, p := range l {
			sb.WriteString(p.String())
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func NewField(lines []string) *Field {
	f := Field{}

	f.RowSize = len(lines[0])
	for y, line := range lines {
		row := make([]Point, f.RowSize)
		for i, c := range line {
			row[i] = Point{Value: lib.Atoi(string(c)), Index: y*f.RowSize + i, Mark: -1}
		}
		f.Points = append(f.Points, row)
	}

	return &f
}

type Pos struct {
	X, Y int
}

func (f *Field) SearchBasin() [3]int {
	max := [3]int{}

	size := len(f.Points)
	for i := 0; i < 10; i++ {
		for x := 0; x < f.RowSize; x++ {
			for y := 0; y < size; y++ {
				// x, y := 2, 2
				value := &f.Points[y][x]
				if value.Value == i {
					f.detect(f.RowSize*y+x, x, y, -1, -1, " ")
				}
			}
		}
	}

	result := make([]int, f.RowSize*size)
	for x := 0; x < f.RowSize; x++ {
		for y := 0; y < size; y++ {
			index := f.Points[y][x].Mark
			if index >= 0 {
				log.Println(index)
				result[index]++
			}
		}
	}

	for i := 0; i < len(result); i++ {
		val := result[i]
		if val > max[0] {
			max[2] = max[1]
			max[1] = max[0]
			max[0] = val
		} else if val > max[1] {
			max[2] = max[1]
			max[1] = val
		} else if val > max[2] {
			max[2] = val
		}
	}

	return max
}

func (f *Field) detect(basin, x, y int, ox, oy int, depth string) {
	value := &f.Points[y][x]
	if value.Mark >= 0 {
		return
	}
	nachbar := f.validates(x, y, ox, oy)
	// log.Println(depth, basin, value, nachbar)

	// log.Println(basin, depth, nachbar)
	for ni := 0; ni < len(nachbar); ni++ {
		n := nachbar[ni]
		nvalue := &f.Points[n[1]][n[0]]
		// log.Println(depth, basin, nvalue, value)
		// if nvalue.Value-value.Value == 1 {
		if nvalue.Value >= value.Value {
			if nvalue.Mark < 0 {
				value.Mark = basin
				// log.Println(depth, basin, n, value, nvalue)
				f.detect(basin, n[0], n[1], x, y, depth+" ")
			}
		}
	}
}

func (f *Field) Search() (result []Pos) {
	size := len(f.Points)

	for y := 0; y < size; y++ {
		for x := 0; x < f.RowSize; x++ {
			if f.eval(x, y) {
				result = append(result, Pos{X: x, Y: y})
			}
		}
	}

	return
}

func (f *Field) eval(x, y int) bool {
	value := f.Points[y][x].Value
	for _, pos := range f.positions(x, y) {
		comp := f.Points[pos[1]][pos[0]].Value
		if value > comp {
			return false
		}
	}
	return true
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

	return
}

func (f *Field) validates(x, y int, cx, cy int) (result [][2]int) {
	check := [2]int{x, y - 1}
	if (check[0] != cx || check[1] != cy) && check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x - 1, y}
	if (check[0] != cx || check[1] != cy) && check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x + 1, y}
	if (check[0] != cx || check[1] != cy) && check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}
	check = [2]int{x, y + 1}
	if (check[0] != cx || check[1] != cy) && check[0] >= 0 && check[0] < f.RowSize && check[1] >= 0 && check[1] < len(f.Points) {
		result = append(result, check)
	}

	return
}
