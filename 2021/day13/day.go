package day13

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
	return fmt.Sprintf("1:%v",
		d.process1(lines),
	)
}

type Cmd struct {
	axis  string
	value int
}

func (d Day) process1(lines []string) string {
	maxX, maxY := 0, 0
	command := false
	commands := []Cmd{}
	for _, line := range lines {
		if len(line) == 0 {
			command = true
		} else if !command {
			xy := lib.Numbers(line)
			if xy[0] > maxX {
				maxX = xy[0]
			}
			if xy[1] > maxY {
				maxY = xy[1]
			}
		} else {
			result := strings.Split(line, " ")
			cmd := strings.Split(result[2], "=")
			c := Cmd{axis: cmd[0], value: lib.Atoi(cmd[1])}
			commands = append(commands, c)
		}
	}
	maxX++
	maxY++

	field := make([][]rune, maxY)
	for y := 0; y < maxY; y++ {
		field[y] = make([]rune, maxX)
	}

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		xy := lib.Numbers(line)
		field[xy[1]][xy[0]] = 1
	}

	sum := 0
	for _, c := range commands {
		if c.axis == "y" {
			align := c.value
			for i := 1; i <= c.value; i++ {
				for x := 0; x < len(field[align]); x++ {
					if field[align+i][x] != 0 {
						field[align-i][x] = 1
					}
				}
			}
			field = field[:len(field)-c.value-1]
		}
		if c.axis == "x" {
			align := c.value
			for y := 0; y < len(field); y++ {
				for i := 1; i <= c.value; i++ {
					if field[y][align+i] != 0 {
						field[y][align-i] = 1
					}
				}
				field[y] = field[y][:len(field[y])-c.value-1]
			}
		}

		if sum == 0 {
			for _, l := range field {
				for _, s := range l {
					if s == 1 {
						sum++
					}
				}
			}
		}
	}

	sb := strings.Builder{}
	for _, l := range field {
		for _, r := range l {
			if r == 0 {
				sb.WriteRune(' ')
			} else {
				sb.WriteRune('#')
			}
		}
		sb.WriteRune('\n')
	}

	return fmt.Sprintf("%v\n%s", sum, sb.String())
}
