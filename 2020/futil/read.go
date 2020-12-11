package futil

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntSlice(path string) ([]int, error) {
	scanner, closeFunc, err := getFileScanner(path)
	if err != nil {
		return nil, err
	}
	defer closeFunc()
	var numbers []int
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func ReadStringSlice(path string) ([]string, error) {
	scanner, closeFunc, err := getFileScanner(path)
	if err != nil {
		return nil, err
	}
	defer closeFunc()
	var strings []string
	for scanner.Scan() {
		text := scanner.Text()
		strings = append(strings, text)
	}
	return strings, nil
}

func ReadBytesSlice(path string) ([][]byte, error) {
	scanner, closeFunc, err := getFileScanner(path)
	if err != nil {
		return nil, err
	}
	defer closeFunc()
	var bytesSlice [][]byte
	for scanner.Scan() {
		text := scanner.Bytes()
		bytesSlice = append(bytesSlice, text)
	}
	return bytesSlice, nil
}

func getFileScanner(path string) (*bufio.Scanner, func() error, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(f)
	return scanner, f.Close, nil
}
