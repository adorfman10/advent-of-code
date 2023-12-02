package one

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adorfman10/advent-of-code/util"
)

func A() int {
	defer util.Duration(util.Track("A"))

	input := util.ReadFile("1/input.txt")
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
		if strings.HasPrefix(s, ns) {
			return n, nil
		}
	}
	return "", fmt.Errorf("not a number")
}

func B() int {
	defer util.Duration(util.Track("B"))

	input := util.ReadFile("1/input.txt")
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
