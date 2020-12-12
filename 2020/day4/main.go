package day4

import (
	"log"
	"regexp"
	"strconv"
)

const (
	byr = "byr"
	iyr = "iyr"
	eyr = "eyr"
	hgt = "hgt"
	hcl = "hcl"
	ecl = "ecl"
	pid = "pid"
	cid = "cid"
)

var yearPattern = regexp.MustCompile("^(\\d{4,})$")
var passwordPattern = regexp.MustCompile("(\\w+):([^\\s]+)")
var colorPattern = regexp.MustCompile("^#[0-9a-f]{6,}$")
var passportPattern = regexp.MustCompile("^[0-9]{9,}$")
var heightPattern = regexp.MustCompile("^(\\d+)(in|cm)$")

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

type Password struct {
	BYR string
	IYR string
	EYR string
	HGT string
	HCL string
	ECL string
	PID string
	CID string
}

func (p Password) IsValid() bool {
	return p.BYR != "" && p.IYR != "" && p.EYR != "" && p.HGT != "" && p.HCL != "" && p.ECL != "" && p.PID != ""
}

func (p Password) IsValidPart2() bool {
	if !p.IsValid() {
		return false
	}

	match := yearPattern.FindStringSubmatch(p.BYR)
	if match == nil {
		return false
	}
	byrString := match[1]
	byr, err := strconv.Atoi(byrString)
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	match = yearPattern.FindStringSubmatch(p.IYR)
	if match == nil {
		return false
	}
	iyrString := match[1]
	iyr, err := strconv.Atoi(iyrString)
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	match = yearPattern.FindStringSubmatch(p.EYR)
	if match == nil {
		return false
	}
	eyrString := match[1]
	eyr, err := strconv.Atoi(eyrString)
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	match = heightPattern.FindStringSubmatch(p.HGT)
	if match == nil {
		return false
	}
	heightString, unit := match[1], match[2]
	height, err := strconv.Atoi(heightString)
	if err != nil {
		return false
	}
	switch unit {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}

	if !colorPattern.MatchString(p.HCL) {
		return false
	}

	if !contains(eyeColors, p.ECL) {
		return false
	}

	if !passportPattern.MatchString(p.PID) {
		return false
	}

	return true
}

func numberOfValidPassports(lines []string) int {
	passwords := reformPasswords(lines)
	validPasswordCount := 0
	for _, password := range passwords {
		if password.IsValid() {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func numberOfValidPassportsPart2(lines []string) int {
	passwords := reformPasswords(lines)
	validPasswordCount := 0
	for _, password := range passwords {
		if password.IsValidPart2() {
			log.Printf("%+v: true", password)
			validPasswordCount++
		} else {
			log.Printf("%+v: false", password)
		}
	}
	return validPasswordCount
}

func reformPasswords(lines []string) []Password {
	var passwords []Password
	var linesOfPassword []string
	for _, line := range lines {
		if line == "" {
			password := gatherPasswordData(linesOfPassword)
			passwords = append(passwords, password)
			linesOfPassword = []string{}
			continue
		}
		linesOfPassword = append(linesOfPassword, line)
	}
	password := gatherPasswordData(linesOfPassword)
	passwords = append(passwords, password)
	return passwords
}

func gatherPasswordData(lines []string) Password {
	var password Password
	for _, line := range lines {
		matches := passwordPattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			assignKeyValue(&password, match[1], match[2])
		}
	}
	return password
}

func assignKeyValue(pwd *Password, key, value string) {
	switch key {
	case byr:
		pwd.BYR = value
	case iyr:
		pwd.IYR = value
	case eyr:
		pwd.EYR = value
	case hgt:
		pwd.HGT = value
	case hcl:
		pwd.HCL = value
	case ecl:
		pwd.ECL = value
	case pid:
		pwd.PID = value
	case cid:
		pwd.CID = value
	}
}

func contains(s []string, t string) bool {
	for _, si := range s {
		if si == t {
			return true
		}
	}
	return false
}
