package day2

import (
	"os"
	"strconv"
	"strings"
)

func Part2() int {
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	f := string(data)
	lines := strings.Split(f, "\n")
	sum := 0
	for _, line := range lines {
		sum += get_power(line)

	}
	return sum
}

func get_power(line string) int {
	max_r := -1
	max_b := -1
	max_g := -1
	game := strings.Split(line, ":")[1]
	sets := strings.Split(game, ";")
	for _, set := range sets {
		raw := strings.Replace(set, " ", "", -1)
		raw = strings.Replace(raw, ",", "", -1)
		update_maxi(raw, &max_r, &max_b, &max_g)
	}

	return (max_r) * (max_b) * (max_g)
}

func update_maxi(raw string, max_r *int, max_b *int, max_g *int) {
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
			if n > *(max_g) {
				*(max_g) = n
			}
			ind += 5
			num = ""

		} else if string(r) == "b" {
			n, _ := strconv.Atoi(num)
			if n > *(max_b) {
				*(max_b) = n
			}
			ind += 4
			num = ""
		} else {
			n, _ := strconv.Atoi(num)
			if n > *(max_r) {
				*(max_r) = n
			}
			ind += 3
			num = ""
		}
	}
}
