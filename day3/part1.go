package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var length int
var Mine []int

func Part1() int {
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	f := string(data)
	lines := strings.Split(f, "\n")
	sum := 0
	length = len(lines[0])

	for ind := range lines {
		top := ""
		bot := ""
		mid := ""
		if ind-1 >= 0 {
			top = lines[ind-1]
		} else {
			i := 0
			for i < length {
				top += "."
				i += 1
			}
		}
		if ind+1 < len(lines) {
			bot = lines[ind+1]

		} else {
			i := 0
			for i < length {
				bot += "."
				i += 1
			}
		}
		mid = lines[ind]

		sum += get_total(top, mid, bot)
	}
	return sum

}

func get_total(top string, mid string, bot string) int {
	start := -1
	end := 0
	total := 0
	for ind := range mid {
		r := mid[ind]
		if r <= 57 && r >= 48 {
			if start == -1 {
				start = ind
			}
			if ind == length-1 {
				end = ind
				total += check_valid(top, mid, bot, start, end)
			}
			continue
		}

		if start != -1 {
			end = ind - 1
			total += check_valid(top, mid, bot, start, end)
			start = -1
			end = 0
		}

	}
	return total
}

func check_valid(top string, mid string, bot string, start int, end int) int {
	i := start
	flag := false
	num := ""
	for i <= end {
		num += string(mid[i])
		i += 1
	}
	i = start
	for i <= end {
		if i == start {
			// check for upper, upper left diagnol, left, down left diagnol, down
			if check_top_bot(top, bot, i) {
				flag = true
				break
			}

			//left
			if i-1 > 0 {
				if check_left_space(top, mid, bot, i) {
					flag = true
					break
				}
			}

		}
		if i == end {
			if check_top_bot(top, bot, i) {
				flag = true
				break
			}

			if i+1 < length {
				if check_right_space(top, mid, bot, i) {
					flag = true
					break
				}
			}
		}
		if i != start && i != end {
			if check_top_bot(top, bot, i) {
				flag = true
				break
			}

		}

		i += 1
	}
	if !flag {
		return 0
	} else {
		n, _ := strconv.Atoi(num)
		Mine = append(Mine, n)
		return n
	}
}

func check_top_bot(top string, bot string, index int) bool {
	n := 0
	if string(top[index]) == "." || (top[index] <= 57 && top[index] >= 48) {
		n += 1
	}
	if string(bot[index]) == "." || (bot[index] <= 57 && bot[index] >= 48) {
		n += 1
	}
	return n != 2
}

func check_left_space(top string, mid string, bot string, index int) bool {
	n := 0
	if string(top[index-1]) == "." || (top[index-1] <= 57 && top[index-1] >= 48) {
		if string(mid[index]) == "7" && string(mid[index-1]) == "7" {
			fmt.Println("I got recognised as .")
		}
		n += 1
	}
	if string(bot[index-1]) == "." || (bot[index-1] <= 57 && bot[index-1] >= 48) {
		n += 1
	}
	if string(mid[index-1]) == "." || (mid[index-1] <= 57 && mid[index-1] >= 48) {
		n += 1
	}
	return n != 3
}

func check_right_space(top string, mid string, bot string, index int) bool {
	n := 0
	if string(top[index+1]) == "." || (top[index+1] <= 57 && top[index+1] >= 48) {
		n += 1
	}
	if string(bot[index+1]) == "." || (bot[index+1] <= 57 && bot[index+1] >= 48) {
		n += 1
	}
	if string(mid[index+1]) == "." || (mid[index+1] <= 57 && mid[index+1] >= 48) {
		n += 1
	}
	return n != 3
}
