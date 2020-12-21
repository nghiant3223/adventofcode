package lutil

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Panicf("cannot convert string to int: %v\n", e)
		return 0
	}
	return i
}
