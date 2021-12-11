package day07

import (
	_ "embed"
	"fmt"
	"log"
	"math"

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
	nums := lib.Numbers(data)

	return fmt.Sprintf("1:%v, 2:%v",
		d.process1(nums),
		d.process2(nums),
	)
}

func (d Day) process1(numbers []int) string {
	max := 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	fuels := make([]int, max)
	for i := 0; i < max; i++ {
		for j := 0; j < len(numbers); j++ {
			num := numbers[j]
			fuels[i] += int(math.Abs(float64(i - num)))
		}
	}

	min := -1
	for i := 0; i < max; i++ {
		if min < 0 || fuels[i] < min {
			min = fuels[i]
		}
	}

	return fmt.Sprintf("%v", min)
}

func (d Day) process2(numbers []int) string {
	max := 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	fuels := make([]int, max)
	for i := 0; i < max; i++ {
		for j := 0; j < len(numbers); j++ {
			num := numbers[j]
			fuels[i] += d.sum(int(math.Abs(float64(i - num))))
		}
	}

	min := -1
	for i := 0; i < max; i++ {
		if min < 0 || fuels[i] < min {
			min = fuels[i]
		}
	}

	return fmt.Sprintf("%v", min)
}

func (d Day) sum(num int) (result int) {
	for i := 1; i <= num; i++ {
		result += i
	}
	return
}
