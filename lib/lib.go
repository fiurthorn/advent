package lib

import (
	"log"
	"strconv"
	"strings"
)

func Atoi(value string) (current int) {
	current, err := strconv.Atoi(value)
	if err != nil {
		log.Panic(err)
	}
	return
}

func Lines(lines string) []string {
	if strings.ContainsRune(lines, '\r') {
		return strings.Split(lines, "\r\n")
	}
	return strings.Split(lines, "\n")
}

func Numbers(line string) (result []int) {
	nums := strings.Split(line, ",")
	for _, num := range nums {
		result = append(result, Atoi(num))
	}
	return
}

func Strings(line string) (result []string) {
	result = strings.Split(line, " ")
	return
}

type Queue []rune

func (q *Queue) Push(x rune) {
	*q = append(*q, x)
}

func (q *Queue) Pop() (el rune) {
	l := len(*q)
	el, *q = (*q)[l-1], (*q)[0:l-1]
	return el
}

func (q *Queue) Unshift() (el rune) {
	l := len(*q)
	el, *q = (*q)[0], (*q)[1:l]
	return el
}

func (q *Queue) Shift(el rune) {
	*q = append(*q, el)
	if len(*q) == 1 {
		return
	}
	copy((*q)[1:], *q)
	(*q)[0] = el
}

func (q *Queue) Clear() {
	*q = []rune{}
}

func (q *Queue) String() string {
	return string(*q)
}

func NewQueue() *Queue {
	return &Queue{}
}
