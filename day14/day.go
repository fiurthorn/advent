package day14

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

type ListNode struct {
	Value rune
	Next  *ListNode
}

type LinkedList struct {
	Start *ListNode
	End   *ListNode
}

func (l *LinkedList) Index() (cache map[rune]int) {
	cache = map[rune]int{}
	for p := l.Start.Next; p != l.End; p = p.Next {
		cache[p.Value]++
	}
	return
}

func (l *LinkedList) Grow(cmd map[string]rune) {
	for p := l.Start.Next; p != l.End; p = p.Next.Next {
		index := string([]rune{p.Value, p.Next.Value})
		ins := &ListNode{Value: cmd[index]}
		ins.Next, p.Next = p.Next, ins
	}
}

func (l *LinkedList) String() string {
	sb := strings.Builder{}
	for p := l.Start.Next; p != l.End; p = p.Next {
		sb.WriteRune(p.Value)
	}
	return sb.String()
}

func (d Day) process1(lines []string) string {
	list := &LinkedList{Start: &ListNode{}, End: &ListNode{}}
	list.Start.Next = list.End

	init := lines[0]
	var prev *ListNode = list.Start
	for _, r := range init {
		cur := &ListNode{Value: r, Next: prev.Next}
		prev.Next, prev = cur, cur
	}

	cmd := map[string]rune{}
	for i := 2; i < len(lines); i++ {
		part := strings.Split(lines[i], " -> ")
		cmd[part[0]] = rune(part[1][0])
	}

	for i := 0; i < 10; i++ {
		list.Grow(cmd)
	}
	min, max := -1, -1
	for _, v := range list.Index() {
		if max == -1 {
			max = v
		} else if v > max {
			max = v
		}
		if min == -1 {
			min = v
		} else if v < min {
			min = v
		}
	}

	return fmt.Sprintf("%d - %d = %v", max, min, max-min)
}

func (d Day) process2(lines []string) string {
	return fmt.Sprintf("%v", 0)
}
