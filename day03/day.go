package day03

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"

	"github.com/fiurthorn/advent/lib"
)

type Day struct{}

//go:embed example.txt
var dayExample string

//go:embed data.txt
var dayData string

func (d Day) Run() {
	log.Printf("Example:  %v", process(dayExample))
	log.Printf("Example:  %v", process(dayData))
}

func process(data string) string {
	lines := lib.Lines(data)
	return fmt.Sprintf("1:%v", process1(lines))
}

func process1(lines []string) string {
	size := len(lines[0])

	ms0th := make([]int, size)
	ms1th := make([]int, size)
	for _, line := range lines {
		for i := 0; i < size; i++ {
			switch line[i] {
			case '0':
				ms0th[i]++
			case '1':
				ms1th[i]++
			}
		}
	}

	gamma := parseGamma(ms0th, ms1th)
	epsilon := parseEpsilon(ms0th, ms1th)

	oxygenLines := make([]string, len(lines))
	copy(oxygenLines, lines)
	oxygen := searchOxygen(oxygenLines)

	co2Lines := make([]string, len(lines))
	copy(co2Lines, lines)
	co2 := searchCO2(co2Lines)

	return fmt.Sprintf("%v * %v = %v |  %v * %v = %v", gamma, epsilon, gamma*epsilon, oxygen, co2, oxygen*co2)
}

func searchCO2(lines []string) (oxygen int64) {
	size := len(lines[0])

	for i := 0; i < size; i++ {
		k0th, k1th := 0, 0
		for _, value := range lines {
			if value[i] == '1' {
				k1th++
			} else {
				k0th++
			}
		}
		cache := []string{}
		for _, value := range lines {
			if k1th >= k0th && value[i] == '0' {
				cache = append(cache, value)
			} else if k1th < k0th && value[i] == '1' {
				cache = append(cache, value)
			}
		}
		if len(cache) == 1 {
			oxygen, _ = strconv.ParseInt(string(cache[0]), 2, 16)
			return
		}
		lines = cache
	}

	return
}

func searchOxygen(lines []string) (oxygen int64) {
	size := len(lines[0])

	for i := 0; i < size; i++ {
		k0th, k1th := 0, 0
		for _, value := range lines {
			if value[i] == '1' {
				k1th++
			} else {
				k0th++
			}
		}
		cache := []string{}
		for _, value := range lines {
			if k1th >= k0th && value[i] == '1' {
				cache = append(cache, value)
			} else if k1th < k0th && value[i] == '0' {
				cache = append(cache, value)
			}
		}
		if len(cache) == 1 {
			oxygen, _ = strconv.ParseInt(string(cache[0]), 2, 16)
			return
		}
		lines = cache
	}

	return
}

func parseGamma(p0th, p1th []int) (gamma int64) {
	size := len(p0th)
	gammaSlice := make([]byte, size)
	for i := 0; i < size; i++ {
		if p1th[i] > p0th[i] {
			gammaSlice[i] = '1'
		} else {
			gammaSlice[i] = '0'
		}
	}
	gamma, _ = strconv.ParseInt(string(gammaSlice), 2, 16)
	return
}

func parseEpsilon(p0th, p1th []int) (gamma int64) {
	size := len(p0th)
	gammaSlice := make([]byte, size)
	for i := 0; i < size; i++ {
		if p1th[i] > p0th[i] {
			gammaSlice[i] = '0'
		} else {
			gammaSlice[i] = '1'
		}
	}
	gamma, _ = strconv.ParseInt(string(gammaSlice), 2, 16)
	return
}
