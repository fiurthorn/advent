package lib

import (
	"fmt"
	"log"
	"sort"
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

type void int
type StringSet map[string]void

const empty = void(0)

func NewStringSet() *StringSet {
	return &StringSet{}
}

func NewStringSetWith(keys ...string) *StringSet {
	s := &StringSet{}

	for _, key := range keys {
		s.Add(key)
	}

	return s
}

func (m *StringSet) Remove(key string) {
	delete(*m, key)
}

func (m *StringSet) Add(key string) {
	(*m)[key] = empty
}

func (m *StringSet) Values() []string {
	keys := []string{}
	for k := range *m {
		keys = append(keys, k)
	}
	return keys
}

func (m *StringSet) SortedValues() []string {
	values := m.Values()
	sort.Strings(values)
	return values
}

func (m *StringSet) String() string {
	return fmt.Sprintf("StringSet[%s]", strings.Join(m.Values(), ", "))
}

func (m *StringSet) Len() int {
	return len(*m)
}

func (m *StringSet) Has(value string) bool {
	_, ok := (*m)[value]
	return ok
}

func (m *StringSet) Clone() (r *StringSet) {
	r = &StringSet{}
	for k, v := range *m {
		(*r)[k] = v
	}
	return
}

func (m *StringSet) CloneWith(value string) (r *StringSet) {
	r = m.Clone()
	r.Add(value)
	return
}
