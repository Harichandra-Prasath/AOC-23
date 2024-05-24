package day2

import (
	"os"
	"strconv"
	"strings"
)

var Total = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part1() int {
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	f := string(data)
	lines := strings.Split(f, "\n")
	sum := 0
	for ind, line := range lines {
		pos := isPossible(line)
		if pos {
			sum += ind + 1
		}

	}
	return sum
}

func isPossible(line string) bool {
	game := strings.Split(line, ":")[1]
	sets := strings.Split(game, ";")
	for _, set := range sets {
		raw := strings.Replace(set, " ", "", -1)
		raw = strings.Replace(raw, ",", "", -1)
		if !check_in_map(raw) {
			return false
		}
	}

	return true
}

func check_in_map(raw string) bool {
	num := ""
	ind := 0
	for ind < len(raw) {
		r := raw[ind]
		if r <= 57 && r >= 48 {
			num += string(r)
			ind += 1
			continue
		}
		if string(r) == "g" {
			n, _ := strconv.Atoi(num)
			if n > Total["green"] {
				return false
			}
			ind += 5
			num = ""

		} else if string(r) == "b" {
			n, _ := strconv.Atoi(num)
			if n > Total["blue"] {
				return false
			}
			ind += 4
			num = ""
		} else {
			n, _ := strconv.Atoi(num)
			if n > Total["red"] {
				return false
			}
			ind += 3
			num = ""
		}
	}

	return true

}
