package one

import (
	"aoc/helper"
	"fmt"
	"strconv"
	"strings"
)

func A() int {
	input := helper.ReadFile("1/input.txt")
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		var nums []string
		for _, c := range line {
			if _, err := strconv.Atoi(string(c)); err == nil {
				nums = append(nums, string(c))
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1]))
		sum += num
	}
	return sum
}

var numMap = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func isNumber(s string) (string, error) {
	for n, ns := range numMap {
		if len(ns) <= len(s) {
			if s[0:len(ns)] == ns {
				return n, nil
			}
		}
	}
	return "", fmt.Errorf("not a number")
}

func B() int {
	input := helper.ReadFile("1/input.txt")
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		var nums []string
		for i, c := range line {
			if _, err := strconv.Atoi(string(c)); err == nil {
				nums = append(nums, string(c))
			}
			if ns, err := isNumber(line[i:]); err == nil {
				nums = append(nums, ns)
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1]))
		sum += num
	}
	return sum
}
