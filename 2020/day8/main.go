package day8

import (
	"log"
	"regexp"
	"strconv"
)

type Instruction struct {
	Operation string
	Argument  int
}

var instructionPattern = regexp.MustCompile(`(\w+) ([+-]\d+)`)

func prepareInput(lines []string) []Instruction {
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		match := instructionPattern.FindStringSubmatch(line)
		operation := match[1]
		argument, err := strconv.Atoi(match[2])
		if err != nil {
			log.Panic(err)
		}
		instruction := Instruction{
			Operation: operation,
			Argument:  argument,
		}
		instructions[i] = instruction
	}
	return instructions
}

func part1(instructions []Instruction) int {
	accumulate := 0
	visited := make(map[int]struct{}, len(instructions))
	i := 0
	for i < len(instructions) {
		_, ok := visited[i]
		if ok {
			return accumulate
		}
		visited[i] = struct{}{}
		instruction := instructions[i]
		switch instruction.Operation {
		case "acc":
			accumulate += instruction.Argument
			i++
		case "jmp":
			i += instruction.Argument
		case "nop":
			i++
		}
	}
	return accumulate
}

func part2(instructions []Instruction) int {
	accumulate := 0
	visited := make(map[int]struct{}, len(instructions))
	i := 0
	lastI := -1
	for i < len(instructions) {
		_, ok := visited[i]
		if ok {
			instructions[lastI].Operation = invertInstruction(instructions[lastI].Operation)
			break
		}
		lastI = i
		visited[i] = struct{}{}
		instruction := instructions[i]
		switch instruction.Operation {
		case "acc":
			accumulate += instruction.Argument
			i++
		case "jmp":
			i += instruction.Argument
		case "nop":
			i++
		}
	}
	return part1(instructions)
}

func invertInstruction(instr string) string {
	if instr == "jmp" {
		return "nop"
	}
	return "acc"
}
