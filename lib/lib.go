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