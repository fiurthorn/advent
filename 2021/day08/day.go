package day08

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"sort"
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

func (d Day) process1(lines []string) string {
	var result [10]int
	var sum int
	for _, line := range lines {
		s := &Segments{}
		s.Samples, s.Requests = d.split(line)

		s.One = s.findBySize(2)
		s.Seven = s.findBySize(3)
		s.Four = s.findBySize(4)
		s.Eight = s.findBySize(7)
		s.findBD()
		s.Five = s.findBySizePattern(5, s.BD)
		s.Three = s.findBySizePattern(5, s.Seven)
		s.Two = s.findBySize(5)
		s.Nine = s.findBySizePattern2(6, s.Seven, s.Five)
		s.Six = s.findBySizePattern(6, s.Five)
		s.Zero = s.findBySize(6)

		ssum := 0
		for i := 0; i < len(s.Requests); i++ {
			digit := s.Find(s.sort(s.Requests[i]))
			result[digit]++
			ssum += int(math.Pow10(3-i)) * digit
		}
		sum += ssum
	}
	return fmt.Sprintf("%v | %v", result[1]+result[4]+result[7]+result[8], sum)
}

type Segments struct {
	Samples     []string
	Requests    []string
	BD          string
	One, Seven  string
	Four, Eight string
	Five, Three string
	Two, Zero   string
	Nine, Six   string
}

func (s *Segments) Find(pattern string) (result int) {
	switch pattern {
	case s.Zero:
		return 0
	case s.One:
		return 1
	case s.Two:
		return 2
	case s.Three:
		return 3
	case s.Four:
		return 4
	case s.Five:
		return 5
	case s.Six:
		return 6
	case s.Seven:
		return 7
	case s.Eight:
		return 8
	case s.Nine:
		return 9
	}
	panic("no number found")
}

func (s *Segments) Contains(a, b string) (result bool) {
	result = true
	for _, r := range b {
		result = result && strings.Contains(a, string(r))
	}
	return
}

func (s *Segments) findBD() {
	s.BD = s.Sub(s.Four, s.One)
}

func (s *Segments) Sub(a, b string) (result string) {
	result = a
	for _, r := range b {
		result = strings.Replace(result, string(r), "", 1)
	}
	return
}

func (s *Segments) findBySizePattern2(size int, pattern, pattern2 string) string {
	for i, length := 0, len(s.Samples); i < length; i++ {
		word := s.Samples[i]
		if len(word) == size && s.Contains(word, pattern) && s.Contains(word, pattern2) {
			s.Samples[i] = s.Samples[length-1]
			s.Samples = s.Samples[:length-1]
			return s.sort(word)
		}
	}
	panic("should not happen!")
}

func (s *Segments) findBySizePattern(size int, pattern string) string {
	for i, length := 0, len(s.Samples); i < length; i++ {
		word := s.Samples[i]
		if len(word) == size && s.Contains(word, pattern) {
			s.Samples[i] = s.Samples[length-1]
			s.Samples = s.Samples[:length-1]
			return s.sort(word)
		}
	}
	panic("should not happen!")
}

func (s *Segments) sort(content string) string {
	r := RuneSlice([]rune(content))
	sort.Sort(r)
	return string(r)
}

func (s *Segments) findBySize(size int) string {
	for i, length := 0, len(s.Samples); i < length; i++ {
		word := s.Samples[i]
		if len(word) == size {
			s.Samples[i] = s.Samples[length-1]
			s.Samples = s.Samples[:length-1]
			return s.sort(word)
		}
	}
	panic("should not happen!")
}

func (d Day) split(line string) ([]string, []string) {
	parts := strings.Split(line, " | ")
	return lib.Strings(parts[0]), lib.Strings(parts[1])
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
