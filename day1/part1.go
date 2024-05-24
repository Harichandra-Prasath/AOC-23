package day1

import (
	"os"
	"strconv"
)

func Part1() int {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	num := ""
	flag := false
	c := ""
	for _, b := range data {
		if string(b) == "\n" {
			num += c
			value, _ := strconv.Atoi(num)
			sum += value
			num = ""
			c = ""
			flag = false
			continue
		}
		check1(b, &num, &flag, &c)
	}
	return sum
}

func check1(b byte, num *string, flag *bool, c *string) {

	r := rune(b)
	if r <= 57 && r >= 48 {
		if !*(flag) {
			*(num) += string(b)
			*(flag) = true
		}
		*(c) = string(b)
	}

}
