package one

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/adorfman10/advent-of-code/util"
)

type idFunc func(string, int, bool) (int, error)

func runner(idFuncs []idFunc) int {
	var wg sync.WaitGroup

	input := util.ReadFile("1/input.txt")
	lines := strings.Split(input, "\n")
	sum := 0
	allNums := make([][]int, len(lines))
	for li, line := range lines {
		wg.Add(2)
		allNums[li] = make([]int, 2)
		go func(line string, li int) {
			defer wg.Done()
			for i := 0; i < len(line); i++ {
				for _, f := range idFuncs {
					if n, err := f(line, i, false); err == nil {
						allNums[li][0] = n
						return
					}
				}
			}
		}(line, li)

		go func(line string, li int) {
			defer wg.Done()
			for i := len(line) - 1; i >= 0; i-- {
				for _, f := range idFuncs {
					if n, err := f(line, i, true); err == nil {
						allNums[li][1] = n
						return
					}
				}
			}
		}(line, li)
	}
	wg.Wait()
	for _, np := range allNums {
		sum += np[0] * 10
		sum += np[1]
	}
	return sum
}

func AGoRoutine() int {
	defer util.Duration(util.Track("AGoRoutine"))
	return runner([]idFunc{isDigitNumber})
}

func BGoRoutine() int {
	defer util.Duration(util.Track("BGoRoutine"))
	return runner([]idFunc{isDigitNumber, isStringNumber})
}

var numMapGoRoutine = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func isStringNumber(s string, idx int, rev bool) (int, error) {
	for n, ns := range numMapGoRoutine {
		rs := s[idx:]
		if strings.HasPrefix(rs, ns) {
			return n, nil
		}
	}
	return 0, fmt.Errorf("not a number")
}

func isDigitNumber(s string, idx int, rev bool) (int, error) {
	c := string(s[idx])
	if n, err := strconv.Atoi(c); err == nil {
		return n, nil
	} else {
		return 0, err
	}
}
