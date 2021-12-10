package dayEX

import (
	_ "embed"
	"log"
)

type Day struct{}

//go:embed example.txt
var dayExample string

//go:embed data.txt
var dayData string

func (d Day) Run() {
	log.Printf("Example:  %v", dayExample)
	log.Printf("Solution: %v", dayData)
}
