package day1

import (
	"os"
	"strconv"
	"strings"
)

var Mine = []int{}

var match = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func Part2() int {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	f := string(data)
	lines := strings.Split(f, "\n")
	for _, line := range lines {
		num := check2(line)
		sum += num
	}
	return sum

}

func check2(line string) int {
	flag := false
	num := ""
	var last string = ""
	ind := 0
	for ind < len(line) {
		c := line[ind]
		if c <= 57 && c >= 48 {
			if !(flag) {
				(num) += string(c)
				(flag) = true
			}
			last = string(c)

			ind += 1
			continue
		}
		text := string(c)
		j := ind + 1
		for {
			if j == len(line) {
				ind += 1
				break
			}
			text += string(line[j])
			if check_in_map(text) == 0 {
				ind += 1
				break
			} else if check_in_map(text) == 1 {
				j += 1
				continue
			} else {
				// it is a word
				val, pos := match[text]
				if !(flag) {
					if pos {
						(num) += val
					}
					(flag) = true
				}
				last = val

				ind += 1
				break
			}

		}

	}
	num += last

	n, _ := strconv.Atoi(num)
	Mine = append(Mine, n)
	return n
}

func check_in_map(text string) int {
	if _, pos := match[text]; pos {
		return 2
	}
	if check_prefix(text) {
		return 1
	}
	return 0

}

func check_prefix(text string) bool {
	for k := range match {
		if strings.HasPrefix(k, text) {
			return true
		}
	}
	return false
}
