package day06

import (
	_ "embed"
	"fmt"
	"log"

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

	return fmt.Sprintf("1:%v, 1:%v, 2:%v",
		d.process1(nums, 18),
		d.process1(nums, 80),
		d.process1(nums, 256),
	)
}

func (d Day) process1(numbers []int, days int) string {
	var fishes = [9]uint64{}

	for _, num := range numbers {
		fishes[num]++
	}

	for y := 0; y < days; y++ {
		zero := fishes[0]

		fishes[0] = fishes[1]
		fishes[1] = fishes[2]
		fishes[2] = fishes[3]
		fishes[3] = fishes[4]
		fishes[4] = fishes[5]
		fishes[5] = fishes[6]
		fishes[6] = fishes[7] + zero
		fishes[7] = fishes[8]
		fishes[8] = zero
	}

	result := fishes[0] +
		fishes[1] +
		fishes[2] +
		fishes[3] +
		fishes[4] +
		fishes[5] +
		fishes[6] +
		fishes[7] +
		fishes[8]

	// for y := 0; y < days; y++ {
	// 	fmt.Print("\r", y)
	// 	l := len(numbers)
	// 	for i := 0; i < l; i++ {
	// 		v := numbers[i]
	// 		if v == 0 {
	// 			numbers[i] = 6
	// 			numbers = append(numbers, 8)
	// 		} else {
	// 			numbers[i] = v - 1
	// 		}
	// 	}
	// }
	return fmt.Sprintf("%d", result)
}
