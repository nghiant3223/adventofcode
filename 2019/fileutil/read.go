package fileutil

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	separator = ","
)

func ReadIntSlice(path string) ([]int, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var numbers []int
	content := string(f)
	for _, text := range strings.Split(content, separator) {
		number, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}
