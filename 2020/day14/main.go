package day14

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const memorySpace = 36
const paddingFormat = "%036s"

var maskPattern = regexp.MustCompile(`mask = (\w+)`)
var writePattern = regexp.MustCompile(`mem\[(\d+)] = (\d+)`)

func part1(lines []string) int64 {
	memory := make(map[string]int64)
	var currentMask string
	for _, line := range lines {
		maskMatch := maskPattern.FindStringSubmatch(line)
		if maskMatch != nil {
			currentMask = maskMatch[1]
			continue
		}
		writeMatch := writePattern.FindStringSubmatch(line)
		if writeMatch != nil {
			address := writeMatch[1]
			value := writeMatch[2]
			binValue := decToBin(value)
			appliedMaskValue := applyMask(currentMask, binValue)
			memory[address] = binStringToDec(appliedMaskValue)
			continue
		}
		log.Panicf("invalid line: %s", line)
	}
	var sum int64 = 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func part2(lines []string) int64 {
	memory := make(map[string]int64)
	var currentMask string
	for _, line := range lines {
		maskMatch := maskPattern.FindStringSubmatch(line)
		if maskMatch != nil {
			currentMask = maskMatch[1]
			continue
		}
		writeMatch := writePattern.FindStringSubmatch(line)
		if writeMatch != nil {
			address := writeMatch[1]
			address = decToBin(address)
			value := writeMatch[2]
			decValue := stringToDec(value)
			appliedMaskValue := applyMaskPart2(currentMask, address)
			addressCombinations := generateCombinations(memorySpace, appliedMaskValue)
			for _, combinationAddress := range addressCombinations {
				memory[combinationAddress] = decValue
			}
			continue
		}
		log.Panicf("invalid line: %s", line)
	}
	var sum int64 = 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func stringToDec(str string) int64 {
	int64Dec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Panicf("cannot parse int: %v", err)
	}
	return int64Dec
}

func decToBin(dec string) string {
	int64Dec, err := strconv.ParseInt(dec, 10, 64)
	if err != nil {
		log.Panicf("cannot parse int: %v", err)
	}
	return strconv.FormatInt(int64Dec, 2)
}

func binStringToDec(bin string) int64 {
	int64Dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Panicf("cannot parse int: %v", err)
	}
	return int64Dec
}

func applyMask(mask string, value string) string {
	mask = fmt.Sprintf(paddingFormat, mask)
	value = fmt.Sprintf(paddingFormat, value)
	result := make([]byte, memorySpace)
	for i := 0; i < memorySpace; i++ {
		switch mask[i] {
		case '0', '1':
			result[i] = mask[i]
		case 'X':
			result[i] = value[i]
		default:
			log.Panicf("invalid mask: %s", mask)
		}
	}
	return string(result)
}

func applyMaskPart2(mask string, value string) string {
	mask = fmt.Sprintf(paddingFormat, mask)
	value = fmt.Sprintf(paddingFormat, value)
	result := make([]byte, memorySpace)

	for i := 0; i < memorySpace; i++ {
		switch mask[i] {
		case 'X':
			result[i] = 'X'
		default:
			result[i] = orBit(value[i], mask[i])
		}
	}

	return string(result)
}

func generateCombinations(space int, mask string) []string {
	return generateCombinationUtil(space, mask, -1, []string{})
}

func generateCombinationUtil(space int, mask string, i int, results []string) []string {
	for j := i + 1; j < space; j++ {
		if mask[j] == 'X' {
			bString0 := []byte(mask)
			bString1 := []byte(mask)
			bString0[j] = '0'
			bString1[j] = '1'
			results = generateCombinationUtil(space, string(bString0), j, results)
			results = generateCombinationUtil(space, string(bString1), j, results)
			return results
		}
	}
	results = append(results, mask)
	return results
}

func orBit(a, b byte) byte {
	if a == '0' && b == '0' {
		return '0'
	}
	return '1'
}
